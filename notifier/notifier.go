package notifier

import (
	"github.com/streadway/amqp"
	"log"
)

var (
	conn *amqp.Connection
	NotificationChan chan string
)

func init() {
	NotificationChan = make(chan string)

	var err error
	conn, err = amqp.Dial("amqp://rabbitmq:rabbitmq@localhost:5672/")
	if err != nil {
		log.Fatal("can't connect to rabbitmq (", err, ")")
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("can't create rabbitmq channel (", err, ")")
	}

	q, err := ch.QueueDeclare("notifications", false, false, false, false, nil)
	if err != nil {
		log.Fatal("can't connect to notification queue (", err, ")")
	}

	go func() {
		for {
			dch, err := ch.Consume(q.Name, "", false, false, false, false, nil)
			if err != nil {
				log.Fatal(err)
			}

			for d := range dch {
				NotificationChan <- string(d.Body)
				d.Ack(false)
			}
		}
	}()
}
