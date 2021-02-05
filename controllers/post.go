package controllers

import (
	"github.com/HETIC-MT-P2021/CQRSES_GROUP1_Consumer/models"
)

func AddPostEvent(id string, event models.Event) {
	document, _, getError := models.GetDocument(id)

	if getError != nil {
		return
	}

	_, upsertError := models.UpsertDocument(id, document)

	if upsertError != nil {
		return
	}

	readModel := models.BuildPostReadModel(id, document)

	models.UpsertReadModel(id, readModel)
}
