package geocode

import (
	"arcgis-sdk/auth"
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

// load environment variables using godotenv

func TestGeocoder_Geocode(t *testing.T) {
	// load environment variables using godotenv
	if err := godotenv.Load("../../.env"); err != nil {

		fmt.Printf("Error loading .env file: %v\n", err)
		return
	}

	authManager := auth.NewApiKeyManager(os.Getenv("ARCGIS_API_KEY"))

	// Create a new Geocoder instance
	geocoder, err := NewGeocoder(authManager)
	if err != nil {
		t.Fatalf("Failed to create Geocoder: %v", err)
	}

	// Define a test address
	testAddress := GeocodeRequestOptions{
		SingleLine: "1600 Pennsylvania Ave NW, Washington, DC",
	}

	// Call the Geocode method
	response, err := geocoder.Geocode(testAddress)
	if err != nil {
		t.Errorf("Geocode returned an error: %v", err)
	}

	// Check the response for the expected values
	// Note: You will need to adjust these checks based on the expected response format and values
	if len(response.Candidates) == 0 {
		t.Errorf("Expected at least one candidate in the response")
	}

	// Add more checks as needed
}
func TestGeocoder_GeocodeInvalidApiKey(t *testing.T) {
	authManager := auth.NewApiKeyManager("invalid-api-key")
	geocoder, _ := NewGeocoder(authManager)

	_, err := geocoder.Geocode(GeocodeRequestOptions{SingleLine: "1600 Pennsylvania Ave NW, Washington, DC"})
	if err == nil {
		t.Errorf("Expected an error with an invalid API key")
	}
}

func TestGeocoder_GeocodeEmptyAddress(t *testing.T) {
	if err := godotenv.Load("../../.env"); err != nil {

		fmt.Printf("Error loading .env file: %v\n", err)
		return
	}

	authManager := auth.NewApiKeyManager(os.Getenv("ARCGIS_API_KEY"))

	geocoder, _ := NewGeocoder(authManager)

	r, err := geocoder.Geocode(GeocodeRequestOptions{})
	// r is Response: {SpatialReference:{WKID:0 LatestWKID:0} Candidates:[]} because the address is empty check the response for the expected values
	if err != nil {
		t.Errorf("Geocode returned an error: %v", err)
	}

	fmt.Printf("Response: %+v\n", r)
	if len(r.Candidates) != 0 {
		t.Errorf("Expected no candidates in the response")
	}
}

func TestGeocoder_GeocodeInvalidAddress(t *testing.T) {
	// load environment variables using godotenv
	if err := godotenv.Load("../../.env"); err != nil {

		fmt.Printf("Error loading .env file: %v\n", err)
		return
	}

	authManager := auth.NewApiKeyManager(os.Getenv("ARCGIS_API_KEY"))

	geocoder, _ := NewGeocoder(authManager)

	_, err := geocoder.Geocode(GeocodeRequestOptions{SingleLine: "Invalid Address"})
	if err == nil {
		t.Errorf("Expected an error with an invalid address")
	}
}