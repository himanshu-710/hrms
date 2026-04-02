package model

type Address struct {
	ID            int    `json:"id"`
	EmployeeID    int    `json:"employee_id"`
	AddressType   string `json:"address_type"`
	Line1         string `json:"line1"`
	Line2         string `json:"line2"`
	City          string `json:"city"`
	State         string `json:"state"`
	PinCode       string `json:"pin_code"`
	Country       string `json:"country"`
	OwnershipType string `json:"ownership_type"`
}

type AddressDTO struct {
	Line1         string `json:"line1"          validate:"required,min=1,max=200"`
	Line2         string `json:"line2"          validate:"max=200"`
	City          string `json:"city"           validate:"required,max=100"`
	State         string `json:"state"          validate:"required,max=100"`
	PinCode       string `json:"pin_code"       validate:"required,max=10"`
	Country       string `json:"country"        validate:"required,max=100"`
	OwnershipType string `json:"ownership_type" validate:"omitempty,oneof=OWNED RENTED"`
}

type AddressesRequest struct {
	EmployeeID      int        `json:"employee_id"`
	Current         AddressDTO `json:"current"           validate:"required"`
	Permanent       AddressDTO `json:"permanent"`
	CopyFromCurrent bool       `json:"copy_from_current"`
}