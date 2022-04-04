package models

type Room struct {
	ID         int    `json:"id,omitempty" gorm:"type:int"`
	HotelID    int    `json:"hotel_id,omitempty" gorm:"type:int"`
	RoomTypeID int    `json:"room_type_id,omitempty" gorm:"type:int"`
	RoomNumber string `json:"room_number,omitempty" gorm:"type:varchar"`
	RoomStatus string `json:"room_status,omitempty" gorm:"type:ENUM('available', 'out_of_service')"`
}

type Rooms []Room
