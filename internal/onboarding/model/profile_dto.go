package model
type OnboardingProfileDTO struct {
	Employee  Employee
	Education []Education
	Experience []Experience
	Addresses []Address
	Documents []EmployeeDocument
	Assets []EmployeeAsset
}