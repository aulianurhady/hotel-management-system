package models

import (
	"encoding/json"
	"net/http"
)

type ResponseData struct {
	RoomQuantity   int            `json:"room_qty"`
	RoomTypeID     int            `json:"room_type_id"`
	CheckinDate    string         `json:"checkin_date"`
	CheckoutDate   string         `json:"checkout_date"`
	TotalPrice     int            `json:"total_price"`
	AvailableRooms AvailableRooms `json:"available_rooms"`
}

type JsonResponse struct {
	StatusCode int         `json:"status_code"`
	Data       interface{} `json:"data,omitempty"`
	Message    string      `json:"message,omitempty"`
}

func Response(w http.ResponseWriter, statusCode int, data interface{}, message string) {
	w.Header().Set("Content-Type", "application/json")

	response := JsonResponse{
		StatusCode: statusCode,
		Data:       data,
		Message:    message,
	}

	json.NewEncoder(w).Encode(response)
}
