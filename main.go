package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/joho/godotenv"
	"github.com/raghuganapathyUCR/arcgis-sdk/auth"
	"github.com/raghuganapathyUCR/arcgis-sdk/geocode"
)

func testGeocoder(AUTH auth.AuthenticationManager) {
	geocoder, err := geocode.NewGeocoder(AUTH)
	if err != nil {
		fmt.Println(err)
		return
	}

	/*
			address: "380 New York St",
		    city: "Redlands",
		    region: "CA",
		    postal: "92373",
	*/

	address := geocode.GeocodeRequestOptions{}

	// addr := "Starbucks"
	start := time.Now()
	r, err := geocoder.Geocode(address)
	sinec := time.Since(start)
	fmt.Println("Time taken to geocode addr: ", sinec)
	if err != nil {
		fmt.Println("MainError: ", err)
		return
	}
	q, _ := json.MarshalIndent(r, "", "  ")

	fmt.Printf("\n Result: %+s", q)
}

func testReverseGeocoder(AUTH auth.AuthenticationManager) {
	point := geocode.Point{
		// X: -117.1956,
		// Y: 34.0564,
		X: -17.1956,
		Y: 4.0564,
	}

	revGeocoder, _ := geocode.NewReverseGeocoder(AUTH)
	rev, err := revGeocoder.ReverseGeocode(point)
	if err != nil {
		fmt.Println("Error reverse geocoding: ", err)
		return
	}
	revq, _ := json.MarshalIndent(rev, "", "  ")
	fmt.Printf("\n Reverse Geocode Result: %+s", revq)

}
func main() {
	err := godotenv.Load() // Load environment variables from .env file
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
		return
	}

	var AUTH = auth.NewApiKeyManager("S")

	// testGeocoder(AUTH)
	testReverseGeocoder(AUTH)
}
