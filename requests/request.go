package requests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func internalRequestOptions(url string, requestOptions *RequestOptions) (RequestOptions, error) {
	// Get default request options
	defaultOptions := getDefaultRequestOptions()

	// Merge default options with provided options
	mergedOptions := RequestOptions{
		HTTPMethod:       requestOptions.HTTPMethod,
		RawResponse:      requestOptions.RawResponse,
		Authentication:   requestOptions.Authentication,
		HideToken:        requestOptions.HideToken,
		Portal:           requestOptions.Portal,
		Credentials:      requestOptions.Credentials,
		MaxURLLength:     requestOptions.MaxURLLength,
		Signal:           requestOptions.Signal,
		SuppressWarnings: requestOptions.SuppressWarnings,
		Params:           make(Params),
		Headers:          make(Headers),
	}

	// Merge headers and params maps
	for k, v := range defaultOptions.Headers {
		mergedOptions.Headers[k] = v
	}
	for k, v := range requestOptions.Headers {
		mergedOptions.Headers[k] = v
	}
	for k, v := range requestOptions.Params {
		mergedOptions.Params[k] = v
	}
	for k, v := range defaultOptions.Params {
		mergedOptions.Params[k] = v
	}

	return mergedOptions, nil

}

func Request(url string, requestOptions *RequestOptions) ([]byte, error) {
	options, err := internalRequestOptions(url, requestOptions)
	if err != nil {
		return []byte{}, err
	}
	resp, err := makeRequest(url, options)
	if err != nil {
		return []byte{}, err
	}
	return resp, nil
}

func makeRequest(requestURL string, options RequestOptions) ([]byte, error) {
	var resp *http.Response
	var err error

	switch options.HTTPMethod {
	case "GET":
		// Construct the URL with query parameters - merge default and provided params
		params := url.Values{}
		for k, v := range options.Params {
			if v != "" {
				params.Add(k, v)
			}
		}
		// add token to the request
		if options.Authentication != nil {
			token, err := options.Authentication.GetToken(requestURL)
			if err != nil {
				return []byte{}, fmt.Errorf("%v", err)
			}
			params.Add("token", token)
		}

		requestURL = fmt.Sprintf("%s?%s", requestURL, params.Encode())
		// Make the GET request
		resp, err = http.Get(requestURL)

		if err != nil {
			return []byte{}, fmt.Errorf("error making GET request: %v", err)
		}

		defer resp.Body.Close()

	case "POST":
		// Marshal the request body as JSON
		jsonBody, err := json.Marshal(options.Params)
		if err != nil {
			return []byte{}, fmt.Errorf("error marshaling request body: %v", err)
		}

		// Make the POST request
		resp, err = http.Post(requestURL, "application/json", bytes.NewBuffer(jsonBody))
		if err != nil {
			return []byte{}, fmt.Errorf("error making POST request: %v", err)
		}
		defer resp.Body.Close()

	default:
		return []byte{}, fmt.Errorf("unsupported HTTP method: %s", options.HTTPMethod)
	}

	// Read the response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, fmt.Errorf("error reading response body: %v", err)
	}
	return body, nil
}
