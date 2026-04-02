package repository

import (
	"context"
	"fmt"
	"hrms/internal/onboarding/model"
)

func (r *OnboardingRepository) GetAssets(employeeID int) ([]model.EmployeeAsset, error) {

	query := `
	SELECT id, employee_id, asset_type, asset_name, asset_category,
	       serial_no, assigned_on, acknowledgement_status, acknowledged_at,
	       condition, assigned_by, notes
	FROM employee_assets
	WHERE employee_id=$1 AND is_active=true
	`

	rows, err := r.DB.Query(context.Background(), query, employeeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.EmployeeAsset

	for rows.Next() {
		var a model.EmployeeAsset

		err := rows.Scan(
			&a.ID, &a.EmployeeID, &a.AssetType, &a.AssetName,
			&a.AssetCategory, &a.SerialNo, &a.AssignedOn,
			&a.AcknowledgementStatus, &a.AcknowledgedAt,
			&a.Condition, &a.AssignedBy, &a.Notes,
		)
		if err != nil {
			return nil, err
		}

		list = append(list, a)
	}

	return list, nil
}

func (r *OnboardingRepository) AcknowledgeAsset(id int) error {

	query := `
	UPDATE employee_assets
	SET acknowledgement_status='ACKNOWLEDGED', acknowledged_at=NOW()
	WHERE id=$1 AND acknowledgement_status='PENDING'
	`

	res, err := r.DB.Exec(context.Background(), query, id)
	if err != nil {
		return err
	}

	if res.RowsAffected() == 0 {
		return fmt.Errorf("already acknowledged or not found")
	}

	return nil
}

func (r *OnboardingRepository) AssignAsset(req model.AssignAssetRequest) error {  

    query := `
    INSERT INTO employee_assets
    (employee_id, asset_type, asset_name, asset_category,
     serial_no, assigned_on, condition, assigned_by, notes, is_active)
    VALUES($1,$2,$3,$4,$5,NOW(),$6,$7,$8,true)
    `

    _, err := r.DB.Exec(context.Background(), query,
        req.EmployeeID, req.AssetType, req.AssetName, req.AssetCategory,
        req.SerialNo, req.Condition, req.AssignedBy, req.Notes,
    )

    return err
}
func (r *OnboardingRepository) GetAssetOwner(id int) (int, error) {
	var employeeID int
	err := r.DB.QueryRow(context.Background(),
		`SELECT employee_id FROM employee_assets WHERE id=$1`, id,
	).Scan(&employeeID)
	return employeeID, err
}