package repository

import (
	"hotel-management-system/app/database"
	"hotel-management-system/app/models"
	"log"
)

var (
	TotalPrice int = 0
)

func GetPrices(checkinDate, checkoutDate string, roomTypeID int) (models.Prices, error) {
	var prices models.Prices

	sql := `
		SELECT id, date, room_type_id, price FROM price WHERE (date >= $1 AND date < $2) AND room_type_id = $3
	`
	data, err := database.DB.Query(sql, checkinDate, checkoutDate, roomTypeID)
	if err != nil {
		log.Printf("Error query %s", err)
		return prices, err
	}

	var totalPrice int = 0
	for data.Next() {
		var perPrice models.Price
		err = data.Scan(&perPrice.ID, &perPrice.Date, &perPrice.RoomTypeID, &perPrice.Price)
		if err != nil {
			log.Printf("Error scan %s", err)
			return prices, err
		}

		totalPrice += perPrice.Price

		prices = append(prices, perPrice)
	}

	TotalPrice = totalPrice

	return prices, nil
}

func GetTotalPrice() int {
	return TotalPrice
}
