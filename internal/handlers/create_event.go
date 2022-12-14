package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
)

const (
	addQuery    = `insert into events (address, price) values ($1, $2)`
	getQuery    = `select * from events where address=$1`
	getAllQuery = `select address, price from events`
	updateQuery = `update events set price = $1`
)

type EventDB struct {
	Id      int64  `json:"id" db:"id"`
	Address string `json:"address" db:"address"`
	Price   int    `json:"price" db:"price"`
}

type EventResult struct {
	Address string `json:"address" db:"address"`
	Price   int    `json:"price" db:"price"`
}

type EventDBAddresses struct {
	Address string `json:"address" db:"address"`
}

type EventDBPrices struct {
	Price int `json:"price" db:"price"`
}

type Storage struct {
	db *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{db: db}
}

type GetEventsResponse struct {
	Event EventDB
}

type GetAllEvents struct {
	Events []*EventResult
}

func (s *Storage) Create(address string, price int) (*EventDB, error) {
	var event EventDB
	rows, err := s.db.Query(addQuery, address, price)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&event.Id); err != nil {
			return nil, err
		}
	}
	return &EventDB{Id: event.Id, Address: event.Address, Price: event.Price}, nil
}

func CreateEvent(storage *Storage, address string, price int) (*EventDB, error) {
	resp, err := storage.Create(address, price)
	if err != nil {
		log.Println("can't create event", err)
		return nil, err
	}
	return resp, nil
}

func (s *Storage) Get(address string) (*GetEventsResponse, error) {
	rows, err := s.db.Query(getQuery, address)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var event EventDB
	for rows.Next() {
		if err := rows.Scan(&event.Id, &event.Address, &event.Price); err != nil {
			return nil, err
		}
	}
	return &GetEventsResponse{Event: event}, nil
}

func (s *Storage) UpdatePrice(price int) (*EventDBPrices, error) {
	var event EventDBPrices
	rows, err := s.db.Query(updateQuery, price)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		if err := rows.Scan(&event.Price); err != nil {
			return nil, err
		}
	}
	return &EventDBPrices{Price: event.Price}, nil
}

func UpdateEvents(storage *Storage) (*EventDBPrices, error) {

	var event *EventDBPrices

	eventRes, err := storage.UpdatePrice(event.Price)
	if err != nil {
		log.Println("update events error", err)
		return nil, err
	}

	event = eventRes

	_, err = json.Marshal(event)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return event, nil
}

func (s *Storage) GetAll() (*GetAllEvents, error) {
	rows, err := s.db.Query(getAllQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var events []*EventResult
	for rows.Next() {
		var event EventResult
		if err := rows.Scan(&event.Address, &event.Price); err != nil {
			return nil, err
		}
		events = append(events, &event)
	}
	return &GetAllEvents{Events: events}, nil
}

func GetEvents(storage *Storage) ([]*EventResult, error) {

	var events []*EventResult

	eventsRes, err := storage.GetAll()
	if err != nil {
		log.Println("can't get all events", err)
		return nil, err
	}

	events = eventsRes.Events

	//_, err = json.Marshal(events)
	//if err != nil {
	//	log.Println("can't marshalling events", err)
	//}
	return events, nil
}

func GetEventByAddress(storage *Storage, address string) (*EventDB, error) {
	var event EventDB

	eventsRes, err := storage.Get(address)
	if err != nil {
		return nil, err
	}

	event = eventsRes.Event

	//_, err = json.Marshal(event)
	//if err != nil {
	//	log.Println("cant marshalling events", err)
	//	return nil, err
	//}

	return &event, nil
}

func ComparePrices(oldPrice, newPrice int) bool {
	if oldPrice != newPrice {
		return true
		fmt.Println("price is change")
	}
	fmt.Println("price not change")
	return false
}

func SendingPriceToEmail(email string) {
	fmt.Println("sending to email about changing price", email)
}
