package models

type Customer struct {
	ID                 string `json:"id"`
	Name               string `json:"name"`
	Type               string `json:"type"`
	ServiceDescription string `json:"serviceDescription"`
	PhoneNumber        string `json:"phoneNumber"`
	Email              string `json:"email"`
	Website            string `json:"website"`
	Address            string `json:"address"`
	PICCustomer        string `json:"picCustomer"`
	Location           string `json:"location"`
	Notes              string `json:"notes"`
	CustomerGroupID    string `json:"customerGroupId"`
}
