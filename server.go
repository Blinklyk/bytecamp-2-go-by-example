package main

import (
	"github.com/Moonlight-Zhao/go-project-example/controller"
	"github.com/Moonlight-Zhao/go-project-example/repository"
	"gopkg.in/gin-gonic/gin.v1"
	"log"
	"net/http"
	"os"
)

func main() {
	if err := Init("./data/"); err != nil {
		os.Exit(-1)
	}
	r := gin.Default()
	r.GET("/community/page/get/:id", func(c *gin.Context) {
		topicId := c.Param("id")
		data := controller.QueryPageInfo(topicId)
		c.JSON(200, data)
	})
	r.POST("/community/topic/post", func(c *gin.Context) {
		var topic repository.Topic
		if err := c.ShouldBind(&topic); err == nil {
			//titleData, ok := c.GetPostForm("title")
			//if ok {
			//	topic.Title = titleData
			//}
			//contentData, ok := c.GetPostForm("content")
			//if ok {
			//	topic.Content = contentData
			//}
			log.Println("title : " + topic.Title)
			log.Println("content : " + topic.Content)
			data := controller.SaveTopic(&topic)
			c.JSON(http.StatusOK, data)
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	log.Fatal(r.Run(":9999"))
}

func Init(filePath string) error {
	if err := repository.Init(filePath); err != nil {
		return err
	}
	return nil
}
