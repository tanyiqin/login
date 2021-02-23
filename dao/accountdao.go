package dao

import (
	"context"
	"github.com/gin-gonic/gin"
	"login/lib"
	"login/model"
	"login/redis"
	"net/http"
	"reflect"
	"strconv"
	"time"
)

const (
	tokenTTL = 5 * time.Minute
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
			token := lib.GetRandomString(16)
			status := redis.RedisDB.Set(context.Background(), AccountName + "|" + strconv.Itoa(SdkType), token, tokenTTL)
			if status.Err() != nil {
				c.String(http.StatusNotFound, "gene token error =", err)
			} else {
				c.JSON(http.StatusOK, gin.H{"account_id":accountInfo.AccountID, "token":token})
			}
		} else {
			c.String(http.StatusNotFound, "auth error")
		}
	}
}

func Test(c *gin.Context) {
	panic("test")
}