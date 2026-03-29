package model

type DashboardRowDTO struct {
	EmployeeID        int      `json:"employee_id"`
	Name              string   `json:"name"`
	DaysSinceJoining  int      `json:"days_since_joining"`
	CompletionPct     int      `json:"completion_pct"`
	IncompleteSections []string `json:"incomplete_sections"`
}