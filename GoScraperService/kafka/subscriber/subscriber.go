package subscriber

import (
	"fmt"
	"github.com/Shopify/sarama"
	"os"
	"os/signal"
	"time"
)

func brokerAddress() string {
	brokerAddr := os.Getenv("BROKER_ADDR")
	if len(brokerAddr) == 0 {
		brokerAddr = "localhost:9092"
	}
	return brokerAddr
}


func getTopic() string {
	topic := os.Getenv("TOPIC")
	if len(topic) == 0 {
		//topic = "default-topic"
		topic = "reflectoring-1"
	}
	return topic
}


func ConfigureKafkaSubscriber() *sarama.Config {
	fmt.Println("Configuring Kafka subscriber...")
	config := sarama.NewConfig()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Offsets.AutoCommit.Interval = 5 * time.Second
	config.Consumer.Return.Errors = true

	return config
}


func CreateKafkaSubscriber(config *sarama.Config) {
	fmt.Println("Starting Kafka subscriber...")
	// create a new consumer for given brokers and configuration
	brokers := []string{brokerAddress()}
	master, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := master.Close(); err != nil {
			panic(err)
		}
	}()

	// describe about the offset here: literal value, sarama.OffsetOldest, sarama.OffsetNewest
	// This is important in case of reconnection
	// create a Kafka partition consumer for given topic
	consumer, err := master.ConsumePartition(getTopic(), 0, sarama.OffsetOldest)
	if err != nil {
		panic(err)
	}

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)

	// Count how many messages processed
	msgCount := 0

	// Get signal for finish
	doneCh := make(chan struct{})
	go func() {
		for {
			// consume and process the messages
			select {
			case err := <- consumer.Errors():
				fmt.Println(err)
			case msg := <-consumer.Messages():
				msgCount ++
				fmt.Println("Received messages", string(msg.Key), string(msg.Value))
			case <-signals:
				fmt.Println("Interrupt is detected")
				doneCh <- struct{}{}


			}
		}
	}()

	<-doneCh
	fmt.Println("Processed", msgCount,"messages")
}

//func main() {
//	fmt.Println("Starting Kafka subscriber...")
//
//	// create the customer configuration
//	config := sarama.NewConfig()
//	config.Consumer.Offsets.Initial = sarama.OffsetOldest
//	config.Consumer.Offsets.AutoCommit.Interval = 5 * time.Second
//	config.Consumer.Return.Errors = true
//
//
//	// create a new consumer for given brokers and configuration
//	brokers := []string{"localhost:9092"}
//	master, err := sarama.NewConsumer(brokers, config)
//	if err != nil {
//		panic(err)
//	}
//
//	defer func() {
//		if err := master.Close(); err != nil {
//			panic(err)
//		}
//	}()
//
//	// describe about the offset here: literal value, sarama.OffsetOldest, sarama.OffsetNewest
//	// This is important in case of reconnection
//	// create a Kafka partition consumer for given topic
//	consumer, err := master.ConsumePartition("", 0, sarama.OffsetOldest)
//	if err != nil {
//		panic(err)
//	}
//
//	signals := make(chan os.Signal, 1)
//	signal.Notify(signals, os.Interrupt)
//
//	// Count how many messages processed
//	msgCount := 0
//
//	// Get signal for finish
//	doneCh := make(chan struct{})
//	go func() {
//		for {
//			// consume and process the messages
//			select {
//			case err := <- consumer.Errors():
//				fmt.Println(err)
//			case msg := <-consumer.Messages():
//				msgCount ++
//				fmt.Println("Received messages", string(msg.Key), string(msg.Value))
//			case <-signals:
//				fmt.Println("Interrupt is detected")
//				doneCh <- struct{}{}
//
//
//			}
//		}
//	}()
//
//	<-doneCh
//	fmt.Println("Processed", msgCount,"messages")
//
//}