package geocode

import (
	"arcgis-sdk/auth"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

// load environment variables using godotenv

var (
	authManager auth.AuthenticationManager
	geocoder    *Geocoder
	revGeocoder *ReverseGeocoder
)

func setup() {
	apiKey := os.Getenv("ARCGIS_KEY")
	if apiKey == "" {
		if err := godotenv.Load("../.env"); err != nil {
			panic("Error loading .env file: " + err.Error())
		}
		apiKey = os.Getenv("ARCGIS_KEY")
		if apiKey == "" {
			panic("ARCGIS_KEY environment variable not set")
		}
	}
	authManager = auth.NewApiKeyManager(apiKey)
	geocoder, _ = NewGeocoder(authManager)
	revGeocoder, _ = NewReverseGeocoder(authManager)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestReverseGeocoder_ReverseGeocode(t *testing.T) {

	// Define a test location
	testLocation := Location{
		Latitude:  34.0564,
		Longitude: -117.1956,
	}

	// Call the ReverseGeocode method
	response, err := revGeocoder.ReverseGeocode(testLocation)
	if err != nil {
		t.Errorf("ReverseGeocode returned an error: %v", err)
	}

	// Check the response for the expected values
	// Note: You will need to adjust these checks based on the expected response format and values
	if len(response.Address) == 0 {
		t.Errorf("Expected at least one address in the response")
	}

	// Add more checks as needed
}

func TestGeocoder_Geocode(t *testing.T) {

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

func TestReverseGeocoder_ReverseGeocodeInvalidLocation(t *testing.T) {

	_, err := revGeocoder.ReverseGeocode([]float64{34.0564, -117.1956, 0.0, 0.0})
	if err == nil {
		t.Errorf("Expected an error with an invalid location")
	}
}
