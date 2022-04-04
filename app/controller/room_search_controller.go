package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"hotel-management-system/app/client"
	"hotel-management-system/app/models"
	"hotel-management-system/app/repository"
	"log"
	"net/http"
	"strconv"
	"time"
)

func HandleRoomSearch(w http.ResponseWriter, r *http.Request) {
	res := models.ResponseData{}
	if err := r.ParseForm(); err != nil {
		models.Response(w, http.StatusInternalServerError, res, err.Error())
		return
	}
	checkinDate := r.FormValue("checkin_date")
	checkoutDate := r.FormValue("checkout_date")
	roomQuantity := r.FormValue("room_qty")
	roomTypeID := r.FormValue("room_type_id")
	promoID := r.FormValue("promo_id")

	convRoomTypeID, _ := strconv.Atoi(roomTypeID)
	convRoomQuantity, _ := strconv.Atoi(roomQuantity)
	convPromoID, _ := strconv.Atoi(promoID)

	format := "2006-01-02"
	checkinFormated, _ := time.Parse(format, checkinDate)
	checkoutFormated, _ := time.Parse(format, checkoutDate)

	res.RoomQuantity = convRoomQuantity
	res.RoomTypeID = convRoomTypeID
	res.CheckinDate = checkinFormated.Format(format)
	res.CheckoutDate = checkoutFormated.Format(format)

	dataPrices, err := repository.GetPrices(checkinDate, checkoutDate, convRoomTypeID)
	if err != nil {
		models.Response(w, http.StatusInternalServerError, res, err.Error())
		return
	}

	dataRooms, err := repository.GetRooms(convRoomTypeID)
	if err != nil {
		models.Response(w, http.StatusInternalServerError, res, err.Error())
		return
	}

	diff := checkoutFormated.Sub(checkinFormated)
	totalDiffDays := diff.Hours() / 24

	if len(dataPrices) != int(totalDiffDays) {
		log.Println("No available rooms in the time range")
		models.Response(w, http.StatusNoContent, nil, errors.New("No room available in the time range").Error())
		return
	}

	dataAvailableRooms := models.AvailableRooms{}
	if len(dataPrices) >= 1 {
		var newPricesObj models.Prices
		for _, val1 := range dataPrices {
			newPricesObj = append(newPricesObj, models.Price{
				Date:  val1.Date,
				Price: val1.Price,
			})
		}

		for _, val2 := range dataRooms {
			dataAvailableRooms = append(dataAvailableRooms, models.AvailableRoom{
				RoomID:     val2.ID,
				RoomNumber: val2.RoomNumber,
				Prices:     newPricesObj,
			})
		}
	}

	totalPrice := repository.GetTotalPrice() * convRoomQuantity

	if convPromoID != 0 {
		params := models.RequestPromoService{
			RoomQuantity:   convRoomQuantity,
			RoomTypeID:     convRoomTypeID,
			CheckinDate:    checkinDate,
			CheckoutDate:   checkoutDate,
			TotalPrice:     totalPrice,
			AvailableRooms: dataAvailableRooms,
			PromoID:        convPromoID,
		}

		jsonBody, _ := json.Marshal(params)
		response, _ := client.SendRequest("http://127.0.0.1:8081/redeem-promo", "POST", string(jsonBody))
		str := fmt.Sprintf("%#v", response)
		fmt.Println("RESPONSE: ", str)

		newdataAvailableRooms := models.AvailableRooms{}
		if len(dataPrices) >= 1 {
			for idx, val := range dataRooms {
				newdataAvailableRooms = append(newdataAvailableRooms, models.AvailableRoom{
					RoomID:     val.ID,
					RoomNumber: val.RoomNumber,
					Prices:     response.AvailableRooms[idx].Prices,
				})
			}
		}

		res.TotalPrice = response.TotalPrice
		res.AvailableRooms = dataAvailableRooms
	} else {
		res.TotalPrice = totalPrice
		res.AvailableRooms = dataAvailableRooms
	}

	models.Response(w, http.StatusOK, res, "")
}
