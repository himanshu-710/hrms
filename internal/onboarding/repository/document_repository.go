package repository

import (
	"context"
	"hrms/internal/onboarding/model"
)

func (r *OnboardingRepository) SaveDocument(doc model.EmployeeDocument) error {
	query := `
	INSERT INTO employee_documents
	(employee_id, doc_category, file_name, s3_url, file_size_kb, mime_type)
	VALUES($1,$2,$3,$4,$5,$6)
	`
	_, err := r.DB.Exec(context.Background(), query,
		doc.EmployeeID, doc.DocCategory, doc.FileName,
		doc.S3URL, doc.FileSizeKB, doc.MimeType,
	)
	return err
}

func (r *OnboardingRepository) GetDocuments(employeeID int) ([]model.EmployeeDocument, error) {
	query := `
	SELECT id, employee_id, doc_category, file_name, s3_url, verification_status, uploaded_at
	FROM employee_documents WHERE employee_id=$1
	ORDER BY uploaded_at DESC
	`
	rows, err := r.DB.Query(context.Background(), query, employeeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.EmployeeDocument
	for rows.Next() {
		var doc model.EmployeeDocument
		err := rows.Scan(
			&doc.ID, &doc.EmployeeID, &doc.DocCategory,
			&doc.FileName, &doc.S3URL,
			&doc.VerificationStatus, &doc.UploadedAt,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, doc)
	}
	return list, nil
}

func (r *OnboardingRepository) DeleteDocument(id int) error {
	_, err := r.DB.Exec(context.Background(),
		`DELETE FROM employee_documents WHERE id=$1`, id)
	return err
}

func (r *OnboardingRepository) VerifyDocument(id int, status string, note string) error {
	query := `
	UPDATE employee_documents
	SET verification_status=$1, rejection_note=$2
	WHERE id=$3
	`
	_, err := r.DB.Exec(context.Background(), query, status, note, id)
	return err
}

func (r *OnboardingRepository) GetDocumentOwner(id int) (int, error) {
	var employeeID int
	err := r.DB.QueryRow(context.Background(),
		`SELECT employee_id FROM employee_documents WHERE id=$1`, id,
	).Scan(&employeeID)
	return employeeID, err
}

func (r *OnboardingRepository) GetPendingDocuments() ([]model.PendingDocumentDTO, error) {
	query := `
	SELECT d.id, d.doc_category, d.file_name, d.s3_url, d.verification_status,
		   e.first_name || ' ' || e.last_name as employee_name
	FROM employee_documents d
	JOIN employees e ON d.employee_id = e.id
	WHERE d.verification_status = 'PENDING'
	ORDER BY d.uploaded_at DESC
	`
	rows, err := r.DB.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.PendingDocumentDTO
	for rows.Next() {
		var doc model.PendingDocumentDTO
		err := rows.Scan(
			&doc.ID, &doc.DocCategory, &doc.FileName, &doc.S3URL,
			&doc.VerificationStatus, &doc.EmployeeName,
		)
		if err != nil {
			return nil, err
		}
		list = append(list, doc)
	}
	return list, nil
}
