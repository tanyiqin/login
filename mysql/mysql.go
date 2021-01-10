package mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var err error
var SqlDB *gorm.DB

func init() {
	dsn := "root:123456@tcp(192.168.137.200:3306)/"
	gorm.Open()
}
