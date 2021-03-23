package rabbit

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP1_Consumer/consts"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP1_Consumer/controllers"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP1_Consumer/models"
)

// Receive gets consume message from a queue
func Receive() {
	msgs, err := CommandChannel.Consume(
		CreatePostQueue.Name, // queue
		"",                   // consumer
		true,                 // auto-ack
		false,                // exclusive
		false,                // no-local
		false,                // no-wait
		nil,                  // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			post := models.Post{}
			log.Printf("Received a message: %s", d.Body)

			err = json.Unmarshal(d.Body, &post)

			if err != nil {
				log.Println("Couldn't get post from body")
				log.Println(err)
			}

			createPostEvent := models.Event{
				EventType: consts.POST_CREATED_EVENT_TYPE,
				Payload:   post,
			}

			controllers.AddPostEvent(fmt.Sprintf("%s", post.ID), createPostEvent)
		}
	}()

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
