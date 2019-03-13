package vonigo

import "errors"

// Config This is where we will add anything we need to make this library work
// the way we want it to
type Config struct {
	BaseURL    string
	AppVersion string
	Username   string
	Password   string
	Company    string
}

func (c Config) validate() error {
	if c.BaseURL == "" {
		return errors.New("BaseURL is required")
	}

	if c.AppVersion == "" {
		return errors.New("AppVersion is required")
	}

	if c.Username == "" {
		return errors.New("Username is required")
	}

	if c.Password == "" {
		return errors.New("Password is required")
	}

	if c.Company == "" {
		return errors.New("Company is required")
	}

	return nil
}
