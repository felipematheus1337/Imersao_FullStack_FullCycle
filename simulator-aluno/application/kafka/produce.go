package kafka

import (
	"encoding/json"
	route2 "github.com/felipematheus1337/Imersao_FullStack_FullCycle/application/route"
	"github.com/felipematheus1337/Imersao_FullStack_FullCycle/infra/kafka"
	"log"
	"os"
	"time"
)

func Produce(msg *ckafka.Message) {
	producer := kafka.NewKafkaProducer()
	route := route2.NewRoute()
	json.Unmarshal(msg.Value, &route)
	route.LoadPositions()
	positions, err := route.ExportJsonPositions()
	if err != nil {
		log.Println(err.Error())
	}
	for _, p := range positions {
		kafka.Publish(p, os.Getenv("KafkaProduceTopic"), producer)
		time.Sleep(time.Millisecond * 500)
	}
}
