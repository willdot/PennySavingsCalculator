package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/willdot/PennySavingsCalculator/calculator"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type request struct {
	Start time.Time
	End   time.Time
}

func main() {

	var PORT string
	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "8080"
	}
	
	router := mux.NewRouter()

	router.HandleFunc("/calculate", GetBudget)

	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:4200"},
	  })

	handler := c.Handler(router)

	srv := &http.Server{
		Handler: handler,
		Addr:    ":" + PORT,
	  }
	
	  log.Fatal(srv.ListenAndServe())
}

// GetBudget takes a start and end date and returns how much to save in that period
func GetBudget(w http.ResponseWriter, r *http.Request) {

	var req request

	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&req)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	amount, err := calculator.CalculateHowMuchToSaveBetweenDays(req.Start, req.End)

	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	result := fmt.Sprintf("£%v", float64(amount)/100)

	json.NewEncoder(w).Encode(result)
}
