package ipilot

type Event struct {
	AccountID                 int    `json:"Account_id"`
	Object                    string `json:"event_object"`
	Action                    string `json:"event_action"`
	Type                      string `json:"event_type"`
	OrderType                 string `json:"order_type"`
	Module                    string `json:"module"`
	PortCount                 int    `json:"port_count"`
	NumberCount               int    `json:"number_count"`
	NumberAssigned            int    `json:"number_assigned"`
	CallpathCount             int    `json:"call_path_count"`
	PONNumber                 string `json:"pon_number"`
	ExternalCustomerAccountID string `json:"external_customer_account_id"`
	ResellerOwnSKU            string `json:"reseller_own_sku,omitempty"`
	PlatformOwnSKU            string `json:"platform_own_sku,omitempty"`
	SystemSKU                 string `json:"system_sku,omitempty"`
}
