package dao

import (
	"github.com/gin-gonic/gin"
	"login/lib"
	"login/model"
	"net/http"
	"reflect"
	"strconv"
)

func AccountCreate(c *gin.Context) {
	account := model.AccountInfo{}
	err := c.Bind(&account)
	if err != nil {
		c.String(http.StatusNotFound, "create account err = %s", err.Error())
	} else {
		if err := lib.CreateAccount(account.AccountName, account.PassWord); err != nil {
			c.String(http.StatusNotFound, "create account err = %s", err.Error())
		} else {
			c.String(http.StatusOK, "create ok")
		}
	}
}

func AccountLogin(c *gin.Context) {
	AccountName := c.DefaultQuery("account_name", "")
	Password := c.DefaultQuery("password","")
	SdkType, _ := strconv.Atoi(c.DefaultQuery("sdk_type", "0"))
	if accountInfo, err := lib.GetAccountByName(AccountName, SdkType); err != nil {
		c.String(http.StatusNotFound, "get account err = %s", err.Error())
	} else {
		salt := accountInfo.Salt
		if reflect.DeepEqual(lib.MD5(Password, []byte(salt)), accountInfo.PassWord) {
			c.JSON(http.StatusOK, gin.H{"account_id":accountInfo.AccountID})
		} else {
			c.String(http.StatusNotFound, "auth error")
		}
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