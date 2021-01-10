package dao

import (
	"github.com/gin-gonic/gin"
	"login/model"
	"net/http"
)

func AccountCreate(c *gin.Context) {
	account := model.Account{}
	err := c.Bind(&account)
	if err != nil {
		c.String(http.StatusBadRequest, "create account err = %s", err.Error())
	} else {

	}
}

//func MustLogin() gin.HandlerFunc{ // 必须登录
//	return func(c *gin.Context) {
//		if _, status := c.GetQuery("token"); !status {
//			c.String(http.StatusUnauthorized, "queshao token")
//			c.Abort()
//		}
//		// c.Next()
//	}
//}
//
//func GetTopicDetail(c *gin.Context) {
//	c.JSON(200, model.CreateTopic(101,"tizibiaoti"))
//}
//
//func GetTopicList(c *gin.Context) {
//	query:=model.TopicQuery{}
//	err := c.BindQuery(&query)
//	if err != nil {
//		c.String(400, "param error%s", err.Error())
//	} else {
//		c.JSON(200, query)
//	}
//}
//
//func NewTopic(c *gin.Context) {
//	Topic := model.Topic{}
//	err := c.BindJSON(&Topic)
//	if err != nil {
//		c.String(400, "param error%s", err.Error())
//	} else {
//		c.JSON(200, Topic)
//	}
//}
//
//func DeleteTopic(c *gin.Context) {
//	c.String(200, "shanchu")
//}