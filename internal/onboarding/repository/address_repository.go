package repository

import (
	"context"
	"hrms/internal/onboarding/model"
)

func (r *OnboardingRepository) SaveAddresses(req model.AddressesRequest) error {

	query := `
	INSERT INTO employee_addresses
	(employee_id,address_type,line1,city,state,country,pin_code)
	VALUES($1,$2,$3,$4,$5,$6,$7)
	ON CONFLICT(employee_id,address_type)
	DO UPDATE SET
	line1=EXCLUDED.line1,
	city=EXCLUDED.city,
	state=EXCLUDED.state,
	country=EXCLUDED.country,
	pin_code=EXCLUDED.pin_code
	`

	_, err := r.DB.Exec(context.Background(),
		query,
		req.EmployeeID, "CURRENT",
		req.Current.Line1, req.Current.City,
		req.Current.State, req.Current.Country, req.Current.PinCode,
	)
	if err != nil {
		return err
	}

	_, err = r.DB.Exec(context.Background(),
		query,
		req.EmployeeID, "PERMANENT",
		req.Permanent.Line1, req.Permanent.City,
		req.Permanent.State, req.Permanent.Country, req.Permanent.PinCode,
	)

	return err
}