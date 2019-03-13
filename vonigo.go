package vonigo

import (
	"github.com/Sirupsen/logrus"
)

var logger *logrus.Logger
var securityToken string
var baseURL string
var appVersion string
var username string
var password string

// Init Here is where we will
func Init(c Config) error {
	err := c.validate()
	if err != nil {
		return err
	}
	baseURL = c.BaseURL
	appVersion = c.AppVersion
	username = c.Username
	password = c.Password

	return nil
}
