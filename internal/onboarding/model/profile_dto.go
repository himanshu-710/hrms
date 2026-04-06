package model

type CreateEmployeeRequest struct {
	FirstName             string `json:"first_name"              validate:"required,min=1,max=100"`
	LastName              string `json:"last_name"               validate:"required,min=1,max=100"`
	Email                 string `json:"personal_email"          validate:"required,email"`
	Department            string `json:"department"`
	EmploymentContextRole string `json:"employment_context_role" validate:"omitempty,oneof=EMPLOYEE HR "`
}

type VerifyDocumentRequest struct {
	Status string `json:"status" validate:"required,oneof=VERIFIED REJECTED"`
	Note   string `json:"note"`
}

type OnboardingProfileDTO struct {
	Employee   Employee               `json:"employee"`
	Education  []Education            `json:"education"`
	Experience []Experience           `json:"experience"`
	Addresses  []Address              `json:"addresses"`
	Documents  []EmployeeDocument     `json:"documents"`
	Assets     []EmployeeAsset        `json:"assets"`
	Identity   []IdentityDocument     `json:"identity"`
	Relations  map[string]interface{} `json:"relations"`
}

type PrimaryDetailsRequest struct {
	FirstName     string `json:"first_name"     validate:"required,min=1,max=100"`
	MiddleName    string `json:"middle_name"    validate:"max=100"`
	LastName      string `json:"last_name"      validate:"required,min=1,max=100"`
	DisplayName   string `json:"display_name"   validate:"max=100"`
	Gender        string `json:"gender"         validate:"required,oneof=MALE FEMALE NON_BINARY PREFER_NOT_TO_SAY"`
	DOB           string `json:"dob"            validate:"required"`
	MaritalStatus string `json:"marital_status" validate:"oneof=SINGLE MARRIED DIVORCED WIDOWED"`
	BloodGroup    string `json:"blood_group"    validate:"max=5"`
	Nationality   string `json:"nationality"    validate:"max=100"`
}

type ContactRequest struct {
	PersonalEmail string `json:"personal_email" validate:"omitempty,email"`
	MobileNo      string `json:"mobile_no"      validate:"omitempty,min=10,max=15"`
	WorkNo        string `json:"work_no"        validate:"omitempty,max=15"`
	ResidenceNo   string `json:"residence_no"   validate:"omitempty,max=15"`
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
	EmployeeID       int     `json:"employee_id"        validate:"required"`
	Degree           string  `json:"degree"             validate:"required,min=1,max=100"`
	Branch           string  `json:"branch"             validate:"max=100"`
	University       string  `json:"university"         validate:"required,min=1,max=200"`
	CGPAOrPct        float64 `json:"cgpa_or_pct"        validate:"min=0,max=100"`
	YearOfJoining    int     `json:"year_of_joining"    validate:"required,min=1950,max=2100"`
	YearOfCompletion int     `json:"year_of_completion" validate:"min=1950,max=2100"`
}

type ExperienceRequest struct {
	EmployeeID     int     `json:"employee_id"      validate:"required"`
	CompanyName    string  `json:"company_name"     validate:"required,min=1,max=200"`
	Designation    string  `json:"designation"      validate:"required,min=1,max=100"`
	EmploymentType string  `json:"employment_type"  validate:"required"`
	StartDate      string  `json:"start_date"       validate:"required"`
	EndDate        *string `json:"end_date"`
	IsCurrent      bool    `json:"is_current"`
	Industry       string  `json:"industry"         validate:"max=100"`
	Description    string  `json:"description"      validate:"max=1000"`
}

type AssignAssetRequest struct {
	EmployeeID    int    `json:"employee_id"    validate:"required"`
	AssetType     string `json:"asset_type"     validate:"required"`
	AssetName     string `json:"asset_name"     validate:"required,min=1,max=200"`
	AssetCategory string `json:"asset_category" validate:"required"`
	SerialNo      string `json:"serial_no"      validate:"max=100"`
	Condition     string `json:"condition"      validate:"required,oneof=GOOD DAMAGED UNDER_REPAIR"`
	AssignedBy    int    `json:"assigned_by"    validate:"required"`
	Notes         string `json:"notes"          validate:"max=500"`
}
