package main

import (
	"fmt"
	route2 "github.com/felipematheus1337/Imersao_FullStack_FullCycle/application/route"
	"github.com/felipematheus1337/Imersao_FullStack_FullCycle/infra/kafka"
	"github.com/joho/godotenv"
	"log"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	msgChannel := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()

	for msg := range msgChannel {
		fmt.Println(string(msg.Value))
	}

	producer := kafka.NewKafkaProducer()
	kafka.Publish("ola", "readtest", producer)

	route := route2.Route{
		ID:       "1",
		ClientID: "1",
	}

	route.LoadPositions()
	stringJson, _ := route.ExportJsonPositions()
	lastIndex := len(stringJson) - 1
	fmt.Println(stringJson[lastIndex])
}
