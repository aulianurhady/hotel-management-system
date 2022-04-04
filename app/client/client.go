package client

import (
	"encoding/json"
	"fmt"
	"hotel-management-system/app/models"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

func SendRequest(url string, method string, params string) (models.ResponseData, error) {
	var response models.ResponseData
	Request, err := http.NewRequest(method, url, strings.NewReader(params))
	if err != nil {
		return response, err
	}

	Request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(Request)
	if err != nil {
		return response, err
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)

	str := fmt.Sprintf("%#v", resp.Body)
	fmt.Println("RESP BODY: ", str)

	json.Unmarshal(body, &response)

	return response, nil
}

func SetParams(List models.AvailableRooms, roomQty, totalPrice, promoID int, checkinDate, checkoutDate string) *url.Values {
	params := url.Values{}
	str := fmt.Sprintf("%#v", List)
	params.Set("available_rooms", str)
	params.Set("total_price", strconv.Itoa(totalPrice))
	params.Set("room_qty", strconv.Itoa(roomQty))
	params.Set("checkin_date", checkinDate)
	params.Set("checkout_date", checkoutDate)
	params.Set("promo_id", strconv.Itoa(promoID))

	return &params
}
