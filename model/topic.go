package model

type Account struct {
	AccountName string `form:"account_name"`
	PassWord string `form:"password"`
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