package vonigo

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Franchise - a single franchise listing in the vonigo account
type Franchise struct {
	FranchiseID        int    `json:"franchiseID"`
	OptionID           int    `json:"optionID"`
	FranchiseName      string `json:"franchiseName"`
	GmtOffsetFranchise int    `json:"gmtOffsetFranchise"`
	Sequence           int    `json:"sequence"`
	IsActive           bool   `json:"isActive"`
}

// GetFranchises reach out to vonigo and get a list of all their franchises
func GetFranchises() ([]Franchise, error) {
	f := []Franchise{}

	if !hasSecurityToken() {
		err := getSecurityToken()
		if err != nil {
			return f, err
		}
	}

	params := map[string]interface{}{
		"pageNo":        "1",
		"pageSize":      "100",
		"method":        "0",
		"securityToken": securityToken,
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(params)

	reqURL, err := buildURL(baseURL, "api/v1/resources/franchises/")
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", reqURL, buf)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)

	body, _ := ioutil.ReadAll(resp.Body)

	// debug info
	log.Info("Get all franchise request URL: ", reqURL)
	log.Info("Request Payload: ", params)
	log.Info("Get all franchise request Response: ", string(body))

	type franchiseListResponsePayload struct {
		Company       string      `json:"company"`
		SecurityToken string      `json:"securityToken"`
		ErrNo         int         `json:"errNo"`
		ErrMsg        string      `json:"errMsg"`
		DateNow       string      `json:"dateNow"`
		Franchises    []Franchise `json:"Franchises"`
	}
	franchiseResponse := franchiseListResponsePayload{}

	err = json.Unmarshal(body, &franchiseResponse)
	if err != nil {
		log.Error(err.Error())
		return f, err
	}
	return franchiseResponse.Franchises, nil
}

// ChangeFranchise - change the current session token to work with a different
// franchise
func ChangeFranchise(id int) error {
	franchiseIDAsString := strconv.Itoa(id)
	reqURL, err := buildURL(baseURL, "api/v1/security/session/")
	reqURL += "?securityToken=" + securityToken + "&method=0&franchiseID=" + franchiseIDAsString

	type changeFranchiseResponsePayload struct {
		Company       string `json:"company"`
		SecurityToken string `json:"securityToken"`
		ErrNo         int    `json:"errNo"`
		ErrMsg        string `json:"errMsg"`
		DateNow       string `json:"dateNow"`
		Errors        []struct {
			FieldID   int         `json:"fieldID"`
			FieldName interface{} `json:"fieldName"`
			ErrNo     int         `json:"errNo"`
			ErrMsg    string      `json:"errMsg"`
		} `json:"Errors"`
	}
	responsePayload := changeFranchiseResponsePayload{}
	buf := new(bytes.Buffer)

	req, err := http.NewRequest("POST", reqURL, buf)
	req.Header.Add("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Info(reqURL)
	log.Info(string(body))

	err = json.Unmarshal(body, &responsePayload)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	if responsePayload.ErrMsg != "Franchise changed successfully." {
		return errors.New(responsePayload.Errors[0].ErrMsg)
	}

	return nil
}
