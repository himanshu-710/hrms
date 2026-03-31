package model

type OnboardingProfileDTO struct {
	Employee   Employee           `json:"employee"`
	Education  []Education        `json:"education"`
	Experience []Experience       `json:"experience"`
	Addresses  []Address          `json:"addresses"`
	Documents  []EmployeeDocument `json:"documents"`
	Assets     []EmployeeAsset    `json:"assets"`

	Identity  []IdentityDocument     `json:"identity"`  
	Relations map[string]interface{} `json:"relations"` 
}

type PrimaryDetailsRequest struct {
	FirstName     string `json:"first_name"`
	MiddleName    string `json:"middle_name"`
	LastName      string `json:"last_name"`
	DisplayName   string `json:"display_name"`
	Gender        string `json:"gender"`
	DOB           string `json:"dob"`
	MaritalStatus string `json:"marital_status"`
	BloodGroup    string `json:"blood_group"`
	Nationality   string `json:"nationality"`
}
type ContactRequest struct {
	PersonalEmail string `json:"personal_email"`
	MobileNo      string `json:"mobile_no"`
	WorkNo        string `json:"work_no"`
	ResidenceNo   string `json:"residence_no"`
}
type RelationsRequest struct {
	Mother struct {
		Name       string `json:"name"`
		DOB        string `json:"dob"`
		Occupation string `json:"occupation"`
		Contact    string `json:"contact"`
	} `json:"mother"`

	Father struct {
		Name       string `json:"name"`
		DOB        string `json:"dob"`
		Occupation string `json:"occupation"`
		Contact    string `json:"contact"`
	} `json:"father"`

	Spouse struct {
		Name    string `json:"name"`
		DOB     string `json:"dob"`
		Contact string `json:"contact"`
	} `json:"spouse"`

	Children []struct {
		Name string `json:"name"`
		DOB  string `json:"dob"`
	} `json:"children"`
}
type EducationRequest struct {
	EmployeeID       int     `json:"employee_id"`
	Degree           string  `json:"degree"`
	Branch           string  `json:"branch"`
	University       string  `json:"university"`
	CGPAOrPct        float64 `json:"cgpa_or_pct"`
	YearOfJoining    int     `json:"year_of_joining"`
	YearOfCompletion int     `json:"year_of_completion"`
}

type ExperienceRequest struct {
	EmployeeID     int     `json:"employee_id"`
	CompanyName    string  `json:"company_name"`
	Designation    string  `json:"designation"`
	EmploymentType string  `json:"employment_type"`
	StartDate      string  `json:"start_date"`
	EndDate        *string `json:"end_date"`
	IsCurrent      bool    `json:"is_current"`
	Industry       string  `json:"industry"`
	Description    string  `json:"description"`
}

type AssignAssetRequest struct {
    EmployeeID    int    `json:"employee_id"`
    AssetType     string `json:"asset_type"`
    AssetName     string `json:"asset_name"`
    AssetCategory string `json:"asset_category"`
    SerialNo      string `json:"serial_no"`
    Condition     string `json:"condition"`
    AssignedBy    int    `json:"assigned_by"`
    Notes         string `json:"notes"`
}

type CreateEmployeeRequest struct {
    FirstName  string `json:"first_name"`
    LastName   string `json:"last_name"`
    Email      string `json:"personal_email"`
    Department string `json:"department"`
}

type VerifyDocumentRequest struct {
    Status string `json:"status"`
    Note   string `json:"note"`
}