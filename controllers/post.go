package controllers

import (
	"fmt"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP1_Consumer/models"
)

func AddPostEvent(id string, event models.Event) {
	eventStore, _, getError := models.GetEventStore(id)

	if getError != nil && fmt.Sprintf("%s", getError) != "key not found" {
		fmt.Println(getError)
		return
	}

	eventStore = models.AddEventToDocument(eventStore, event)

	readModel, err := models.BuildPostReadModel(id, eventStore)

	if err != nil {
		fmt.Println(err)
		return
	}

	_, upsertError := models.UpsertDocument(id, eventStore)

	if upsertError != nil {
		fmt.Println("Error, Couldn't insert event store:")
		fmt.Println(upsertError)
		return
	}

	fmt.Println("Post successfuly stored !")

	_, upsertError = models.UpsertReadModel(id, readModel)

	if upsertError != nil {
		fmt.Println("Error, Couldn't insert read model:")
		fmt.Println(upsertError)
		return
	}

	fmt.Println("Read models successfully updated")
}
