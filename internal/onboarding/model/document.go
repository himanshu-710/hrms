package model

import "time"

type IdentityDocument struct {
	ID int `json:"id"`

	EmployeeID int `json:"employee_id"`

	DocType string `json:"doc_type"`

	EncryptedDocNumber string `json:"encrypted_doc_number"`

	CreatedAt time.Time `json:"created_at"`
}

type EmployeeDocument struct {
	ID int `json:"id"`

	EmployeeID int `json:"employee_id"`

	DocCategory string `json:"doc_category"`

	FileName string `json:"file_name"`

	S3URL string `json:"s3_url"`

	VerificationStatus string `json:"verification_status"`

	VerifiedBy *int `json:"verified_by"`

	VerifiedAt *time.Time `json:"verified_at"`

	RejectionNote *string `json:"rejection_note"`

	UploadedAt time.Time `json:"uploaded_at"`
}
