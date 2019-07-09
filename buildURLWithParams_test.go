package vonigo

import (
	"testing"
)

func TestCanBuildURLWithParams(t *testing.T) {

	// setup some local variables

	baseurl := "example.com"
	pathurl := "data/url"
	paramsurl := map[string]string{
		"query1": "query1value",
		"query2": "query2value",
	}

	// pass those locals vars into buildURL

	url1, _, err := buildURLWithParams(baseurl, pathurl, paramsurl)

	if err != nil {
		t.Error("URL Failed to build")
	}
	// Vonigo requires a trailing slash between the path and query params
	if url1 != "https://example.com/data/url/?query1=query1value&query2=query2value" {
		t.Error("BuildURL failed to correctly build the URL. Expected: " + "https://example.com/data/url/?query1=query1value&query2=query2value" + " Got " + url1)
	}
}

func TestBaseUrlWithParamsCantHaveProtocal(t *testing.T) {

	baseurl := "https://example.com"
	pathurl := ""
	paramsurl := map[string]string{}

	_, _, err := buildURLWithParams(baseurl, pathurl, paramsurl)

	if err == nil && err.Error() != "Do not pass the protocal with the url, only the host" {
		t.Error("Failed to catch incorrect base URL format")
	}
}
