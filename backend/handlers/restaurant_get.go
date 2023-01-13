package handlers

import (
	"backend/common"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func RestaurantGetHandler(w http.ResponseWriter, r *http.Request) {
	restaurant := common.Restaurant{}
	byte_data, err := json.Marshal(restaurant)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(byte_data); err != nil {
		fmt.Println(err)
	}
}
