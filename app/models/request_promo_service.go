package models

type RequestPromoService struct {
	RoomQuantity   int            `json:"room_qty"`
	RoomTypeID     int            `json:"room_type_id"`
	CheckinDate    string         `json:"checkin_date"`
	CheckoutDate   string         `json:"checkout_date"`
	TotalPrice     int            `json:"total_price"`
	AvailableRooms AvailableRooms `json:"available_rooms"`
	PromoID        int            `json:"promo_id"`
}
