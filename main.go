package main

import (
	"github.com/tiagompalte/go-rest-api/database"
	"github.com/tiagompalte/go-rest-api/routes"
)

func main() {
	database.ConectaDatabase()
	routes.HandleResponse()
}
