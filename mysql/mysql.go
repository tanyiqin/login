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
	dsn := fmt.Sprintf("%s:%s@%s(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		SqlConf.User, SqlConf.Password, SqlConf.Protocl, SqlConf.Host, SqlConf.Port, SqlConf.DBName)
	//dsn := "root:123456@tcp(192.168.137.200:3306)/login_server"
	SqlDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Println("dsn=",dsn)
	if err != nil {
		fmt.Println("mysql connect err=", err)
		return
	}

}
