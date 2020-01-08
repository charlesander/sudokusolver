package servers

import (
	"encoding/json"
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
	if !ok  {
		log.Println("Url Param 'cellValues[]' is missing")
		return
	} else {
		intCellValues, err = utilties.ConvertStringSliceToIntSlice(stringCellValues)

	}

	board, err := boards.NewBoard(cells.NewFactory(), intCellValues)
	if(err != nil) {
		panic("TODO RETURN ERROR")
	}
	solvers.Solve(board)

	values := []int{}
	for i := 0; i < boards.CELL_COUNT; i++ {
		values  = append(values, board.GetCell(i).GetCellValue())
	}
	json.NewEncoder(w).Encode(values)

}