package vonigo

// WorkOrderResponse this struct represents the structure of a Vonigo response when we ask for work orders
type WorkOrderResponse struct {
	Company       string      `json:"company"`
	SecurityToken string      `json:"securityToken"`
	ErrNo         int         `json:"errNo"`
	ErrMsg        string      `json:"errMsg"`
	DateNow       string      `json:"dateNow"`
	WorkOrders    []WorkOrder `json:"WorkOrders"`
}
