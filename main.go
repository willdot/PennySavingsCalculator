package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	calculator "PennyBudgetCalculator/calculator"

	"github.com/gorilla/mux"
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

	err := http.ListenAndServe(":"+PORT, router)

	if err != nil {
		log.Fatal(err)
	}
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
