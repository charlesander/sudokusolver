package main

import (
	"fmt"
	"sudokusolver/pkg/servers"
)

func main() {
	fmt.Println("Sudoku solvers")

	router := servers.GenerateHandler()

	servers.LaunchServer(router)
}


