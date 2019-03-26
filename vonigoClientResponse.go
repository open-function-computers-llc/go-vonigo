package vonigo

// ClientsResponse this struct represents the structure of a Vonigo response when we ask for clients
type ClientsResponse struct {
	Company       string   `json:"company"`
	SecurityToken string   `json:"securityToken"`
	ErrNo         int      `json:"errNo"`
	ErrMsg        string   `json:"errMsg"`
	DateNow       string   `json:"dateNow"`
	Clients       []Client `json:"Clients"`
}

// ClientResponse this stuct is the response when requesting a single client details
type ClientResponse struct {
	Company       string `json:"company"`
	SecurityToken string `json:"securityToken"`
	ErrNo         int    `json:"errNo"`
	ErrMsg        string `json:"errMsg"`
	DateNow       string `json:"dateNow"`
	Client        Client `json:"Client"`
}
