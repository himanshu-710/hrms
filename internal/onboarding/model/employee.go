package model

import "time"

type Employee struct {
	ID int `json:"id"`

	EmployeeCode string `json:"employee_code"`

	FirstName   string `json:"first_name"`
	MiddleName  string `json:"middle_name"`
	LastName    string `json:"last_name"`
	DisplayName string `json:"display_name"`

	Gender string `json:"gender"`

	DOB *time.Time `json:"dob"`

	MaritalStatus string `json:"marital_status"`
	BloodGroup    string `json:"blood_group"`

	IsPhysicallyHandicapped bool    `json:"is_physically_handicapped"`
	DisabilityPercentage    float64 `json:"disability_percentage"`

	Nationality string `json:"nationality"`

	WorkEmail    string `json:"work_email"`
	PersonalEmail string `json:"personal_email"`
	MobileNo      string `json:"mobile_no"`
	WorkNo        string `json:"work_no"`
	ResidenceNo   string `json:"residence_no"`

	Designation string `json:"designation"`

	DepartmentID         *int `json:"department_id"`
	ReportingManagerID   *int `json:"reporting_manager_id"`
	DottedLineManagerID  *int `json:"dotted_line_manager_id"`

	DateOfJoining *time.Time `json:"date_of_joining"`

	EmploymentType string `json:"employment_type"`
	ProbationEndDate *time.Time `json:"probation_end_date"`

	Relations map[string]interface{} `json:"relations"`

	IsActive bool `json:"is_active"`

	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
