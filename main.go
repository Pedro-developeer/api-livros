package main

import (
	"github.com/pedrodevelopeer/webapi-with-go/database"
	"github.com/pedrodevelopeer/webapi-with-go/server"
)

func main() {
	database.StartDB()
	s := server.NewServer()

	s.Run()
}
