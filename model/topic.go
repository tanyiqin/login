package model

import "time"

type AccountInfo struct {
	AccountID	int64	`form:"account_id" gorm:"account_id" json:"account_id"`
	AccountName string `form:"account_name" binding:"required" gorm:"account_name" json:"account_name"`
	SdkType	int `form:"sdk_type" gorm:"sdk_type" json:"sdk_type"`
	PassWord string `form:"password" gorm:"password" json:"-"`
	Salt	string `gorm:"salt" json:"-"`
	CreateTime	time.Time `gorm:"create_time,autoCreateTime" json:"-"`
	UpdateTime	time.Time `gorm:"update_time,autoUpdateTime" json:"-"`
}

//type Topic struct {
//	TopicID int `json:"id"`
//	TopicTitle string `json:"title" binding:"min=4,max=20"`
//	TopicShortTitle string `json:"stitle" binding:"required,nefield=TopicTitle"`
//	UserIP string `json:"ip" binding:"ipv4"`
//	TopicScore int `json:"score" binding:"omitempty,gt=5"`
//	TopicUrl string `json:"url" binding:"omitempty,topicurl"`
//}
//
//func CreateTopic(id int, title string) Topic{
//	return Topic{TopicID: id, TopicTitle: title}
//}
//
//type TopicQuery struct {
//	UserName string `json:"username" form:"username"`
//	Page int `json:"page" form:"page" binding:"required"`
//	PageSize int `json:"pagesize"`
//}