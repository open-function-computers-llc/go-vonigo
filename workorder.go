package vonigo

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// WorkOrder is a order that has yet to be completed by the service provider.
type WorkOrder struct {
	ObjectID               string `json:"objectID"`
	Name                   string `json:"name"`
	Title                  string `json:"title"`
	DateCreated            string `json:"dateCreated"`
	DateLastEdited         string `json:"dateLastEdited"`
	DateService            string `json:"dateService"`
	CountClients           string `json:"countClients"`
	CountContacts          string `json:"countContacts"`
	CountContactsSecondary string `json:"countContactsSecondary"`
	CountLocations         string `json:"countLocations"`
	CountNotes             string `json:"countNotes"`
	CountDocuments         string `json:"countDocuments"`
	CountEmails            string `json:"countEmails"`
	CountJobs              string `json:"countJobs"`
	CountWorkOrders        string `json:"countWorkOrders"`
	CountInvoices          string `json:"countInvoices"`
	CountReceipts          string `json:"countReceipts"`
	CountPayments          string `json:"countPayments"`
	CountQuotes            string `json:"countQuotes"`
	CountCases             string `json:"countCases"`
	CountTasks             string `json:"countTasks"`
	CountReminders         string `json:"countReminders"`
	CountCrews             string `json:"countCrews"`
	CountCharges           string `json:"countCharges"`
	CountExpenses          string `json:"countExpenses"`
	IsUseContacts          string `json:"isUseContacts"`
	IsUseSecondaryContacts string `json:"isUseSecondaryContacts"`
	IsUseLocations         string `json:"isUseLocations"`
	IsUseNotes             string `json:"isUseNotes"`
	IsUseDocuments         string `json:"isUseDocuments"`
	IsUseEmails            string `json:"isUseEmails"`
	IsUseJobs              string `json:"isUseJobs"`
	IsUseWorkOrders        string `json:"isUseWorkOrders"`
	IsUseInvoices          string `json:"isUseInvoices"`
	IsUseReceipts          string `json:"isUseReceipts"`
	IsUsePayments          string `json:"isUsePayments"`
	IsUseQuotes            string `json:"isUseQuotes"`
	IsUseCases             string `json:"isUseCases"`
	IsUseTasks             string `json:"isUseTasks"`
	IsUseServiceReminder   string `json:"isUseServiceReminder"`
	IsUseCrews             string `json:"isUseCrews"`
	IsUseCharges           string `json:"isUseCharges"`
	IsUseExpenses          string `json:"isUseExpenses"`
	IsUseTips              string `json:"isUseTips"`
	IsUseLabels            string `json:"isUseLabels"`
	IsUseDiscounts         string `json:"isUseDiscounts"`
	IsUseCampaigns         string `json:"isUseCampaigns"`
	IsUsePromos            string `json:"isUsePromos"`
	IsCanEdit              string `json:"isCanEdit"`
	IsCanDelete            string `json:"isCanDelete"`
	IsCanCopy              string `json:"isCanCopy"`
	IsCanConvert           string `json:"isCanConvert"`
	IsCanActivate          string `json:"isCanActivate"`
	IsCanDeactivate        string `json:"isCanDeactivate"`
	IsCanComplete          string `json:"isCanComplete"`
	IsActive               string `json:"isActive"`
	Fields                 []struct {
		FieldID    int    `json:"fieldID"`
		FieldValue string `json:"fieldValue"`
		OptionID   int    `json:"optionID"`
	} `json:"Fields"`
	Relations []struct {
		ObjectTypeID int    `json:"objectTypeID"`
		ObjectID     int    `json:"objectID"`
		Name         string `json:"name"`
		RelationType string `json:"relationType"`
		IsActive     string `json:"isActive"`
	} `json:"Relations"`
}

// GetClientWorkOrders - Get all work orders for a single client/account
func GetClientWorkOrders(clientID string) ([]WorkOrder, error) {
	worders := []WorkOrder{}
	worderResponse := WorkOrderResponse{}
	httpclient := &http.Client{}
	log.Info("get work orders for account!")

	if !hasSecurityToken() {
		err := getSecurityToken()
		if err != nil {
			return worders, err
		}
	}

	params := map[string]string{}
	params["clientID"] = clientID
	params["securityToken"] = securityToken
	params["isCompleteObject"] = "true"

	reqURL, err := buildURL(baseURL, "api/v1/data/WorkOrders/")
	if err != nil {
		return worders, err
	}
	log.Info(reqURL)

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(params)

	req, err := http.NewRequest("POST", reqURL, buf)
	req.Header.Add("Content-Type", "application/json")
	resp, err := httpclient.Do(req)

	body, _ := ioutil.ReadAll(resp.Body)

	err = checkVonigoError(body)

	if err != nil {
		return worders, err
	}

	err = json.Unmarshal(body, &worderResponse)
	if err != nil {
		return worders, err
	}
	return worderResponse.WorkOrders, nil
}

// GetServiceDate - helper function to pull the next service date by converting
// vonigo's UNIX timestamp into a timestamp friendly for humans
func (o WorkOrder) GetServiceDate() string {
	return getStringTime(o.DateService)
}
