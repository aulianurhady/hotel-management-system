package models

type AvailableRoom struct {
	RoomID     int    `json:"room_id,omitempty"`
	RoomNumber string `json:"room_number,omitempty"`
	Prices     Prices `json:"price,omitempty"`
}

type AvailableRooms []AvailableRoom
