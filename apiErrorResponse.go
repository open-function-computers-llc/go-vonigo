package vonigo

type errorResponse struct {
	Company       string `json:"company"`
	SecurityToken string `json:"securityToken"`
	ErrNo         int    `json:"errNo"`
	ErrMsg        string `json:"errMsg"`
	DateNow       string `json:"dateNow"`
	Errors        []struct {
		FieldID   int    `json:"fieldID"`
		FieldName string `json:"fieldName"`
		ErrNo     int    `json:"errNo"`
		ErrMsg    string `json:"errMsg"`
	} `json:"Errors"`
}
