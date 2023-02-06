package main

import (
	"fmt"
	kafka2 "github.com/felipematheus1337/Imersao_FullStack_FullCycle/application/kafka"
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
		go kafka2.Produce(msg)

	}

}
