package repository

import "hrms/internal/onboarding/model"

func (r *OnboardingRepository) SaveAddresses(req model.AddressesRequest) error {

	query := `
	INSERT INTO employee_addresses
	(employee_id,address_type,street,city,state,country,postal_code)
	VALUES($1,$2,$3,$4,$5,$6,$7)
	ON CONFLICT(employee_id,address_type)
	DO UPDATE SET
	street=EXCLUDED.street,
	city=EXCLUDED.city,
	state=EXCLUDED.state,
	country=EXCLUDED.country,
	postal_code=EXCLUDED.postal_code
	`

	_, err := r.DB.Exec(
		query,
		req.EmployeeID,
		"CURRENT",
		req.Current.Street,
		req.Current.City,
		req.Current.State,
		req.Current.Country,
		req.Current.PostalCode,
	)

	if err != nil {
		return err
	}

	_, err = r.DB.Exec(
		query,
		req.EmployeeID,
		"PERMANENT",
		req.Permanent.Street,
		req.Permanent.City,
		req.Permanent.State,
		req.Permanent.Country,
		req.Permanent.PostalCode,
	)

	return err
}