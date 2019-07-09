package vonigo

import (
	"errors"
	"net/url"
	"strings"
)

func buildURL(base string, path string) (string, error) {

	if strings.Contains(base, "http") {
		return "", errors.New("Do not pass the protocal with the url, only the host")
	}

	// Vonigo requires a trailing '/' on the path of requests
	lastCharacterOfPath := path[len(path)-1:]

	if lastCharacterOfPath != "/" {
		path = path + "/"
	}

	u := &url.URL{
		Scheme: "https",
		Host:   base,
		Path:   path,
	}

	return u.String(), nil
}
