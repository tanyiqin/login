package conf

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

var (
	SqlConf MySqlConf
)

type Config struct {
	MySqlConf `yaml:"mysql"`
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
	var conf Config
	confData, err := ioutil.ReadFile("conf/conf.yaml")
	if err != nil {
		panic(1)
		return
	}
	err = yaml.Unmarshal(confData, &conf)
	if err != nil {
		fmt.Print("umarshal err =",err)
		panic(2)
		return
	}
	SqlConf = conf.MySqlConf
	//yfile, err := os.Open("conf/conf.yaml")
	//defer yfile.Close()
	//if err != nil {
	//	log.Panic("err in conf Open yaml,err=%v", err)
	//}
	//ydecode := yaml.NewDecoder(yfile)
	//
	//var c Config
	//err = ydecode.Decode(&c)
	//if err != nil {
	//	log.Panic("err in conf Decode,err=%v", err)
	//}
	//conf = c
}