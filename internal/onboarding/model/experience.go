package model
type Experience struct {
	ID int `json:"id"`

	EmployeeID int `json:"employee_id"`

	Company string `json:"company"`
	Role    string `json:"role"`

	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`

	CurrentlyWorking bool `json:"currently_working"`
}