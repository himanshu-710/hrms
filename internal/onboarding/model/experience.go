package model

type Experience struct {
	ID int `json:"id"`

	EmployeeID int `json:"employee_id"`

	CompanyName string `json:"company_name"`
	Designation string `json:"designation"`

	EmploymentType string `json:"employment_type"` 

	StartDate string `json:"start_date"`
	EndDate   *string `json:"end_date"` 

	IsCurrent bool `json:"is_current"`

	Industry    string `json:"industry"`
	Description string `json:"description"`
}
