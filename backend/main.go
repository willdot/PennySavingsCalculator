package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"github.com/willdot/PennySavingsCalculator/backend/calculator"
)

type request struct {
	Start time.Time
	End   time.Time
}

func main() {

	var port string
	if port = os.Getenv("PORT"); port == "" {
		port = "8080"
	}

	var apiKey string
	if apiKey = os.Getenv("APIKEY"); apiKey == "" {
		log.Fatal(errors.New("No api key set"))
		return
	}

	router := mux.NewRouter()

	router.HandleFunc("/calculate", CheckAPIKey(apiKey, GetBudget()))

	c := cors.New(cors.Options{
		AllowedOrigins: []string{
			"https://1pdailysavingscalculatorapp.azurewebsites.net",
		},
	})

	handler := c.Handler(router)

	err := http.ListenAndServe(":"+port, handler)

	if err != nil {
		log.Fatal(err)
	}
}

// GetBudget takes a start and end date and returns how much to save in that period
func GetBudget() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

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

		result := fmt.Sprintf("£%.2f", float64(amount)/100)

		json.NewEncoder(w).Encode(result)
	}
}

// CheckAPIKey makes sure the origin is correct. I only want my api to be accessible from the front end I create
func CheckAPIKey(apiKey string, h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("ApiKey") != apiKey {
			http.Error(w, errors.New("Api key not valid").Error(), http.StatusForbidden)
			return
		}
		h.ServeHTTP(w, r)
	})
}
