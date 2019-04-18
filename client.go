package vonigo

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
)

// Client this is a client in the format that Vonigo returns to us
type Client struct {
	ObjectID               string `json:"objectID"`
	Name                   string `json:"name"`
	Title                  string `json:"title"`
	DateCreated            string `json:"dateCreated"`
	DateLastEdited         string `json:"dateLastEdited"`
	DateFirstUsed          string `json:"dateFirstUsed"`
	DateLastUsed           string `json:"dateLastUsed"`
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

// GetEmail return the email address for the given Account
func (c Client) GetEmail() string {
	for _, value := range c.Fields {
		if value.FieldID == fieldMap["email"] {
			return value.FieldValue
		}
	}
	return ""
}

// GetClients - Get all clients
func GetClients(params map[string]string) ([]Client, error) {
	clients := []Client{}
	clientResponse := ClientsResponse{}
	log.Info("get clients!")

	if !hasSecurityToken() {
		err := getSecurityToken()
		if err != nil {
			return clients, err
		}
	}

	log.Info("Params for Clients lookup: ", params)

	mergedParams, _ := getBaseParams("retrieve")

	for i, item := range params {

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
	log.Info(reqURL)
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
