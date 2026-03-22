package model
type EmployeeAsset struct {
	ID int `json:"id"`

	EmployeeID int `json:"employee_id"`

	AssetCode string `json:"asset_code"`

	AssetType string `json:"asset_type"`

	AssetName string `json:"asset_name"`

	AssignedOn string `json:"assigned_on"`

	AcknowledgementStatus string `json:"acknowledgement_status"`

	AcknowledgedAt string `json:"acknowledged_at"`
}