package controllers

import (
	"fmt"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP1_Consumer/consts"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP1_Consumer/models"
)

func AddPostEvent(id string, event models.Event) {
	eventStore, _, getError := models.GetEventStore(id)

	if getError != nil && fmt.Sprintf("%s", getError) != "key not found" {
		fmt.Println(getError)
		return
	}

	eventStore, didUpdate := models.AddEventToEventStore(eventStore, event)

	if didUpdate == false {
		if event.EventType == consts.POST_CREATED_EVENT_TYPE {
			fmt.Println("This post already exist")
		} else {
			fmt.Println("This post is already deleted or doesn't exist")
		}

		return
	}

	readModel, err := models.BuildPostReadModel(id, eventStore)

	if err != nil {
		fmt.Println(err)
		return
	}

	_, upsertError := models.UpsertEventStore(id, eventStore)

	if upsertError != nil {
		fmt.Println("Error, Couldn't insert event store:")
		fmt.Println(upsertError)
		return
	}

	fmt.Println("Event successfuly stored !")

	if event.EventType == consts.POST_DELETED_EVENT_TYPE {
		_, deleteError := models.DeleteReadModel(id)

		if deleteError != nil {
			fmt.Println("Error, Couldn't delete read model:")
			fmt.Println(upsertError)
			return
		}
	} else {
		_, upsertError = models.UpsertReadModel(id, readModel)

		if upsertError != nil {
			fmt.Println("Error, Couldn't insert read model:")
			fmt.Println(upsertError)
			return
		}
	}

	fmt.Println("Read models successfully updated")
}
