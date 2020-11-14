package main

import "ServiceMesh/GoScraperService/kafka/subscriber"

func main() {
	kafkaSubscriberConfig := subscriber.ConfigureKafkaSubscriber()
	subscriber.CreateKafkaSubscriber(kafkaSubscriberConfig)
}