package models

import (
	"fmt"

	"github.com/couchbase/gocb"
)

type Event struct {
	EventType string
	Payload   interface{}
}

type EventStore struct {
	Store []Event
}

func AddEventToDocument(eventStore EventStore, event Event) EventStore {
	if len(eventStore.Store) == 0 {
		eventStore.Store = []Event{event}
	} else {
		eventStore.Store = append(eventStore.Store, event)
	}

	return eventStore
}

func UpsertDocument(id string, eventStore EventStore) (gocb.Cas, error) {
	cas, err := EventBucket.Upsert(id, eventStore, 0)

	if err != nil {
		fmt.Println("Couldn't Upsert eventStore")
		fmt.Println(err)

		return cas, err
	}

	return cas, nil
}

func GetDocument(id string) (EventStore, gocb.Cas, error) {
	var eventStore EventStore

	cas, err := EventBucket.Get(id, &eventStore)

	if err != nil {
		fmt.Println("Couldn't Get eventStore")
		fmt.Println(err)

		return eventStore, cas, err
	}

	return eventStore, cas, nil
}
