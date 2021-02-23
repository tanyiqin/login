package mysql

import (
	"fmt"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	. "login/conf"
	"login/logger"
)

var err error
var SqlDB *gorm.DB

func InitMysqlDB(config Config) {
	mysqlConfig := config.GetMysqlConfig()
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		mysqlConfig.User, mysqlConfig.Password, mysqlConfig.Protocl, mysqlConfig.Host, mysqlConfig.Port, mysqlConfig.DBName)
	SqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.TempLogger.Fatal("init mysql db error", zap.Error(err))
	}
}