package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Estructura para mapear la respuesta JSON
type GoldResponse struct {
	Price float64 `json:"price"`
}

func getGoldPrice() (float64, error) {
	apiURL := "https://api.gold-api.com/price/XAU"
	resp, err := http.Get(apiURL)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	// Decodificar la respuesta JSON en la estructura GoldResponse
	var goldResp GoldResponse
	if err := json.NewDecoder(resp.Body).Decode(&goldResp); err != nil {
		return 0, err
	}

	return goldResp.Price, nil
}

func getMandarinPrice() float64 {
	// El precio de una mandarina es aproximadamente $0.33 USD
	return 0.33
}

func calculateMandarins(goldPrice, mandarinPrice float64) float64 {
	return goldPrice / mandarinPrice
}

func goldHandler(w http.ResponseWriter, r *http.Request) {
	goldPrice, err := getGoldPrice()
	if err != nil {
		http.Error(w, "Error fetching gold price", http.StatusInternalServerError)
		return
	}

	mandarinPrice := getMandarinPrice()
	mandarins := calculateMandarins(goldPrice, mandarinPrice)

	// Solo enviar el número de mandarinas como respuesta
	w.Header().Set("Content-Type", "text/plain")
	fmt.Fprintf(w, "%.2f", mandarins)
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/mandarins", goldHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
