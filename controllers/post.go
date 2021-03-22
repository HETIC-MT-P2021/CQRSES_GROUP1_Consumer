package controllers

import (
	"fmt"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP1_Consumer/models"
)

func AddPostEvent(id string, event models.Event) {
	document, _, getError := models.GetDocument(id)

	if getError != nil && fmt.Sprintf("%e", getError) != "key not found" {
		fmt.Println(getError)
		return
	}

	document = models.AddEventToDocument(document, event)

	_, upsertError := models.UpsertDocument(id, document)

	if upsertError != nil {
		fmt.Println(upsertError)
		return
	}

	fmt.Println("Post successfuly stored !")

	readModel, err := models.BuildPostReadModel(id, document)

	if err != nil {
		// remove event from document
		return
	}

	models.UpsertReadModel(id, readModel)

	fmt.Println("Read models successfully updated")
}
