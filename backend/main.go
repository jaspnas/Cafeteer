package main

import (
	"backend/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/restaurant-list", handlers.RestaurantListHandler)
	http.HandleFunc("/restaurant", handlers.RestaurantGetHandler)
	if err := http.ListenAndServe(os.Getenv("SERVER_ADDR"), nil); err != nil {
		log.Fatal(err)
	}

}
