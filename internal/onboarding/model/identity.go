package model
import "time"

type IdentityRequest struct {
	EmployeeID int                    `json:"employee_id"`
	DocType    string                 `json:"doc_type"`
	DocNumber  string                 `json:"doc_number"`
	NameOnDoc  string                 `json:"name_on_doc"`
	IssueDate  *time.Time             `json:"issue_date"`
	ExpiryDate *time.Time             `json:"expiry_date"`
	ExtraInfo  map[string]interface{} `json:"extra_info"`
}