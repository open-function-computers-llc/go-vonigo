package vonigo

import (
	"errors"
	"net/url"
	"strings"
)

func buildURL(base string, path string, qparams map[string]string) (string, url.Values, error) {
	v := url.Values{}

	if strings.Contains(base, "http") {
		return "", v, errors.New("Do not pass the protocal with the url, only the host")
	}

	for key, value := range qparams {
		v.Add(key, value)
	}

	// Vonigo requires a trailing '/' on the path of requests
	lastCharacterOfPath := path[len(path)-1:]

	if lastCharacterOfPath != "/" {
		path = path + "/"
	}

	u := &url.URL{
		Scheme:   "https",
		Host:     base,
		Path:     path,
		RawQuery: v.Encode(),
	}

	return u.String(), v, nil

}
