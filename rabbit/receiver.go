package rabbit

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP1_Consumer/consts"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP1_Consumer/controllers"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP1_Consumer/models"
)

// ReceiveCreatePost gets consume message from a queue
func ReceiveCreatePost() {
	msgs, err := CommandChannel.Consume(
		consts.CREATE_POST_COMMAND_QUEUE, // queue
		"",                               // consumer
		true,                             // auto-ack
		false,                            // exclusive
		false,                            // no-local
		false,                            // no-wait
		nil,                              // args
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			post := models.Post{}

			err = json.Unmarshal(d.Body, &post)

			if err != nil {
				log.Println("Couldn't get post from body")
				log.Println(err)
			}

			createPostEvent := models.Event{
				EventType: consts.POST_CREATED_EVENT_TYPE,
				Payload:   post,
			}

			controllers.AddPostEvent(fmt.Sprintf("%d", post.ID), createPostEvent)
		}
	}()

	log.Printf("Create Queue Waiting for messages. To exit press CTRL+C")
}

// ReceiveUpdatePost gets consume message from a queue
func ReceiveUpdatePost() {
	msgs, err := CommandChannel.Consume(
		consts.UPDATE_POST_COMMAND_QUEUE, // queue
		"",                               // consumer
		true,                             // auto-ack
		false,                            // exclusive
		false,                            // no-local
		false,                            // no-wait
		nil,                              // args
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			post := models.Post{}

			err = json.Unmarshal(d.Body, &post)

			if err != nil {
				log.Println("Couldn't get post from body")
				log.Println(err)
			}

			updatePostEvent := models.Event{
				EventType: consts.POST_UPDATED_EVENT_TYPE,
				Payload:   post,
			}

			controllers.AddPostEvent(fmt.Sprintf("%d", post.ID), updatePostEvent)
		}
	}()

	log.Printf("Update Queue Waiting for messages. To exit press CTRL+C")
}

// ReceiveDeletePost gets consume message from a queue
func ReceiveDeletePost() {
	msgs, err := CommandChannel.Consume(
		consts.DELETE_POST_COMMAND_QUEUE, // queue
		"",                               // consumer
		true,                             // auto-ack
		false,                            // exclusive
		false,                            // no-local
		false,                            // no-wait
		nil,                              // args
	)
	failOnError(err, "Failed to register a consumer")

	go func() {
		for d := range msgs {
			post := models.Post{}

			err = json.Unmarshal(d.Body, &post)

			if err != nil {
				log.Println("Couldn't get post from body")
				log.Println(err)
			}

			updatePostEvent := models.Event{
				EventType: consts.POST_DELETED_EVENT_TYPE,
				Payload:   post,
			}

			controllers.AddPostEvent(fmt.Sprintf("%d", post.ID), updatePostEvent)
		}
	}()

	log.Printf("Delete Queue Waiting for messages. To exit press CTRL+C")
}
