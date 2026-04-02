package model

import "time"

type IdentityDocument struct {
	ID int `json:"id"`

	EmployeeID int `json:"employee_id"`

	DocType string `json:"doc_type"` 

	DocNumber string `json:"doc_number"` 

	NameOnDoc string `json:"name_on_doc"`

	IssueDate *time.Time `json:"issue_date"`
	ExpiryDate *time.Time `json:"expiry_date"`

	ExtraInfo map[string]interface{} `json:"extra_info"` 

	CreatedAt time.Time `json:"created_at"`
}

type EmployeeDocument struct {
	ID int `json:"id"`

	EmployeeID int `json:"employee_id"`

	DocCategory string `json:"doc_category"` 

	FileName string `json:"file_name"`

	S3URL string `json:"s3_url"`
	 PresignedURL       string     `json:"presigned_url"` 

	FileSizeKB int `json:"file_size_kb"`
	MimeType   string `json:"mime_type"`

	VerificationStatus string `json:"verification_status"` 

	VerifiedBy *int `json:"verified_by"`
	VerifiedAt *time.Time `json:"verified_at"`

	RejectionNote *string `json:"rejection_note"`

	UploadedAt time.Time `json:"uploaded_at"`
}
type UploadDocumentRequest struct {
	EmployeeID  int    `form:"employee_id"`
	DocCategory string `form:"doc_category"`
}