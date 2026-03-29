package model

type Education struct {
	ID int `json:"id"`

	EmployeeID int `json:"employee_id"`

	Degree string `json:"degree"`
	Branch string `json:"branch"`

	University string `json:"university"`

	CGPAOrPct float64 `json:"cgpa_or_pct"`

	YearOfJoining    int `json:"year_of_joining"`
	YearOfCompletion int `json:"year_of_completion"`

	CertificateURL string `json:"certificate_url"` 
}
