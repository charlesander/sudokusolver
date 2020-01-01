package main

import (
	"sudokusolver/pkg/servers"
)

func main() {

	router := servers.GenerateHandler()

	servers.LaunchServer(router)
}


