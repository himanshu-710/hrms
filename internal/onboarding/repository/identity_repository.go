package repository

import (
	"context"
	"encoding/json"
	"hrms/internal/onboarding/model"
)

func (r *OnboardingRepository) SaveIdentity(req model.IdentityRequest) error {

	extraJSON, err := json.Marshal(req.ExtraInfo)
	if err != nil {
		return err
	}

	query := `
	INSERT INTO employee_identity_documents
	(employee_id, doc_type, doc_number, name_on_doc, issue_date, expiry_date, extra_info)
	VALUES($1,$2,$3,$4,$5,$6,$7)
	ON CONFLICT(employee_id, doc_type)
	DO UPDATE SET
	doc_number=EXCLUDED.doc_number,
	name_on_doc=EXCLUDED.name_on_doc,
	issue_date=EXCLUDED.issue_date,
	expiry_date=EXCLUDED.expiry_date,
	extra_info=EXCLUDED.extra_info
	`

	_, err = r.DB.Exec(context.Background(), query,
		req.EmployeeID, req.DocType, req.DocNumber, req.NameOnDoc,
		req.IssueDate, req.ExpiryDate, extraJSON,
	)

	return err
}

func (r *OnboardingRepository) GetIdentity(employeeID int) ([]model.IdentityDocument, error) {

	query := `
	SELECT id, employee_id, doc_type, doc_number, name_on_doc,
	issue_date, expiry_date, extra_info, created_at
	FROM employee_identity_documents
	WHERE employee_id=$1
	`

	rows, err := r.DB.Query(context.Background(), query, employeeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.IdentityDocument

	for rows.Next() {
		var doc model.IdentityDocument
		var extra []byte

		err := rows.Scan(
			&doc.ID, &doc.EmployeeID, &doc.DocType, &doc.DocNumber,
			&doc.NameOnDoc, &doc.IssueDate, &doc.ExpiryDate,
			&extra, &doc.CreatedAt,
		)
		if err != nil {
			return nil, err
		}

		if extra != nil {
			if err := json.Unmarshal(extra, &doc.ExtraInfo); err != nil {
				return nil, err
			}
		}

		list = append(list, doc)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return list, nil
}