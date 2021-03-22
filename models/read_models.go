package models

import (
	"fmt"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP1_Consumer/consts"
	"github.com/couchbase/gocb"
	"github.com/jinzhu/copier"
)

func updateModel(post Post, event Event) Post {
	err := copier.Copy(post, event.Payload)

	return post, err
}

func BuildPostReadModel(id string, document Document) (Post, err) {
	var readmodel Post

	for _, event := range document.Store {
		if event.EventType == consts.POST_CREATED_EVENT_TYPE || event.EventType == consts.POST_UPDATED_EVENT_TYPE {
			readmodel, err = updateModel(readmodel, event)

			if err != nil {
				return readmodel, err
			}
		}
	}

	return readmodel, nil
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
