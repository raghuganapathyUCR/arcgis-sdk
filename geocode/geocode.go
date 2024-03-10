package geocode

import (
	"arcgis-sdk/auth"
	"arcgis-sdk/requests"
	"arcgis-sdk/utils"
	"encoding/json"
	"errors"
	"fmt"
)

const (
	ARCGIS_ONLINE_GEOCODING_URL = "https://geocode.arcgis.com/arcgis/rest/services/World/GeocodeServer"
)

// GeocodeOptions is a struct that defines the options for the GeocodeService.
type Geocoder struct {
	Authentication auth.AuthenticationManager
}

type GeocodeResponse struct {
	SpatialReference struct {
		WKID       int `json:"wkid"`
		LatestWKID int `json:"latestWkid"`
	} `json:"spatialReference"`
	Candidates []struct {
		Address  string `json:"address"`
		Location struct {
			X float64 `json:"x"`
			Y float64 `json:"y"`
		} `json:"location"`
		Score      float64 `json:"score"`
		Attributes struct {
			PlaceName string `json:"PlaceName"`
		} `json:"attributes"`
		Extent struct {
			Xmin float64 `json:"xmin"`
			Ymin float64 `json:"ymin"`
			Xmax float64 `json:"xmax"`
			Ymax float64 `json:"ymax"`
		} `json:"extent"`
	} `json:"candidates"`
}

// GeocodeService is a service that provides geocoding and reverse geocoding capabilities.
type MultiLineAddress struct {
	SingleLine   string   `json:"singleLine,omitempty"`
	Address      string   `json:"address,omitempty"`
	Address2     string   `json:"address2,omitempty"`
	Address3     string   `json:"address3,omitempty"`
	Neighborhood string   `json:"neighborhood,omitempty"`
	City         string   `json:"city,omitempty"`
	Subregion    string   `json:"subregion,omitempty"`
	OutFields    []string `json:"outFields,omitempty"`
	Region       string   `json:"region,omitempty"`
	Postal       string   `json:"postal,omitempty"`
	PostalExt    string   `json:"postalExt,omitempty"`
	CountryCode  string   `json:"countryCode,omitempty"`
	MagicKey     string   `json:"magicKey,omitempty"`
}

func NewGeocoder(auth auth.AuthenticationManager) (*Geocoder, error) {
	if auth == nil {
		return nil, errors.New("authentication manager needed to create a Geocoder")
	}
	return &Geocoder{
		Authentication: auth,
	}, nil
}

func (g *Geocoder) Geocode(address string, isMultiLine bool) (GeocodeResponse, error) {
	var geocodeResponse GeocodeResponse
	var err error
	if isMultiLine {
		geocodeResponse, err = g.findAddressCandidatesMultiLine(address)
	} else {
		geocodeResponse, err = g.findAddressCandidatesSingleLine(address)
	}
	return geocodeResponse, err
}

func (g *Geocoder) findAddressCandidatesSingleLine(address string) (GeocodeResponse, error) {
	url := fmt.Sprintf("%s/findAddressCandidates", utils.CleanURL(ARCGIS_ONLINE_GEOCODING_URL))
	params := map[string]string{
		"singleLine": address,
	}
	return g.requestGeocodeService(url, params)
}

func (g *Geocoder) findAddressCandidatesMultiLine(address string) (GeocodeResponse, error) {
	url := fmt.Sprintf("%s/findAddressCandidates", utils.CleanURL(ARCGIS_ONLINE_GEOCODING_URL))
	// address in this case is a string that represents a JSON object
	// so we need to unmarshal it into a MultiLineAddress struct
	var multiLineAddress MultiLineAddress
	err := json.Unmarshal([]byte(address), &multiLineAddress)
	if err != nil {
		fmt.Println("Error unmarshalling multiline address: ", err)
		return GeocodeResponse{}, err
	}
	params := map[string]string{
		"singleLine":   multiLineAddress.SingleLine,
		"address":      multiLineAddress.Address,
		"address2":     multiLineAddress.Address2,
		"address3":     multiLineAddress.Address3,
		"neighborhood": multiLineAddress.Neighborhood,
		"city":         multiLineAddress.City,
		"subregion":    multiLineAddress.Subregion,
		"region":       multiLineAddress.Region,
		"postal":       multiLineAddress.Postal,
		"postalExt":    multiLineAddress.PostalExt,
		"countryCode":  multiLineAddress.CountryCode,
		"magicKey":     multiLineAddress.MagicKey,
		"outFields":    utils.JoinStringArray(multiLineAddress.OutFields),
	}
	return g.requestGeocodeService(url, params)

}

func (g *Geocoder) requestGeocodeService(url string, params map[string]string) (GeocodeResponse, error) {
	requestOptions := requests.RequestOptions{
		Params:           params,
		HTTPMethod:       requests.GET,
		Authentication:   g.Authentication,
		RawResponse:      true,
		SuppressWarnings: true,
	}
	resp, err := requests.Request(url, &requestOptions)

	if err != nil {
		return GeocodeResponse{}, err
	}
	var geocodeResponse GeocodeResponse
	err = json.Unmarshal(resp, &geocodeResponse)
	if err != nil {
		return GeocodeResponse{}, err
	}
	return geocodeResponse, nil
}
