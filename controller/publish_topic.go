package controller

import (
	"github.com/Moonlight-Zhao/go-project-example/repository"
	"github.com/Moonlight-Zhao/go-project-example/service"
)

func SaveTopic(topic *repository.Topic) *PageData {
	service.SaveTopic(topic)
	if topic.Id <= 0 {
		return &PageData{
			Code: -1,
			Msg:  "id错误",
		}
	}
	return &PageData{
		Code: 0,
		Msg:  "success",
		Data: topic.Id,
	}
}
