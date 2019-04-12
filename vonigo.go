package vonigo

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

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

	mergedParams, _ := getBaseParams("retrieve")

	for i, item := range params {

		if i == "minutes" {
			start, end := getMinutesAgo(item)
			mergedParams["dateStart"] = start
			mergedParams["dateEnd"] = end
		}
		mergedParams[i] = item
	}

	reqURL, reqParams, err := buildURL(baseURL, "api/v1/data/Clients", mergedParams)
	if err != nil {
		return nil, err
	}

	resp, err := http.PostForm(reqURL, reqParams)

	body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &clientResponse)
	if err != nil {
		return nil, err
	}
	return clientResponse.Clients, nil
}

// GetClient - Get a single client
func GetClient(id int) (Client, error) {
	stringID := strconv.Itoa(id)
	client := Client{}
	clientResponse := ClientResponse{}

	if !hasSecurityToken() {
		err := getSecurityToken()
		if err != nil {
			return client, err
		}
	}
	params, _ := getBaseParams("retrieve")
	params["objectID"] = stringID

	reqURL, urlValues, err := buildURL(baseURL, "api/v1/data/Clients", params)
	if err != nil {
		return client, err
	}
	fmt.Println(reqURL)
	resp, err := http.PostForm(reqURL, urlValues)

	body, _ := ioutil.ReadAll(resp.Body)

	err = checkVonigoError(body)

	if err != nil {
		return client, err
	}

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
func GetLead(id int) (Client, error) {
	return GetClient(id)
}
