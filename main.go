package main

import (
	"github.com/Rayelisson/api-go-gin/database"
	"github.com/Rayelisson/api-go-gin/routes"
)

func main() {
	database.ConectaComBancoDeDados()
	routes.HandleRequests()
}
