package models

type Hotel struct {
	ID        int    `json:"id,omitempty" gorm:"primaryKey;unique;type:int"`
	HotelName string `json:"hotel_name,omitempty" gorm:"type:varchar"`
	Address   string `json:"address,omitempty" gorm:"type:varchar"`
}

type Hotels []Hotel
