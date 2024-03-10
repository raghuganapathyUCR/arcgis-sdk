package utils

import "strings"

// CleanURL trims leading and trailing spaces and removes the trailing slash from a URL string.
func CleanURL(url string) string {
	// Trim leading and trailing spaces
	url = strings.TrimSpace(url)

	// Remove the trailing slash if one was included
	url = strings.TrimSuffix(url, "/")
	return url
}

func JoinStringArray(arr []string) string {
	return strings.Join(arr, ",")
}
