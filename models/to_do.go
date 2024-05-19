package models

import "time"

type ToDo struct {
	ID              int       `json:"id"`
	Description     string    `json:"description"`
	CreatedOn       time.Time `json:"createdon"`
	ChangedOn       time.Time `json:"changedon"`
	DeletedOn       time.Time `json:"deletedon"`
	PercentComplete float64   `json:"percentcomplete"`
	User            string    `json:"user"`
}
