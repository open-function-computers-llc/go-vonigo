package vonigo

import (
	"testing"
)

func TestCanBuildURL(t *testing.T) {

	// setup some local variables

	baseurl := "example.com"
	pathurl := "data/url"
	paramsurl := map[string]string{
		"query1": "query1value",
		"query2": "query2value",
	}

	// pass those locals vars into buildURL

	url1, err := buildURL(baseurl, pathurl, paramsurl)

	if err != nil {
		t.Error("URL Failed to build")
	}

	if url1 != "https://example.com/data/url?query1=query1value&query2=query2value" {
		t.Error("BuildURL failed to correctly build the URL")
	}
}

func TestBaseUrlCantHaveProtocal(t *testing.T) {

	baseurl := "https://example.com"
	pathurl := ""
	paramsurl := map[string]string{}

	_, err := buildURL(baseurl, pathurl, paramsurl)

	if err == nil && err.Error() != "Do not pass the protocal with the url, only the host" {
		t.Error("Failed to catch incorrect base URL format")
	}
}
