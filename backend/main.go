package main

import (
	"github/ANNMAINAWANGARI/FintechApp/api"
)

func main() {
	server := api.NewServer(".")
	server.Start(8000)
}