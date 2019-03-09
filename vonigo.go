package vonigo

import (
	"github.com/Sirupsen/logrus"
)

// Logger this is where we will log all interactions
var Logger logrus.Logger

var isInitialized bool

var securityToken string

// Init Here is where we will
func Init(c Config) {
	securityToken = c.SecurityToken

	isInitialized = true
}
