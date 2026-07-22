package main

import (
	"gotickets/internal/config"
	"gotickets/internal/server"
)

func main() {
	//load environment variable
	cfg := config.LoadEnv()
	
	//connect to database
	db := config.ConnectDatabase(cfg)

	//start the server
	server.Start(db, cfg)

}
