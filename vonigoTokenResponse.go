package vonigo

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

type tokenResponse struct {
	Company       string `json:"company"` // this is not necessarily what company you are working with, but rather "vonigo"
	SecurityToken string `json:"securityToken"`
	ErrNo         int    `json:"errNo"`
	ErrMsg        string `json:"errMsg"`
	DateNow       string `json:"dateNow"`
	Session       struct {
		CompanyID                string `json:"companyID"`
		FranchiseID              string `json:"franchiseID"`
		FranchiseName            string `json:"franchiseName"`
		OfficeID                 string `json:"officeID"`
		RouteID                  string `json:"routeID"`
		DatabaseVersionID        string `json:"databaseVersionID"`
		ScheduleValidationModeID string `json:"scheduleValidationModeID"`
		FirstName                string `json:"firstName"`
		LastName                 string `json:"lastName"`
		UserPick                 string `json:"userPick"`
		UserPic                  string `json:"userPic"`
		ScheduleTimeStart        string `json:"scheduleTimeStart"`
		ScheduleTimeEnd          string `json:"scheduleTimeEnd"`
		ScheduleViewType         string `json:"scheduleViewType"`
		GmtOffsetFranchise       string `json:"gmtOffsetFranchise"`
		GmtOffsetUser            string `json:"gmtOffsetUser"`
		ScheduleDayMin           string `json:"scheduleDayMin"`
		ScheduleDayMax           string `json:"scheduleDayMax"`
		ScheduleDateCurrent      string `json:"scheduleDateCurrent"`
	} `json:"Session"`
	SecurityGroups []struct {
		GroupID    int    `json:"groupID"`
		Name       string `json:"name"`
		IsSelected bool   `json:"isSelected"`
	} `json:"securityGroups"`
}

func getSecurityToken() error {
	params := map[string]string{
		"appVersion": appVersion,
		"username":   username,
		"password":   password,
		"company":    company,
	}

	log.Info(params)
	reqURL, _, err := buildURLWithParams(baseURL, "api/v1/security/login/", params)
	if err != nil {
		return err
	}

	log.Info(reqURL)
	resp, err := http.Get(reqURL)
	if err != nil {
		return err
	}
	body, _ := ioutil.ReadAll(resp.Body)

	err = checkVonigoError(body)
	if err != nil {
		return err
	}

	tknresp := tokenResponse{}

	err = json.Unmarshal(body, &tknresp)
	if err != nil {
		return err
	}
	log.Info("This is the response from getting the token:")
	log.Info(tknresp)
	securityToken = tknresp.SecurityToken
	log.Info(securityToken)

	return nil
}

func getBaseParams(action string) (map[string]interface{}, error) {
	value := map[string]interface{}{
		"securityToken": securityToken,
	}

	validActions := map[string]interface{}{
		"retrieve": "1",
		"update":   "2",
		"create":   "3",
	}

	if _, ok := validActions[action]; !ok {
		return value, errors.New("Invalid action")
	}

	// value["method"] = validActions[action]
	return value, nil
}
