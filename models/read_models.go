package models

import (
	"errors"
	"fmt"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP1_Consumer/consts"
	"github.com/couchbase/gocb"
	"github.com/jinzhu/copier"
)

func updateModel(post *Post, payload Post) (*Post, error) {
	err := copier.Copy(&post, &payload)

	return post, err
}

func BuildPostReadModel(id string, eventStore EventStore) (Post, error) {
	var readmodel Post

	for _, event := range eventStore.Store {
		if event.EventType != consts.POST_DELETED_EVENT_TYPE {
			if readmodel.ID != 0 && event.EventType == consts.POST_CREATED_EVENT_TYPE {
				return readmodel, errors.New("There cannot be two createPost Event")
			}
			readmodel, err := updateModel(&readmodel, event.Payload)

			if err != nil {
				return *readmodel, err
			}
		} else {
			return readmodel, nil
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

func GetReadModel(id string) (Post, gocb.Cas, error) {
	var post Post
	cas, error := ReadBucket.Get(id, &post)

	return post, cas, error
}

func DeleteReadModel(id string) (gocb.Cas, error) {
	_, actualCas, getErr := GetReadModel(id)

	if getErr != nil {
		return actualCas, nil
	}

	return ReadBucket.Remove(id, actualCas)
}
