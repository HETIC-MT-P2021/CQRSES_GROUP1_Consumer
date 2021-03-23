package models

import (
	"fmt"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP1_Consumer/consts"
	"github.com/couchbase/gocb"
)

type Event struct {
	EventType string
	Payload   Post
}

type EventStore struct {
	Store []Event
}

func AddEventToEventStore(eventStore EventStore, event Event) (EventStore, bool) {
	didUpdate := false

	if len(eventStore.Store) == 0 && event.EventType == consts.POST_CREATED_EVENT_TYPE {
		eventStore.Store = []Event{event}
		didUpdate = true
	} else if len(eventStore.Store) > 0 && eventStore.Store[len(eventStore.Store)-1].EventType != consts.POST_DELETED_EVENT_TYPE && event.EventType != consts.POST_CREATED_EVENT_TYPE {
		// No update for deleted posts and creation only for create event
		eventStore.Store = append(eventStore.Store, event)
		didUpdate = true
	}

	return eventStore, didUpdate
}

func UpsertEventStore(id string, eventStore EventStore) (gocb.Cas, error) {
	cas, err := EventBucket.Upsert(id, eventStore, 0)

	if err != nil {
		fmt.Println("Couldn't Upsert eventStore")
		fmt.Println(err)

		return cas, err
	}

	return cas, nil
}

func GetEventStore(id string) (EventStore, gocb.Cas, error) {
	var eventStore EventStore

	cas, err := EventBucket.Get(id, &eventStore)

	if err != nil {
		fmt.Println("Couldn't Get eventStore")
		fmt.Println(err)

		return eventStore, cas, err
	}

	return eventStore, cas, nil
}
