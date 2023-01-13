package handlers

import (
	"backend/common"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type restaurantList struct {
	Restaurants []common.RestaurantMetadata `json:"restaurants"`
}

func RestaurantListHandler(w http.ResponseWriter, r *http.Request) {

	resList := restaurantList{}

	byteResponse, err := json.Marshal(resList)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(byteResponse); err != nil {
		log.Println(err)
	}

}
