package main

import (
	"arcgis-sdk/auth"    // Path: ../../api/auth/auth.go
	"arcgis-sdk/geocode" // Path: ../../api/geocode/geocode.go
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load() // Load environment variables from .env file
	if err != nil {
		fmt.Printf("Error loading .env file: %v\n", err)
		return
	}

	var AUTH = auth.NewApiKeyManager(os.Getenv("ARCGIS_KEY"))
	geocoder, err := geocode.NewGeocoder(AUTH)
	if err != nil {
		fmt.Println(err)
		return
	}

	// address := geocode.MultiLineAddress{
	// 	Address:     "306 Phase 3",
	// 	Address2:    "Golden Park Apartments",
	// 	Address3:    "Devarachikkanahalli Main Road",
	// 	City:        "Bangalore",
	// 	Region:      "Karnataka",
	// 	Postal:      "560076",
	// 	CountryCode: "IND",
	// 	OutFields:   []string{"PlaceName"},
	// }

	// t, _ := json.Marshal(address)
	addr := "Starbucks"
	start := time.Now()
	r, err := geocoder.Geocode(addr, false)
	sinec := time.Since(start)
	fmt.Println("Time taken to geocode addr: ", sinec)
	if err != nil {
		fmt.Println("MainError: ", err)
		return
	}
	fmt.Printf("\nAddress: %+v", r)
}
