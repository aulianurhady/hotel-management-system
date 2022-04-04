package repository

import (
	"hotel-management-system/app/database"
	"hotel-management-system/app/models"
	"log"
)

func GetRooms(roomTypeID int) (models.Rooms, error) {
	var rooms models.Rooms

	availableRooms := "available"
	sql := `
		SELECT id, hotel_id, room_type_id, room_number, room_status FROM room
		WHERE room_type_id = $1 AND room_status = $2
	`
	data, err := database.DB.Query(sql, roomTypeID, availableRooms)
	if err != nil {
		log.Printf("Error query %s", err)
		return rooms, err
	}

	for data.Next() {
		var perRoom models.Room
		err = data.Scan(&perRoom.ID, &perRoom.HotelID, &perRoom.RoomTypeID, &perRoom.RoomNumber, &perRoom.RoomStatus)
		if err != nil {
			log.Printf("Error scan %s", err)
			return rooms, err
		}

		rooms = append(rooms, perRoom)
	}

	return rooms, nil
}
