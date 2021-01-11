package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	. "login/conf"
)

var err error
var SqlDB *gorm.DB

func init() {
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?parseTime=True&%s",
		SqlConf.User, SqlConf.Password, SqlConf.Protocl, SqlConf.Host, SqlConf.Port, SqlConf.DBName, SqlConf.Args)
	SqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("mysql connect err=", err)
		return
	}
}
