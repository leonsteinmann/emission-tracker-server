package model

import (
	"time"
)

type Record struct {
	ID            int       `json:"id"`
	Name          string    `json:"name"`
	Datetime      time.Time `json:"datetime"`
	UserID        *int      `json:"user_id"`
	CategoryID    int       `json:"category_id"`
	SubcategoryID int       `json:"subcategory_id"`
	Amount        float64   `json:"amount"`
	UnitType      string    `json:"unit_type"`
	InputDatetime time.Time `json:"input_datetime"`
}
