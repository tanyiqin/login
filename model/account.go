package model

import "time"

type AccountInfo struct {
	AccountID	int64	`form:"account_id" gorm:"primaryKey" json:"account_id"`
	AccountName string `form:"account_name" binding:"required" json:"account_name"`
	SdkType	int `form:"sdk_type" json:"sdk_type"`
	PassWord string `form:"password" gorm:"Column:password" json:"password"`
	Salt	string `gorm:"salt" json:"-"`
	CreateTime	time.Time `gorm:"autoCreateTime" json:"-"`
	UpdateTime	time.Time `gorm:"autoUpdateTime" json:"-"`
}