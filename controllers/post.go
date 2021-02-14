package controllers

import (
	"fmt"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP1_Consumer/models"
)

func AddPostEvent(id string, event models.Event) {
	document, _, getError := models.GetDocument(id)

	if getError != nil {
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
	/*
		readModel := models.BuildPostReadModel(id, document)

		models.UpsertReadModel(id, readModel)
	*/
}
