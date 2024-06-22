package main

import (
	"log"
	"go-csrf/server"
	"go-csrf/server/middleware/myJwt"
	"go-csrf/db"
)

var host = "localhost"
var port = "8000"

func main() {
	db.InitDB()

	jwtErr := myJwt.InitJWT()
	if jwtErr != nil {
		log.Println("Error initializing the JWT!")
		log.Fatal(jwtErr)
	}

	serverErr := server.StartServer(host, port)
	if serverErr != nil {
		log.Println("Error starting the server!")	
		log.Fatal(serverErr)
	}
}
