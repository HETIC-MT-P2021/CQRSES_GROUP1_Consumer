package controllers

import (
	"fmt"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP1_Consumer/models"
)

func AddPostEvent(id string, event models.Event) {
	eventStore, _, getError := models.GetEventStore(id)

	if getError != nil && fmt.Sprintf("%e", getError) != "key not found" {
		fmt.Println(getError)
		return
	}

	eventStore = models.AddEventToDocument(eventStore, event)

	_, upsertError := models.UpsertDocument(id, eventStore)

	if upsertError != nil {
		fmt.Println(upsertError)
		return
	}

	fmt.Println("Post successfuly stored !")

	readModel, err := models.BuildPostReadModel(id, eventStore)
	fmt.Println(readModel)
	if err != nil {
		// remove event from eventStore
		return
	}

	models.UpsertReadModel(id, readModel)

	fmt.Println("Read models successfully updated")
}
