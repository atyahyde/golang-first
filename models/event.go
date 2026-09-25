package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Name        string    `json:"name" binding:"required"`
	Description string    `json:"description" binding:"required"`
	Location    string    `json:"location" binding:"required"`
	UserId      int       `json:"user_id"`
	Datetime    time.Time `json:"datetime" binding:"required"`
}

var events []Event = []Event{}

// fungsi untuk simpan data event
func (e Event) Save() {
	events = append(events, e)
}

// fungsi menampilkan tampil semua data event
func GetAllEvents() []Event {
	return events
}
