package main

import (
	"arcgis-sdk/auth"    // Path: ../../api/auth/auth.go
	"arcgis-sdk/geocode" // Path: ../../api/geocode/geocode.go
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
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

	address := geocode.GeocodeRequestOptions{
		// Address:   "380 New York St",
		// City:      "Redlands",
		// Region:    "CA",
		// Postal:    "92373",
		// OutFields: []string{"*"},
		SingleLine: "1600 Pennsylvania Ave NW, Washington, DC",
	}

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
	point := geocode.Location{
		Latitude:  34.0564,
		Longitude: -117.1956,
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

	var AUTH = auth.NewApiKeyManager(os.Getenv("ARCGIS_KEY"))

	testGeocoder(AUTH)
}
