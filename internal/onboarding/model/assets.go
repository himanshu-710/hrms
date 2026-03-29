package model

import "time"

type EmployeeAsset struct {
	ID                     int        `json:"id"`
	EmployeeID             int        `json:"employee_id"`
	AssetType              string     `json:"asset_type"`
	AssetName              string     `json:"asset_name"`
	AssetCategory          string     `json:"asset_category"`
	SerialNo               string     `json:"serial_no"`
	AssignedOn             time.Time  `json:"assigned_on"`
	AcknowledgementStatus  string     `json:"acknowledgement_status"`
	AcknowledgedAt         *time.Time `json:"acknowledged_at"`
	Condition              string     `json:"condition"`
	AssignedBy             int        `json:"assigned_by"`
	Notes                  string     `json:"notes"`
	IsActive               bool       `json:"is_active"`
	ReturnedOn *time.Time `json:"returned_on"`
}