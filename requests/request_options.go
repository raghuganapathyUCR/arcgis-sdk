package requests

import "github.com/raghuganapathyUCR/arcgis-sdk/auth"

type ResponseFormat string
type HTTPMethod string

const (
	GET  HTTPMethod = "GET"
	POST HTTPMethod = "POST"
)
const (
	JSON    ResponseFormat = "json"
	GeoJSON ResponseFormat = "geojson"
	Text    ResponseFormat = "text"
	HTML    ResponseFormat = "html"
	Image   ResponseFormat = "image"
	Zip     ResponseFormat = "zip"
)

type Params map[string]string
type Headers map[string]string

type RequestOptions struct {
	Params           Params                     `json:"params,omitempty"`
	HTTPMethod       HTTPMethod                 `json:"httpMethod,omitempty"`
	RawResponse      bool                       `json:"rawResponse,omitempty"`
	Authentication   auth.AuthenticationManager `json:"authentication,omitempty"`
	HideToken        bool                       `json:"hideToken,omitempty"`
	Portal           string                     `json:"portal,omitempty"`
	Credentials      string                     `json:"credentials,omitempty"`
	MaxURLLength     int                        `json:"maxUrlLength,omitempty"`
	Headers          map[string]string          `json:"headers,omitempty"`
	Signal           interface{}                `json:"signal,omitempty"`
	SuppressWarnings bool                       `json:"suppressWarnings,omitempty"`
}

func getDefaultRequestOptions() RequestOptions {
	return RequestOptions{
		HTTPMethod: POST,
		Params: map[string]string{
			"f": string(JSON),
		},
	}
}
