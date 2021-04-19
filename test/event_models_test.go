package test

import (
	"testing"
	"time"

	"github.com/HETIC-MT-P2021/CQRSES_GROUP1_Consumer/consts"
	"github.com/HETIC-MT-P2021/CQRSES_GROUP1_Consumer/models"
)

func TestAddCreateEventToEventStore(t *testing.T) {
	var eventStore models.EventStore

	post := models.Post{
		ID:        0,
		Title:     "Test",
		Content:   "Test",
		AuthorID:  0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	createPostEvent := models.Event{
		EventType: consts.POST_CREATED_EVENT_TYPE,
		Payload:   post,
	}

	eventStore, didUpdate := models.AddEventToEventStore(eventStore, createPostEvent)

	if didUpdate != true {
		t.Errorf("AddEventToEventStore - CreateEvent, didUpdate should be true")
	}
	if len(eventStore.Store) != 1 {
		t.Errorf("AddEventToEventStore - CreateEvent, Store length should be 1")
	}

	eventStore, didUpdate = models.AddEventToEventStore(eventStore, createPostEvent)

	if didUpdate != false {
		t.Errorf("AddEventToEventStore - CreateEvent, didUpdate should be false")
	}
	if len(eventStore.Store) != 1 {
		t.Errorf("AddEventToEventStore - CreateEvent, Store length should be 1")
	}
}

func TestAddUpdateEventToEventStore(t *testing.T) {
	var eventStore models.EventStore

	post := models.Post{
		ID:        0,
		Title:     "Test",
		Content:   "Test",
		AuthorID:  0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	updatePostEvent := models.Event{
		EventType: consts.POST_UPDATED_EVENT_TYPE,
		Payload:   post,
	}

	eventStore, didUpdate := models.AddEventToEventStore(eventStore, updatePostEvent)

	if didUpdate != false {
		t.Errorf("AddEventToEventStore - UpdateEvent, didUpdate should be false")
	}
	if len(eventStore.Store) != 0 {
		t.Errorf("AddEventToEventStore - UpdateEvent, Store length should be 0")
	}

	createPostEvent := models.Event{
		EventType: consts.POST_CREATED_EVENT_TYPE,
		Payload:   post,
	}

	eventStore, didUpdate = models.AddEventToEventStore(eventStore, createPostEvent)
	eventStore, didUpdate = models.AddEventToEventStore(eventStore, updatePostEvent)

	if didUpdate != true {
		t.Errorf("AddEventToEventStore - UpdateEvent, didUpdate should be true")
	}
	if len(eventStore.Store) != 2 {
		t.Errorf("AddEventToEventStore - UpdateEvent, Store length should be 2")
	}

	eventStore, didUpdate = models.AddEventToEventStore(eventStore, updatePostEvent)

	if didUpdate != true {
		t.Errorf("AddEventToEventStore - UpdateEvent, didUpdate should be true")
	}
	if len(eventStore.Store) != 3 {
		t.Errorf("AddEventToEventStore - UpdateEvent, Store length should be 3")
	}
}

func TestAddDeleteEventToEventStore(t *testing.T) {
	var eventStore models.EventStore

	post := models.Post{
		ID:        0,
		Title:     "Test",
		Content:   "Test",
		AuthorID:  0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	deletePostEvent := models.Event{
		EventType: consts.POST_DELETED_EVENT_TYPE,
		Payload:   post,
	}

	eventStore, didUpdate := models.AddEventToEventStore(eventStore, deletePostEvent)

	if didUpdate != false {
		t.Errorf("AddEventToEventStore - DeleteEvent 1, didUpdate should be false")
	}
	if len(eventStore.Store) != 0 {
		t.Errorf("AddEventToEventStore - DeleteEvent 1, Store length should be 0")
	}

	createPostEvent := models.Event{
		EventType: consts.POST_CREATED_EVENT_TYPE,
		Payload:   post,
	}

	eventStore, didUpdate = models.AddEventToEventStore(eventStore, createPostEvent)
	eventStore, didUpdate = models.AddEventToEventStore(eventStore, deletePostEvent)

	if didUpdate != true {
		t.Errorf("AddEventToEventStore - DeleteEvent 1, didUpdate should be true")
	}
	if len(eventStore.Store) != 2 {
		t.Errorf("AddEventToEventStore - DeleteEvent 1, Store length should be 2")
	}

	updatePostEvent := models.Event{
		EventType: consts.POST_UPDATED_EVENT_TYPE,
		Payload:   post,
	}

	var eventStore2 models.EventStore

	eventStore2, didUpdate = models.AddEventToEventStore(eventStore2, createPostEvent)
	eventStore2, didUpdate = models.AddEventToEventStore(eventStore2, updatePostEvent)
	eventStore2, didUpdate = models.AddEventToEventStore(eventStore2, deletePostEvent)

	if didUpdate != true {
		t.Errorf("AddEventToEventStore - DeleteEvent 2, didUpdate should be true")
	}
	if len(eventStore2.Store) != 3 {
		t.Errorf("AddEventToEventStore - DeleteEvent 2, Store length should be 3")
	}

	var eventStore3 models.EventStore

	eventStore3, didUpdate = models.AddEventToEventStore(eventStore3, createPostEvent)
	eventStore3, didUpdate = models.AddEventToEventStore(eventStore3, deletePostEvent)
	eventStore3, didUpdate = models.AddEventToEventStore(eventStore3, updatePostEvent)

	if didUpdate != false {
		t.Errorf("AddEventToEventStore - DeleteEvent 3, didUpdate should be false")
	}
	if len(eventStore3.Store) != 2 {
		t.Errorf("AddEventToEventStore - DeleteEvent 3, Store length should be 2")
	}
}
