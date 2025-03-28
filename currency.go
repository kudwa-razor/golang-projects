package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const apiKey = "a2c11ed7da00b96b53efd3ac" // Your API Key
const apiURL = "https://v6.exchangerate-api.com/v6/%s/latest/USD"

// ExchangeRates struct to store API response
type ExchangeRates struct {
	Rates map[string]float64 `json:"conversion_rates"`
}

// Fetch exchange rates from API
func GetExchangeRates() (map[string]float64, error) {
	url := fmt.Sprintf(apiURL, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch exchange rates: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %s", resp.Status)
	}

	var data ExchangeRates
	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, fmt.Errorf("failed to parse JSON response: %w", err)
	}

	return data.Rates, nil
}

// Convert currency using exchange rates
func ConvertCurrency(amount float64, from string, to string, rates map[string]float64) (float64, error) {
	rateFrom, existsFrom := rates[from]
	rateTo, existsTo := rates[to]

	if !existsFrom || !existsTo {
		return 0, fmt.Errorf("unsupported currency: %s or %s", from, to)
	}

	converted := amount * (rateTo / rateFrom)
	return converted, nil
}
