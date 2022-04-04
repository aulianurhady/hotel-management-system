package routes

import (
	"hotel-management-system/app/controller"
	"net/http"
)

func Routes() {
	http.HandleFunc("/roomsearch", controller.HandleRoomSearch)
}
