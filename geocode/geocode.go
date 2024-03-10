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
			LocName    string  `json:"Loc_name,omitempty"`
			Status     string  `json:"Status,omitempty"`
			MatchAddr  string  `json:"Match_addr,omitempty"`
			LongLabel  string  `json:"LongLabel,omitempty"`
			ShortLabel string  `json:"ShortLabel,omitempty"`
			AddrType   string  `json:"Addr_type,omitempty"`
			Type       string  `json:"Type,omitempty"`
			PlaceName  string  `json:"PlaceName,omitempty"`
			PlaceAddr  string  `json:"Place_addr,omitempty"`
			Phone      string  `json:"Phone,omitempty"`
			URL        string  `json:"URL,omitempty"`
			Rank       int     `json:"Rank,omitempty"`
			AddBldg    string  `json:"AddBldg,omitempty"`
			AddNum     string  `json:"AddNum,omitempty"`
			AddNumFrom string  `json:"AddNumFrom,omitempty"`
			AddNumTo   string  `json:"AddNumTo,omitempty"`
			AddRange   string  `json:"AddRange,omitempty"`
			Side       string  `json:"Side,omitempty"`
			StPreDir   string  `json:"StPreDir,omitempty"`
			StPreType  string  `json:"StPreType,omitempty"`
			StName     string  `json:"StName,omitempty"`
			StType     string  `json:"StType,omitempty"`
			StDir      string  `json:"StDir,omitempty"`
			BldgType   string  `json:"BldgType,omitempty"`
			BldgName   string  `json:"BldgName,omitempty"`
			LevelType  string  `json:"LevelType,omitempty"`
			LevelName  string  `json:"LevelName,omitempty"`
			UnitType   string  `json:"UnitType,omitempty"`
			UnitName   string  `json:"UnitName,omitempty"`
			SubAddr    string  `json:"SubAddr,omitempty"`
			StAddr     string  `json:"StAddr,omitempty"`
			Block      string  `json:"Block,omitempty"`
			Sector     string  `json:"Sector,omitempty"`
			Nbrhd      string  `json:"Nbrhd,omitempty"`
			District   string  `json:"District,omitempty"`
			City       string  `json:"City,omitempty"`
			MetroArea  string  `json:"MetroArea,omitempty"`
			Subregion  string  `json:"Subregion,omitempty"`
			Region     string  `json:"Region,omitempty"`
			RegionAbbr string  `json:"RegionAbbr,omitempty"`
			Territory  string  `json:"Territory,omitempty"`
			Zone       string  `json:"Zone,omitempty"`
			Postal     string  `json:"Postal,omitempty"`
			PostalExt  string  `json:"PostalExt,omitempty"`
			Country    string  `json:"Country,omitempty"`
			CntryName  string  `json:"CntryName,omitempty"`
			LangCode   string  `json:"LangCode,omitempty"`
			Distance   float64 `json:"Distance,omitempty"`
			X          float64 `json:"X,omitempty"`
			Y          float64 `json:"Y,omitempty"`
			DisplayX   float64 `json:"DisplayX,omitempty"`
			DisplayY   float64 `json:"DisplayY,omitempty"`
			Xmin       float64 `json:"Xmin,omitempty"`
			Xmax       float64 `json:"Xmax,omitempty"`
			Ymin       float64 `json:"Ymin,omitempty"`
			Ymax       float64 `json:"Ymax,omitempty"`
			ExInfo     string  `json:"ExInfo,omitempty"`
		} `json:"attributes,omitempty"`
		Extent struct {
			Xmin float64 `json:"xmin"`
			Ymin float64 `json:"ymin"`
			Xmax float64 `json:"xmax"`
			Ymax float64 `json:"ymax"`
		} `json:"extent"`
	} `json:"candidates"`
}

// GeocodeService is a service that provides geocoding and reverse geocoding capabilities.
type GeocodeRequestOptions struct {
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

func (g *Geocoder) Geocode(address GeocodeRequestOptions) (GeocodeResponse, error) {
	var geocodeResponse GeocodeResponse
	var err error

	if err != nil {
		fmt.Println("Error marshalling address: ", err)
		return GeocodeResponse{}, err
	}
	geocodeResponse, err = g.findAddressCandidatesMultiLine(address)

	return geocodeResponse, err
}

func (g *Geocoder) findAddressCandidatesMultiLine(address GeocodeRequestOptions) (GeocodeResponse, error) {
	url := fmt.Sprintf("%s/findAddressCandidates", utils.CleanURL(ARCGIS_ONLINE_GEOCODING_URL))

	params := map[string]string{
		"singleLine":   address.SingleLine,
		"address":      address.Address,
		"address2":     address.Address2,
		"address3":     address.Address3,
		"neighborhood": address.Neighborhood,
		"city":         address.City,
		"subregion":    address.Subregion,
		"region":       address.Region,
		"postal":       address.Postal,
		"postalExt":    address.PostalExt,
		"countryCode":  address.CountryCode,
		"magicKey":     address.MagicKey,
		"outFields":    utils.JoinStringArray(address.OutFields),
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

	// print the response to understand how the server responds when errors occur
	if err != nil {
		return GeocodeResponse{}, err
	}

	var errorResponse struct {
		Error struct {
			Code    int      `json:"code"`
			Message string   `json:"message"`
			Details []string `json:"details"`
		} `json:"error"`
	}
	if err := json.Unmarshal([]byte(resp), &errorResponse); err == nil && errorResponse.Error.Code != 0 {
		return GeocodeResponse{}, fmt.Errorf("server error: %s (code %d)", errorResponse.Error.Message, errorResponse.Error.Code)
	}
	// Continue with the error check
	if errorResponse.Error.Code == 498 {
		return GeocodeResponse{}, errors.New("invalid token")
	}

	// Unmarshal the response into the GeocodeResponse struct
	var geocodeResponse GeocodeResponse
	if err := json.Unmarshal([]byte(resp), &geocodeResponse); err != nil {
		return GeocodeResponse{}, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return geocodeResponse, nil
}
