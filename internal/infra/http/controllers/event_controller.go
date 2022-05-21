package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/test_server/internal/service/database"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/test_server/internal/domain/event"
)

type EventController struct {
	service *event.Service
}

func NewEventController(s *event.Service) *EventController {
	return &EventController{
		service: s,
	}
}

func (c *EventController) FindAll() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		events := (*c.service).FindAll()

		err := success(w, events)
		if err != nil {
			fmt.Printf("EventController.FindAll(): %s", err)
		}
	}
}

func (c *EventController) FindOne() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.FindOne(): %s", err)
			}
			return
		}

		event, err := (*c.service).FindOne(id)
		if err != nil {
			err = notFound(w, err)
			if err != nil {
				fmt.Println(err)
			}
			return
		}

		err = success(w, event)
		if err != nil {
			fmt.Printf("EventController.FindOne(): %s", err)
		}
	}
}

func (c *EventController) Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sess := database.OpenConnection()
		defer database.CloseConnection(sess)

		eventsCol := sess.Collection("event")

		var event event.Event
		err := json.NewDecoder(r.Body).Decode(&event)
		if err != nil {
			fmt.Println(err)
			err = badRequest(w, err)
			if err != nil {
				fmt.Println(err)
			}
			return
		}
		_, err = eventsCol.Insert(&event)
		if err != nil {
			fmt.Println(err)
			err := badRequest(w, err)
			if err != nil {
				fmt.Println(err)
			}
			return
		}

		err = noContent(w)
		if err != nil {
			fmt.Printf("EventController.Create(): %s", err)
		}
	}
}

func (c *EventController) Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.Update(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.Update: %s", err)
			}
			return
		}

		sess := database.OpenConnection()
		defer database.CloseConnection(sess)

		eventsCol := sess.Collection("event")

		var event event.Event

		eventRes := eventsCol.Find(id)
		err = eventRes.One(&event)
		if err != nil {
			err = notFound(w, err)
			if err != nil {
				fmt.Println(err)
			}
			return
		}

		err = json.NewDecoder(r.Body).Decode(&event)
		if err != nil {
			fmt.Println(err)
			err = badRequest(w, err)
			if err != nil {
				fmt.Println(err)
			}
			return
		}

		err = eventRes.Update(&event)
		if err != nil {
			fmt.Println(err)
			err := badRequest(w, err)
			if err != nil {
				fmt.Println(err)
			}
			return
		}

		err = noContent(w)
		if err != nil {
			fmt.Printf("EventController.Update(): %s", err)
		}
	}
}
func (c *EventController) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		id, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
		if err != nil {
			fmt.Printf("EventController.Delete(): %s", err)
			err = internalServerError(w, err)
			if err != nil {
				fmt.Printf("EventController.Delete: %s", err)
			}
			return
		}

		sess := database.OpenConnection()
		defer database.CloseConnection(sess)

		eventsCol := sess.Collection("event")

		var event event.Event

		eventRes := eventsCol.Find(id)
		err = eventRes.One(&event)
		if err != nil {
			err = notFound(w, err)
			if err != nil {
				fmt.Println(err)
			}
			return
		}

		err = eventRes.Delete()
		if err != nil {
			fmt.Println(err)
			err := badRequest(w, err)
			if err != nil {
				fmt.Println(err)
			}
			return
		}

		err = noContent(w)
		if err != nil {
			fmt.Printf("EventController.Delete(): %s", err)
		}
	}
}
