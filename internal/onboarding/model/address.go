package model

type Address struct {
	Street     string `json:"street"`
	City       string `json:"city"`
	State      string `json:"state"`
	Country    string `json:"country"`
	PostalCode string `json:"postal_code"`
}

type AddressesRequest struct {
	EmployeeID int `json:"employee_id"`

	Current   Address `json:"current"`
	Permanent Address `json:"permanent"`

	CopyFromCurrent bool `json:"copy_from_current"`
}
