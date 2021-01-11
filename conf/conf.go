package conf

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	conf Config
	SqlConf = conf.mysqlConf
)

type Config struct {
	mysqlConf *MySqlConf
}

type MySqlConf struct {
	User	string `yaml:"user"`
	Password	string	`yaml:"password"`
	Protocl	string	`yaml:"protocl"`
	Host string `yaml:"host"`
	Port int `yaml:"port"`
	DBName string `yaml:"dbname"`
	Args string `yaml:"args"`
}

func init() {
	confData, err := ioutil.ReadFile("conf/conf.yaml")
	if err != nil {
		panic(1)
		return
	}
	err = yaml.Unmarshal(confData, &conf)
	if err != nil {
		panic(2)
		return
	}
}