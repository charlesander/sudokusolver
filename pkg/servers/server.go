package servers

import (
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"sudokusolver/pkg/boards"
	"sudokusolver/pkg/cells"
	"sudokusolver/pkg/solvers"
	utilties "sudokusolver/pkg/utilites"
	"time"
)

type ResponseTemplateStruct struct {
	Complete bool
	Message  string
	Body     []int
}

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
	router.HandleFunc("/solve", handleSolve).Methods("GET")

	return router
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "pong")
}

func handleSolve(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, X-Request-With")

	stringCellValues, ok := r.URL.Query()["cellValues[]"]
	intCellValues := []int{}
	var err error
	if !ok {
		displayError(w, errors.New("Url Param 'cellValues[]' is missing"))
		return
	} else {
		intCellValues, err = utilties.ConvertStringSliceToIntSlice(stringCellValues)
	}

	board, err := boards.NewBoard(cells.NewFactory(), intCellValues)
	if err != nil {
		displayError(w, err)
		return
	}
	ok, err = solvers.Solve(board)
	if err != nil {
		displayError(w, err)
		return
	} else if !ok {
		displayError(w, errors.New("Unable to solve"))
		return
	}
	values := []int{}
	for i := 0; i < boards.CELL_COUNT; i++ {
		cell, err := board.GetCell(i)
		if err != nil {
			displayError(w, err)
			return
		}
		values = append(values, cell.GetCellValue())
	}

	var jsonResponse = ResponseTemplateStruct{true, "OK", values}

	json.NewEncoder(w).Encode(jsonResponse)
}

func displayError(w http.ResponseWriter, err error) {
	var jsonResponse = ResponseTemplateStruct{false, err.Error(), nil}
	json.NewEncoder(w).Encode(jsonResponse)
}
