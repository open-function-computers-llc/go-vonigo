package vonigo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
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
	// this is for testing only
	if baseURL == "https://example.com" {
		securityToken = "fakeSecurityToken"
		return nil
	}

	// https://stackoverflow.com/questions/24493116/how-to-send-a-post-request-in-go
	// all the paramaters that vonigo wants are in the URL, even though we are making a POST request
	formValues := url.Values{}
	response, err := http.PostForm(getTokenURL(), formValues)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	fmt.Println(string(body))

	return nil
}

func getTokenURL() string {
	return baseURL + "/api/v1/security/login/?appVersion=" + appVersion + "&company=" + company + "&password=" + password + "&userName=" + username
}
