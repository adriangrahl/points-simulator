# points-simulator

Golang service that provides fake route points via kafka.

## Connection test with kafka

Within kafka, wait for a message in a test topic
```bash
	kafka-console-consumer --bootstrap-server=localhost:9092 --topic=readtest
```


```golang
	producer := kafka.NewKafkaProducer()
	kafka.Publish("My Message", "readtest", producer)
```golang



