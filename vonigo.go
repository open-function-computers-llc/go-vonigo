package vonigo

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/Sirupsen/logrus"
)

// Gloabl security token set in securityToken.go
var log *logrus.Logger
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
	log = c.Logger

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

// ***** Clients ***** //

// GetClients - Get all clients
func GetClients(params map[string]string) ([]Client, error) {
	clients := []Client{}
	clientResponse := ClientsResponse{}

	if !hasSecurityToken() {
		err := getSecurityToken()
		if err != nil {
			return clients, err
		}
	}

	params["securityToken"] = securityToken

	reqURL, err := buildURL(baseURL, "api/v1/data/Clients", params)
	if err != nil {
		return nil, err
	}
	var emptyPostValues url.Values
	resp, err := http.PostForm(reqURL, emptyPostValues)

	body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &clientResponse)
	if err != nil {
		return nil, err
	}
	return clientResponse.Clients, nil
}

// GetClient - Get a single client
func GetClient(params map[string]string) (Client, error) {
	client := Client{}
	clientResponse := ClientResponse{}

	if !hasSecurityToken() {
		err := getSecurityToken()
		if err != nil {
			return client, err
		}
	}

	params["securityToken"] = securityToken
	params["method"] = "1"

	reqURL, err := buildURL(baseURL, "api/v1/data/Clients", params)
	if err != nil {
		return client, err
	}
	var emptyPostValues url.Values
	resp, err := http.PostForm(reqURL, emptyPostValues)

	body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &clientResponse)
	if err != nil {
		return client, err
	}
	return clientResponse.Client, nil
}

// ***** Leads ***** //

// GetLeads - Get all leads
func GetLeads(params map[string]string) ([]Client, error) {
	return GetClients(params)
}

// GetLead - Get a single lead
func GetLead(params map[string]string) (Client, error) {
	return GetClient(params)
}
