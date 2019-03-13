package vonigo

import (
	"crypto/md5"
	"encoding/hex"
)

var securityToken string
var baseURL string
var appVersion string
var username string
var password string
var company string
var isInitialized bool

// Init Check for all the required package level variables, and get a vonigo security token
func Init(c Config) error {
	err := c.validate()
	if err != nil {
		return err
	}
	baseURL = c.BaseURL
	appVersion = c.AppVersion
	username = c.Username

	// vonigo wants the MD5 hash of the password, not the raw text password
	rawPassword := c.Password
	hash := md5.Sum([]byte(rawPassword))
	password = hex.EncodeToString(hash[:])

	company = c.Company

	err = getSecurityToken()
	if err != nil {
		return err
	}

	isInitialized = true
	return nil
}
