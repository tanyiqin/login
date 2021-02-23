package main

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"go.uber.org/zap"
	"login/conf"
	"login/dao"
	"login/logger"
	"login/mysql"
	"login/redis"
	"net"
	"net/http"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
	"time"
)

func main() {
	// 初始化数据库
	mainConfig := conf.ParseConfig("./conf/conf.yaml")
	mysql.InitMysqlDB(mainConfig)
	redis.InitRedisDB(mainConfig)
	// 初始化日志
	loggerConfig := mainConfig.GetLoggerConfig()
	var log, flog *zap.Logger = nil, nil
	if loggerConfig.File != "" && loggerConfig.Rotation == true{
		flog = logger.NewRotationJSONFileLogger(mainConfig, logger.LevelMapToZapLevel(loggerConfig.Level))
	} else if loggerConfig.File != "" {
		flog = logger.NewJSONFileLogger(loggerConfig.File, logger.LevelMapToZapLevel(loggerConfig.Level))
	}

	if loggerConfig.Stdout == true && flog != nil{
		log = logger.NewMultiLogger(flog, logger.TempLogger)
	} else if flog != nil {
		log = flog
	} else {
		log = logger.TempLogger
	}

	// gin 中间件
	// ssl
	secureMiddleware := secure.New(secure.Options{
		SSLRedirect: true,
	})
	secureFunc := func() gin.HandlerFunc {
		return func(c *gin.Context) {
			err := secureMiddleware.Process(c.Writer, c.Request)
			if err != nil {
				c.Abort()
				return
			}
		}
	}
	// 日志
	loggerFunc := func() gin.HandlerFunc{
		return func(c *gin.Context) {
			start := time.Now()
			path := c.Request.URL.Path
			query := c.Request.URL.RawQuery
			c.Next()

			cost := time.Since(start)
			log.Info(path,
				zap.Int("status", c.Writer.Status()),
				zap.String("method", c.Request.Method),
				zap.String("path", path),
				zap.String("query", query),
				zap.String("ip", c.ClientIP()),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
				zap.Duration("cost", cost))
		}
	}
	// 错误捕获
	recoveryFunc := func() gin.HandlerFunc{
		return func(c *gin.Context) {
			defer func() {
				if err := recover(); err != nil {
					var brokenPipe bool
					if ne, ok := err.(*net.OpError); ok {
						if se, ok := ne.Err.(*os.SyscallError); ok {
							stringError := strings.ToLower(se.Error())
							if strings.Contains(stringError, "broken pipe") || strings.Contains(stringError, "connection reset by peer") {
								brokenPipe = true
							}
						}
					}
					httpRequest, _ := httputil.DumpRequest(c.Request, false)
					headers := strings.Split(string(httpRequest), "\r\n")
					for idx, header := range headers {
						current := strings.Split(header, ":")
						if current[0] == "Authorization" {
							headers[idx] = current[0] + ": *"
						}
					}

					log.Error("[Recovery from panic]",
						zap.Any("error", err),
						zap.String("request", string(httpRequest)),
						zap.String("stack", string(debug.Stack())))

					if brokenPipe {
						c.Error(err.(error))
						c.Abort()
					} else {
						c.AbortWithStatus(http.StatusInternalServerError)
					}
				}
			}()
			c.Next()
		}
	}
	router := gin.New()
	gin.Default()
	accountGroup := router.Group("/account")
	{
		accountGroup.Use(secureFunc(), recoveryFunc(), loggerFunc())
		accountGroup.POST("/create", dao.AccountCreate)
		accountGroup.GET("/login", dao.AccountLogin)
		accountGroup.GET("/test", dao.Test)
	}
	router.RunTLS(":443", "./ssl/server.crt", "./ssl/server.key")
}