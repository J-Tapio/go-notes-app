package main

import (
	"log"

	"go-notes-app/db"
	"go-notes-app/handler"
	"go-notes-app/server"
)

func init() {
	db.InitDBConnection()
	handler.InitRouter()
}

func tidyUp() {
	log.Println("Exited the app.")
}

func main() {
	server.Start()

	defer tidyUp()
	defer db.Disconnect()
}
