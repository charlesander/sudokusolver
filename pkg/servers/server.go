package servers

import (
	"bytes"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"strconv"
	"sudokusolver/pkg/boards"
	"sudokusolver/pkg/cells"
	"sudokusolver/pkg/solvers"
	utilties "sudokusolver/pkg/utilites"
	"time"
)

func LaunchServer(router *mux.Router) {
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}

func GenerateHandler() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/ping", handlePing).Methods("GET")
	router.HandleFunc("/solve/{sudokuCellValues}", handleSolve).Methods("GET")

	return router
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "pong")
}

func handleSolve(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	sudokuCellValues := params["sudokuCellValues"]
	cellValues := utilties.ExplodeToIntSlice(sudokuCellValues)

	board := boards.NewBoard(cells.NewFactory(), cellValues)
	solvers.Solve(board)

	fmt.Println(cellValues)


	buf := bytes.Buffer{}
	for i := 0; i < boards.CELL_COUNT; i++ {
		buf.WriteString(strconv.Itoa(board.GetCell(i).GetCellValue()))
		//fmt.Print( i , ",")
		if (i+1)%9 == 0 {
			buf.WriteString("\n")
		}
	}
	result := buf.String()
	//io.WriteString(w, strconv.FormatBool(board.CheckComplete()))
	io.WriteString(w, result)
}