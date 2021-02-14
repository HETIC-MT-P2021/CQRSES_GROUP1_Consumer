package models

import (
	"fmt"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP1_Consumer/consts"
	"github.com/couchbase/gocb"
	"github.com/jinzhu/copier"
)

func updateModel(post Post, event Event) Post {
	copier.Copy(post, event.Payload)

	return post
}

func BuildPostReadModel(id string, document Document) Post {
	var readmodel Post

	for _, event := range document.Store {
		if event.EventType == consts.POST_CREATED_EVENT_TYPE || event.EventType == consts.POST_UPDATED_EVENT_TYPE {
			readmodel = updateModel(readmodel, event)
		}
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
