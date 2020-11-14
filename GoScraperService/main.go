package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

type Item struct {
	Id string `json:"id"`
	Name string `json:"name"`

}

func main() {
	fmt.Println("Starting a RabbitMQ consumer...")

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		"spring-boot-queue",
		false,
		false,
		false,
		false,
		nil)

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil)
	fmt.Println(msgs)
	//forever := make(chan bool)
	//item := Item{}
	//go func() {
	//	for d := range msgs {
	//		log.Printf("Received a message: %s", d.Body)
	//
	//		err := json.Unmarshal(d.Body,&item)
	//		if err != nil {
	//			log.Fatalf(err.Error())
	//		}
	//		log.Println(item.Id)
	//		log.Println(item.Name)
	//
	//		err = ch.Publish(
	//			"",
	//			q.Name,
	//			false,
	//			false,
	//			amqp.Publishing{
	//				Headers:         nil,
	//				ContentType:     "text/plain",
	//				ContentEncoding: "",
	//				DeliveryMode:    0,
	//				Priority:        0,
	//				CorrelationId:   "",
	//				ReplyTo:         "",
	//				Expiration:      "",
	//				MessageId:       "",
	//				Timestamp:       time.Time{},
	//				Type:            "",
	//				UserId:          "",
	//				AppId:           "",
	//				Body:            []byte(fmt.Sprintf("Got it! Item %s was added to shopping cart!", d.Body)),
	//			})
	//
	//		if err != nil {
	//			log.Fatalf(err.Error())
	//		}
	//	}
	//}()
	//
	//<-forever
}