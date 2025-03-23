package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/roughtomato/notifyMe/db"
	"github.com/roughtomato/notifyMe/handlers"
)

func main() {
	gormDB, err := db.InitDB()
	if err != nil {
		log.Fatal("Failed to initialize database:", err)
	}

	r := mux.NewRouter()
	notificationHandler := handlers.NewNotificationHandler(gormDB)

	r.HandleFunc("/notifications/{userId}", notificationHandler.GetAllNotificationsByUserId).Methods("GET")
	r.HandleFunc("/notifications", notificationHandler.AddNotification).Methods("POST")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
