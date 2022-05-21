package event

import (
	"fmt"
	"github.com/test_server/internal/service/database"
)

type Repository interface {
	FindAll() []Event
	FindOne(id int64) (*Event, error)
}

//const EventsCount int64 = 10

type repository struct {
	// Some internal data
}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) FindAll() []Event {
	sess := database.OpenConnection()
	defer database.CloseConnection(sess)

	eventsCol := sess.Collection("event")
	var events []Event
	err := eventsCol.Find().All(&events)
	if err != nil {
		fmt.Println("eventsCol.Find: ", err)
	}

	return events
}

func (r *repository) FindOne(id int64) (*Event, error) {
	sess := database.OpenConnection()
	defer database.CloseConnection(sess)

	eventsCol := sess.Collection("event")
	var event Event
	err := eventsCol.Find(id).One(&event)

	return &event, err
}
