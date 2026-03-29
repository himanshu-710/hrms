package model

type Address struct {
	ID            int    `json:"id"`
	EmployeeID    int    `json:"employee_id"`
	AddressType   string `json:"address_type"`   // ✅ ADD THIS
	Line1         string `json:"line1"`
	Line2         string `json:"line2"`
	City          string `json:"city"`
	State         string `json:"state"`
	PinCode       string `json:"pin_code"`
	Country       string `json:"country"`
	OwnershipType string `json:"ownership_type"`
}

type AddressesRequest struct {
	EmployeeID int `json:"employee_id"`

	Current   Address `json:"current"`
	Permanent Address `json:"permanent"`

	CopyFromCurrent bool `json:"copy_from_current"`
}
