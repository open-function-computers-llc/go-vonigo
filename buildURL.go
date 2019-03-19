package vonigo

import (
	"errors"
	"net/url"
	"strings"
)

func buildURL(base string, path string, qparams map[string]string) (string, error) {

	if strings.Contains(base, "http") {
		return "", errors.New("Do not pass the protocal with the url, only the host")
	}

	v := url.Values{}

	for key, value := range qparams {
		v.Add(key, value)
	}

	u := &url.URL{
		Scheme:   "https",
		Host:     base,
		Path:     path,
		RawQuery: v.Encode(),
	}

	return u.String(), nil

}
