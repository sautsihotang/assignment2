package main

import (
	"assignment2/config"
	"assignment2/routes"
)

func main() {
	config.StartDB()
	routes.StartingServer().Run(":8080")
}
