package lib

import (
	"crypto"
	"crypto/md5"
	"encoding/hex"
	"login/model"
	"login/mysql"
)

const (
	SdkTypeDefault = iota
)

// 获取账号信息
func GetAccountByName(AccountName string, SdkType int) (model.AccountInfo, error){
	accountInfo := model.AccountInfo{}
	result :=mysql.SqlDB.Table("account_info").Where("account_name = ? AND sdk_type = ?", AccountName, SdkType).First(&accountInfo)
	return accountInfo, result.Error
}
func GetAccountByID(AccountID int) (model.AccountInfo, error) {
	accountInfo := model.AccountInfo{}
	result := mysql.SqlDB.Table("account_info").Where("account_id = ?", AccountID).First(&accountInfo)
	return accountInfo, result.Error
}

// 创建账号
func CreateAccount(AccountName string, Password string) error {
	accountInfo := model.AccountInfo{}
	accountInfo.AccountName = AccountName
	md5.New()
	hex.Encode()
}