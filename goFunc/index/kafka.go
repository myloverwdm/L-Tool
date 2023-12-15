package index

import (
	"LTool/goFunc"
	"encoding/json"
	"fmt"
	"github.com/IBM/sarama"
	"log"
)

type TopicAndPartition struct {
	TopicName string `json:"topicName"`

	Partitions []int `json:"partitions"`
}

type KafkaForm struct {
	Connection string `json:"connection"`
}

const kafkaForm = "kafkaForm"

func GetKafkaForm() KafkaForm {
	var kafka KafkaForm
	kafka = goFunc.Get(kafkaForm, kafka)
	return kafka
}

func UpdateKafkaForm(kafka KafkaForm) string {
	marshal, err := json.Marshal(kafka)
	if err == nil {
		goFunc.Set(kafkaForm, string(marshal))
		return ""
	}
	return err.Error()
}

func GetTopicAndPartition() []TopicAndPartition {

	return []TopicAndPartition{
		TopicAndPartition{"ss", []int{1}},

		TopicAndPartition{"lyb", []int{1, 2, 3, 4}},

		TopicAndPartition{"Lyb", []int{1, 2, 3, 4, 5, 6, 7, 8}},

		TopicAndPartition{"qewesreb", []int{1}},

		TopicAndPartition{"urtygcseb", []int{1, 2, 3, 4, 5, 6, 7, 8}},
	}
}

func main() {
	// 创建Kafka配置
	config := sarama.NewConfig()
	config.Version = sarama.V2_8_0_0 // 设置Kafka版本
	// 创建Kafka客户端
	client, err := sarama.NewClient([]string{"10.100.2.191:6667,10.100.2.192:6667,10.100.2.193:6667"}, config)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
	// 获取所有主题
	topics, err := client.Topics()
	if err != nil {
		log.Fatal(err)
	}
	// 打印所有主题
	// fmt.Println("Topics:")
	for _, topic := range topics {
		fmt.Println(topic)
	}
}
