package vonigo

// ClientResponse this struct represents the structure of a Vonigo response when we ask for clients
type ClientResponse struct {
	Company       string   `json:"company"`
	SecurityToken string   `json:"securityToken"`
	ErrNo         int      `json:"errNo"`
	ErrMsg        string   `json:"errMsg"`
	DateNow       string   `json:"dateNow"`
	Clients       []Client `json:"Clients"`
}
