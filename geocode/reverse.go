package geocode

import (
	"github.com/raghuganapathyUCR/arcgis-sdk/auth"
	"github.com/raghuganapathyUCR/arcgis-sdk/requests"
	"github.com/raghuganapathyUCR/arcgis-sdk/utils"
	"encoding/json"
	"errors"
	"fmt"
)

type ReverseGeocodeResponse struct {
	Address  map[string]interface{} `json:"address"`
	Location Point                  `json:"location"`
}

type SpatialReference struct {
	WKID          *int    `json:"wkid,omitempty"`
	LatestWKID    *int    `json:"latestWkid,omitempty"`
	VCSWKID       *int    `json:"vcsWkid,omitempty"`
	LatestVCSWKID *int    `json:"latestVcsWkid,omitempty"`
	WKT           *string `json:"wkt,omitempty"`
	LatestWKT     *string `json:"latestWkt,omitempty"`
}

type HasZM struct {
	HasZ bool `json:"hasZ"`
	HasM bool `json:"hasM"`
}

type Point struct {
	X                float64          `json:"x"`
	Y                float64          `json:"y"`
	Z                *float64         `json:"z,omitempty"`
	HasZM            HasZM            `json:"hasZM"`
	SpatialReference SpatialReference `json:"spatialReference"`
}

// GeocodeResponse is a struct that represents the response from the GeocodeService.
type Location struct {
	Latitude  float64 `json:"x,omitempty"`
	Longitude float64 `json:"y,omitempty"`
	Lat       float64 `json:"lat,omitempty"`
	Long      float64 `json:"long,omitempty"`
	Z         float64 `json:"z,omitempty"`
}

type ReverseGeocoder struct {
	Authentication auth.AuthenticationManager
}

func NewReverseGeocoder(auth auth.AuthenticationManager) (*ReverseGeocoder, error) {
	if auth == nil {
		return nil, errors.New("authentication manager needed to create a Geocoder")
	}
	return &ReverseGeocoder{
		Authentication: auth,
	}, nil
}

func (g *ReverseGeocoder) ReverseGeocode(coords interface{}) (ReverseGeocodeResponse, error) {
	options := requests.RequestOptions{
		Params: map[string]string{
			"f": "json",
		},
		HTTPMethod: requests.GET,
	}

	switch c := coords.(type) {
	case []float64:
		if len(c) == 3 {
			options.Params["location"] = fmt.Sprintf("%f,%f,%f", c[0], c[1], c[2])
		} else if len(c) == 2 {
			options.Params["location"] = fmt.Sprintf("%f,%f", c[0], c[1])
		} else {
			return ReverseGeocodeResponse{}, errors.New("invalid number of coordinates provided")
		}
	case Location:
		options.Params["location"] = fmt.Sprintf("%f,%f", c.Longitude, c.Latitude)
	case Point:
		if c.Z != nil {
			options.Params["location"] = fmt.Sprintf("%f,%f,%f", c.X, c.Y, *c.Z)
		} else {
			options.Params["location"] = fmt.Sprintf("%f,%f", c.X, c.Y)
		}
	default:
		return ReverseGeocodeResponse{}, errors.New("unsupported coordinate type provided")
	}

	endpoint := fmt.Sprintf("%s/reverseGeocode", utils.CleanURL(ARCGIS_ONLINE_GEOCODING_URL))
	response, err := requests.Request(endpoint, &options)
	if err != nil {
		return ReverseGeocodeResponse{}, err
	}
	// Check for server errors, which are returned in the 200 response
	if err := json.Unmarshal([]byte(response), &errorResponse); err == nil && errorResponse.Error.Code != 0 {
		return ReverseGeocodeResponse{}, fmt.Errorf("server error: %s (code %d)", errorResponse.Error.Message, errorResponse.Error.Code)
	}

	var reverseGeocodeResponse ReverseGeocodeResponse
	err = json.Unmarshal([]byte(response), &reverseGeocodeResponse)
	if err != nil {
		return ReverseGeocodeResponse{}, err
	}

	return reverseGeocodeResponse, nil
}
