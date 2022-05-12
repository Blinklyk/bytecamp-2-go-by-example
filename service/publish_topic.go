package service

import (
	"github.com/Moonlight-Zhao/go-project-example/repository"
	"github.com/Moonlight-Zhao/go-project-example/utils"
	"time"
)

type TopicData struct {
}

var node, _ = utils.NewWorker(1)

func SaveTopic(topic *repository.Topic) error {
	// 雪花生成ID, 填充
	topic.Id = node.GetId()
	// 生成创建时间，填充
	//topic.CreateTime = time.Now().Unix()
	topic.CreateTime = repository.UnixTime(time.Now())
	err := repository.NewTopicDaoInstance().SaveTopic(topic)
	if err != nil {
		return err
	}
	return nil
}
