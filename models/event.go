package models

import (
	"fmt"

	"github.com/couchbase/gocb"
)

type Event struct {
	EventType string
	Payload   interface{}
}

type Document struct {
	Store []Event
}

func UpsertDocument(id string, document Document) (gocb.Cas, error) {
	cas, err := EventBucket.Upsert(id, document, 0)

	if err != nil {
		fmt.Println("Couldn't Upsert document")
		fmt.Println(err)

		return cas, err
	}

	return cas, nil
}

func GetDocument(id string) (Document, gocb.Cas, error) {
	var document Document

	cas, err := EventBucket.Get(id, &document)

	if err != nil {
		fmt.Println("Couldn't Get document")
		fmt.Println(err)

		return document, cas, err
	}

	return document, cas, nil
}
