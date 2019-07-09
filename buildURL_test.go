package vonigo

import (
	"testing"
)

func TestCanBuildURL(t *testing.T) {

	// setup some local variables

	baseurl := "example.com"
	pathurl := "data/url"

	// pass those locals vars into buildURL

	url1, err := buildURL(baseurl, pathurl)

	if err != nil {
		t.Error("URL Failed to build")
	}
	// Vonigo requires a trailing slash between the path and query params
	if url1 != "https://example.com/data/url/" {
		t.Error("BuildURL failed to correctly build the URL. Expected: " + "https://example.com/data/url/" + " Got " + url1)
	}
}

func TestBaseUrlCantHaveProtocal(t *testing.T) {

	baseurl := "https://example.com"
	pathurl := ""

	_, err := buildURL(baseurl, pathurl)

	if err == nil && err.Error() != "Do not pass the protocal with the url, only the host" {
		t.Error("Failed to catch incorrect base URL format")
	}
}
