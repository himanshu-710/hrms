package model

import "time"

type Employee struct {
	ID int `json:"id"`

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`

	DOB        *time.Time `json:"dob"`
	Gender     string     `json:"gender"`
	BloodGroup string     `json:"blood_group"`

	PersonalEmail string `json:"personal_email"`
	MobileNo      string `json:"mobile_no"`
	WorkNo        string `json:"work_no"`

	Department string `json:"department"`

	DateOfJoining *time.Time `json:"date_of_joining"`

	EmploymentContextRole string `json:"employment_context_role"`

	Relations map[string]interface{} `json:"relations"`

	IsActive bool `json:"is_active"`

	CreatedAt time.Time `json:"created_at"`
}
