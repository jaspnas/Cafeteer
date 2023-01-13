package main

import (
	"backend/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/restaurants", handlers.RestaurantListHandler)
	if err := http.ListenAndServe(os.Getenv("SERVER_ADDR"), nil); err != nil {
		log.Fatal(err)
	}

}
