# points-simulator

Golang service that provides fake route points via kafka.

## Connection test with kafka

### Go producer test

Within kafka, wait for a message in a test topic
```bash
kafka-console-consumer --bootstrap-server=localhost:9092 --topic=readtest
```

On a Go function
```golang
producer := kafka.NewKafkaProducer()
kafka.Publish("My Message", "readtest", producer)
```

### Go consumer test
```bash
kafka-console-producer --bootstrap-server=localhost:9092 --topic=readtest
```

On a Go function
```golang
msgChan := make(chan *ckafka.Message)
consumer := kafka.NewKafkaConsumer(msgChan)

go consumer.Consume()

for msg := range msgChan {
	fmt.Println(string(msg.Value))
}
```

