package main

import (
	"github.com/adriangrahl/points-simulator/infra/kafka"
	"github.com/joho/godotenv"
	"log"
	"fmt"
	ckafka "github.com/confluentinc/confluent-kafka-go/kafka"
	producer "github.com/adriangrahl/points-simulator/application/kafka"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}
}

func main() {
	msgChan := make(chan *ckafka.Message)
	consumer := kafka.NewKafkaConsumer(msgChan)
	go consumer.Consume()
	for msg := range msgChan {
		fmt.Println(string(msg.Value))
		go  producer.Produce(msg)
	}

	////////////////consumer test///////////////
	// msgChan := make(chan *ckafka.Message)
	// consumer := kafka.NewKafkaConsumer(msgChan)
	// go consumer.Consume()
	// for msg := range msgChan {
	// 	fmt.Println(string(msg.Value))
	// }


	////////////////producer test///////////////
	// producer := kafka.NewKafkaProducer()
	// kafka.Publish("ola", "readtest", producer)

	for {
		_ = 1
	}

	////////////////get route point test///////////////
	// "fmt"
	// route2 "github.com/adriangrahl/points-simulator/application/route"

	// route := route2.Route{
	// 	ID: "1",
	// 	ClientID: "1",
	// }

	// route.LoadPositions()
	// stringjson, _ := route.ExportJsonPositions()
	// fmt.Println(stringjson[0])
}

