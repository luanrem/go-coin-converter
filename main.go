package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
)

type ExchangeRate struct {
	Base  string             `json:"base"`
	Date  string             `json:"date"`
	Rates map[string]float64 `json:"rates"`
}

func main() {
	args := os.Args[1:]
	valor, _ := strconv.ParseFloat(args[0], 64)
	moeda := args[1]

	// get value from json
	data, err := os.ReadFile("taxa-de-cambio.json")
	if err != nil {
		panic(err)
	}

	taxas := ExchangeRate{}

	jsonErr := json.Unmarshal(data, &taxas)
	if jsonErr != nil {
		panic(jsonErr)
	}

	finalValue := valor * taxas.Rates[moeda]

	// resultado final
	fmt.Printf("%.2f\n", finalValue)
}
