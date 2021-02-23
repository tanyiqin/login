package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	Mysql MysqlConfig `yaml:"mysql"`
	Redis RedisConfig `yaml:"redis"`
	Logger LoggerConfig `yaml:"logger"`
}

type MysqlConfig struct {
	User	string `yaml:"user"`
	Password	string	`yaml:"password"`
	Protocl	string	`yaml:"protocl"`
	Host string `yaml:"host"`
	Port int `yaml:"port"`
	DBName string `yaml:"dbname"`
	Args string `yaml:"args"`
}

type RedisConfig struct {
	Addr string `yaml:"addr"`
	Password string `yaml:"password"`
	DB int `yaml:"db"`
}

type LoggerConfig struct {
	Level string `yaml:"level" usage:"log level includes debug, info, error"`
	Stdout bool `yaml:"stdout"`
	File string `yaml:"file"`
	Rotation bool `yaml:"rotation"`
	MaxSize int `yaml:"maxsize"`
	MaxAge int `yaml:"maxage"`
	MaxBackups int `yaml:"maxbackups"`
	LocalTime bool `yaml:"localtime"`
	Compress bool `yaml:"compress"`
}

func ParseConfig(confPath string) Config {
	var conf Config
	confData, err := ioutil.ReadFile(confPath)
	if err != nil {
		panic("read conf error")
	}
	err = yaml.Unmarshal(confData, &conf)
	if err != nil {
		panic("unmarshal yaml error")
	}
	return conf
}

func (c Config) GetMysqlConfig() MysqlConfig {
	return c.Mysql
}

func (c Config) GteRedisConfig() RedisConfig {
	return c.Redis
}

func (c Config) GetLoggerConfig() LoggerConfig {
	return c.Logger
}