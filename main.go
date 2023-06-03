package main

import (
	"assignment-3/database"
	"assignment-3/routers"
)

func main() {
	database.StartDB()
	var PORT = ":8100"

	routers.StartServer().Run(PORT)
}
