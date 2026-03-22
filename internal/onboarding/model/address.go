package model
type Address struct {
	ID int `json:"id"`

	EmployeeID int `json:"employee_id"`

	AddressType string `json:"address_type"`

	Street string `json:"street"`
	City   string `json:"city"`
	State  string `json:"state"`
	Country string `json:"country"`

	PostalCode string `json:"postal_code"`
}