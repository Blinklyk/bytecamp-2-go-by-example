package repository

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
	"sync"
)

type Topic struct {
	Id         int64  `json:"id"`
	Title      string `form:"title" json:"title"`
	Content    string `form:"content" json:"content"`
	CreateTime int64  `json:"create_time"`
}
type TopicDao struct {
}

var (
	topicDao  *TopicDao
	topicOnce sync.Once
)

func NewTopicDaoInstance() *TopicDao {
	topicOnce.Do(
		func() {
			topicDao = &TopicDao{}
		})
	return topicDao
}
func (*TopicDao) QueryTopicById(id int64) *Topic {
	// 改了topicIndexMap的类型，其他业务中用到了也需要修改
	topicIndexMap.RLock()
	defer topicIndexMap.RUnlock()
	return topicIndexMap.data[id]
}

func (*TopicDao) SaveTopic(topic *Topic) error {
	// 最佳模式打开文件
	file, err := os.OpenFile("./data/topic", os.O_WRONLY|os.O_APPEND, 0777)
	if err != nil {
		return err
	}
	defer file.Close()

	// 打开流
	writer := bufio.NewWriter(file)

	// json 转换
	bytes, err := json.Marshal(*topic)
	if err != nil {
		return err
	}
	// 写入啥
	_, err = writer.WriteString("\n" + string(bytes))
	log.Println(string(bytes))
	if err != nil {
		return err
	}
	// 开写！
	err = writer.Flush()
	log.Println("写入成功， 内容为" + string(bytes))
	if err != nil {
		return err
	}

	// 更新索引
	// 1.重新加载索引 缺点：浪费资源
	// initxxx()
	// 2.更新索引 缺点：不同存储介质的数不一致性（热点问题） and 并发写入索引
	// 解决方案：对写操作加锁
	topicIndexMap.Lock()
	defer topicIndexMap.Unlock()
	topicIndexMap.data[topic.Id] = topic
	return nil
}
