package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	//if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	//	v.RegisterValidation("topicurl", TopicUrl)
	//}
	//
	//v1 := router.Group("/v1/topics")
	//{
	//	v1.GET("", dao.GetTopicList)
	//	v1.GET("/:topic_id", dao.GetTopicDetail)
	//
	//	v1.Use(dao.MustLogin())
	//	{
	//		v1.POST("", dao.NewTopic)
	//		v1.DELETE("/:topic_id", dao.DeleteTopic)
	//	}
	//}
	accountGroup := router.Group("/account")
	{
		accountGroup.POST("/create", dao)
	}
	router.Run()
}
