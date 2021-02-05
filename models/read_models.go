package models

import (
	"fmt"

	"github.com/couchbase/gocb"
)

func updateModel(post Post, event Event) Post {

}

func BuildPostReadModel(id string, document Document) Post {
	var readmodel Post

	for _, event := range document.Store {
		readmodel = updateModel(readmodel, event)
	}

	return readmodel
}

func UpsertReadModel(id string, model interface{}) (gocb.Cas, error) {
	cas, err := ReadBucket.Upsert(id, model, 0)

	if err != nil {
		fmt.Println("Couldn't Upsert readmodel")
		fmt.Println(err)

		return cas, err
	}

	return cas, nil
}
