package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Product struct {
	Name     string  `json:"name"`
	Price    float64 `json:"price"`
	Quantity int     `json:"quantity"`
}

func toJSON(p Product) (string, error) {
	jsonData, err := json.Marshal(p)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func fromJSON(jsonStr string) (Product, error) {
	var p Product
	err := json.Unmarshal([]byte(jsonStr), &p)
	if err != nil {
		return Product{}, err
	}
	return p, nil
}

func main() {
	product := Product{
		Name:     "Iphone 16",
		Price:    999.99,
		Quantity: 7,
	}

	jsonStr, err := toJSON(product)
	if err != nil {
		log.Fatalf("Error encoding to JSON: %v", err)
	}
	fmt.Println("JSON Output:", jsonStr)

	decodedProduct, err := fromJSON(jsonStr)
	if err != nil {
		log.Fatalf("Error decoding from JSON: %v", err)
	}
	fmt.Printf("Decoded Product: %+v\n", decodedProduct)
}
