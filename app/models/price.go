package models

import (
	"time"
)

type Price struct {
	ID         int        `json:"id,omitempty" gorm:"primaryKey;unique;type:int"`
	Date       *time.Time `json:"date,omitempty" gorm:"type:timestamptz"`
	RoomTypeID int        `json:"room_type_id,omitempty" gorm:"type:int"`
	Price      int        `json:"price,omitempty" gorm:"type:int"`
}

type Prices []Price
