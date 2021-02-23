package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"login/conf"
	"os"
	"path/filepath"
	"strings"
)

var TempLogger = zap.NewExample(zap.AddCaller())

func LevelMapToZapLevel(level string) zapcore.Level {
	switch strings.ToLower(level) {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "error":
		return zapcore.ErrorLevel
	default:
		return zapcore.DebugLevel
	}
}

func NewJSONFileLogger(filename string, level zapcore.Level) *zap.Logger {
	output, err := os.OpenFile(filename, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		TempLogger.Error("Could not create log file", zap.Error(err))
		return nil
	}
	return NewJSONLogger(output, level)
}

func NewRotationJSONFileLogger(config conf.Config, level zapcore.Level) *zap.Logger {
	loggerConfig := config.GetLoggerConfig()
	fileName := loggerConfig.File
	logDir := filepath.Dir(fileName)
	if err := os.MkdirAll(logDir, 0755); err != nil {
		TempLogger.Error("Could not create log dir", zap.Error(err))
		return nil
	}
	jsonEncoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		EncodeLevel: zapcore.LowercaseColorLevelEncoder,
		EncodeTime: zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	})

	writeSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   fileName,
		MaxSize:    loggerConfig.MaxSize,
		MaxAge:     loggerConfig.MaxAge,
		MaxBackups: loggerConfig.MaxBackups,
		LocalTime:  loggerConfig.LocalTime,
		Compress:   loggerConfig.Compress,
	})

	core := zapcore.NewCore(jsonEncoder, writeSyncer, level)
	options := []zap.Option{zap.AddCaller()}
	return zap.New(core, options...)
}

func NewMultiLogger(loggers ...*zap.Logger) *zap.Logger {
	cores := make([]zapcore.Core, 0, len(loggers))
	for _, logger := range loggers {
		cores = append(cores, logger.Core())
	}
	teeCore := zapcore.NewTee(cores...)
	options := []zap.Option{zap.AddCaller()}
	return zap.New(teeCore, options...)
}

func NewJSONLogger(output *os.File, level zapcore.Level) *zap.Logger{
	jsonEncoder := zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		EncodeLevel: zapcore.LowercaseColorLevelEncoder,
		EncodeTime: zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller: zapcore.ShortCallerEncoder,
	})
	core := zapcore.NewCore(jsonEncoder, zapcore.Lock(output), level)
	options := []zap.Option{zap.AddCaller()}
	return zap.New(core, options...)
}
