package vonigo

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// Lead this is a client in the format that Vonigo returns to us
type Lead struct {
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

// GetEmail - Get email for lead
func (c Lead) GetEmail() string {
	for _, value := range c.Fields {
		if value.FieldID == fieldMap["email"] {
			return value.FieldValue
		}
	}
	return ""
}

// GetPhoneNumber - Get email for lead
func (c Lead) GetPhoneNumber() string {
	for _, value := range c.Fields {
		if value.FieldID == fieldMap["phone"] {
			return value.FieldValue
		}
	}
	return ""
}

// GetLeadType - Get type of lead: Residential, Commercial, or ""
func (c Lead) GetLeadType() string {
	for _, value := range c.Fields {
		if value.FieldID == fieldMap["type"] {
			if value.OptionID == 59 {
				return "Residential"
			}

			if value.OptionID == 60 {
				return "Commercial"
			}

			return ""
		}
	}
	return ""
}

// GetNextWorkOrderDate - return the string date of the next work order for a lead
func (c Lead) GetNextWorkOrderDate() string {

	clientID := c.ObjectID
	workOrders, err := GetClientWorkOrders(clientID)
	if err != nil {
		return ""
	}

	if len(workOrders) == 0 {
		return ""
	}

	if len(workOrders) == 1 {
		order := workOrders[0]
		return getStringTime(order.DateService)
	}

	currentTime := int(time.Now().Unix())
	timeDiff := 0
	serviceTime := ""
	for _, o := range workOrders {
		service, _ := strconv.Atoi(o.DateService)
		// Initial set
		if timeDiff == 0 && service-currentTime > 0 {
			timeDiff = service - currentTime
			serviceTime = o.DateService
			continue
		}

		// If the order is "closer" to current time and also not in the past
		if service-currentTime > 0 && service-currentTime < timeDiff {
			timeDiff = service - currentTime
			serviceTime = o.DateService
		}
	}
	t := getStringTime(serviceTime)
	return t
}

// GetLeads - Get all leads
func GetLeads(params map[string]string) ([]Lead, error) {
	leads := []Lead{}
	leadResponse := LeadsResponse{}
	client := &http.Client{}
	log.Info("get leads!")

	if !hasSecurityToken() {
		err := getSecurityToken()
		if err != nil {
			return leads, err
		}
	}

	mergedParams, _ := getBaseParams("retrieve")
	mergedParams["dateMode"] = "1"

	for i, item := range params {
		mergedParams[i] = item
	}

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(mergedParams)

	reqURL, err := buildURL(baseURL, "api/v1/data/Leads")
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", reqURL, buf)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)

	body, _ := ioutil.ReadAll(resp.Body)

	err = json.Unmarshal(body, &leadResponse)
	if err != nil {
		log.Error(err.Error())
		return nil, err
	}
	return leadResponse.Leads, nil
}

// GetLead - Get a single client
func GetLead(id int) (Lead, error) {
	stringID := strconv.Itoa(id)
	lead := Lead{}
	leadResponse := LeadResponse{}
	client := &http.Client{}

	if !hasSecurityToken() {
		err := getSecurityToken()
		if err != nil {
			return lead, err
		}
	}
	params, _ := getBaseParams("retrieve")
	params["objectID"] = stringID

	reqURL, err := buildURL(baseURL, "api/v1/data/Leads")
	if err != nil {
		return lead, err
	}
	log.Info(reqURL)

	buf := new(bytes.Buffer)
	json.NewEncoder(buf).Encode(params)

	req, err := http.NewRequest("POST", reqURL, buf)
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)

	body, _ := ioutil.ReadAll(resp.Body)

	err = checkVonigoError(body)

	if err != nil {
		return lead, err
	}

	err = json.Unmarshal(body, &leadResponse)
	if err != nil {
		return lead, err
	}
	return leadResponse.Lead, nil
}
