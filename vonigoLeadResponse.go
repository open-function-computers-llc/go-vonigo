package vonigo

// LeadsResponse this struct represents the structure of a Vonigo response when we ask for clients
type LeadsResponse struct {
	Company       string `json:"company"`
	SecurityToken string `json:"securityToken"`
	ErrNo         int    `json:"errNo"`
	ErrMsg        string `json:"errMsg"`
	DateNow       string `json:"dateNow"`
	Leads         []Lead `json:"Clients"`
}

// LeadResponse this stuct is the response when requesting a single client details
type LeadResponse struct {
	Company       string `json:"company"`
	SecurityToken string `json:"securityToken"`
	ErrNo         int    `json:"errNo"`
	ErrMsg        string `json:"errMsg"`
	DateNow       string `json:"dateNow"`
	Lead          Lead   `json:"Client"`
}
