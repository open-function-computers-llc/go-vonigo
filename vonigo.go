package vonigo

import (
	"github.com/Sirupsen/logrus"
)

var logger *logrus.Logger
var isInitialized bool
var securityToken string

// Init Here is where we will
func Init(c Config) {
	securityToken = c.SecurityToken
	logger = c.Logger

	logger.Info("The vonigo package has been successfully initialized")
	isInitialized = true
}
