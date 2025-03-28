// "https://app.exchangerate-api.com/dashboard/confirmed" -->link for my API
package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/charmbracelet/huh"
)

func main() {
	// Fetch exchange rates
	rates, err := GetExchangeRates()
	if err != nil {
		log.Fatalf("Failed to fetch exchange rates: %v", err)
	}

	var amountStr, from, to string

	// TUI Form
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Enter Amount").
				Prompt("> ").
				Validate(func(s string) error {
					if _, err := strconv.ParseFloat(s, 64); err != nil {
						return fmt.Errorf("Invalid number")
					}
					return nil
				}).
				Value(&amountStr),
		),

		huh.NewGroup(
			huh.NewSelect[string]().
				Title("From Currency").
				Options(
					huh.NewOption("USD", "USD"),
					huh.NewOption("EUR", "EUR"),
					huh.NewOption("GBP", "GBP"),
					huh.NewOption("JPY", "JPY"),
				).
				Value(&from),
		),

		huh.NewGroup(
			huh.NewSelect[string]().
				Title("To Currency").
				Options(
					huh.NewOption("USD", "USD"),
					huh.NewOption("EUR", "EUR"),
					huh.NewOption("GBP", "GBP"),
					huh.NewOption("JPY", "JPY"),
				).
				Value(&to),
		),
	)

	// Run form
	if err := form.Run(); err != nil {
		log.Fatalf("Form error: %v", err)
	}

	// Convert amount
	amount, _ := strconv.ParseFloat(amountStr, 64)
	result, err := ConvertCurrency(amount, from, to, rates)
	if err != nil {
		log.Fatalf("Conversion error: %v", err)
	}

	fmt.Printf("\nConverted Amount: %.2f %s\n", result, to)
}
