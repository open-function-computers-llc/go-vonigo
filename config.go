package vonigo

import (
	"errors"

	"github.com/sirupsen/logrus"
)

// Config This is where we will add anything we need to make this library work
// the way we want it to
type Config struct {
	Logger      *logrus.Logger
	BaseURL     string
	AppVersion  string
	Username    string
	Password    string
	Company     string
	FieldMapper map[string]int
}

func (c Config) validate() error {

	// TODO(Kyle) Test to make sure logrus is being passed in correctly
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
