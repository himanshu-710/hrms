package repository

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"hrms/internal/onboarding/model"
	"hrms/pkg/utils"
)

func (r *OnboardingRepository) CreateEmployee(firstName, lastName, email, department string) error {
	query := `
	INSERT INTO employees(first_name, last_name, personal_email, employee_code, work_email)
	VALUES($1,$2,$3,$4,$5)
	`
	code := fmt.Sprintf("EMP-%d", time.Now().UnixMilli())
	workEmail := fmt.Sprintf("%s.%s@company.com", strings.ToLower(firstName), strings.ToLower(lastName))

	_, err := r.DB.Exec(context.Background(), query, firstName, lastName, email, code, workEmail)
	return err
}

func (r *OnboardingRepository) GetEmployee(id int) (*model.Employee, error) {

	query := `
	SELECT id, first_name, last_name, personal_email, mobile_no, relations
	FROM employees
	WHERE id=$1
	`

	var emp model.Employee
	var relations []byte

	err := r.DB.QueryRow(context.Background(), query, id).Scan(
		&emp.ID,
		&emp.FirstName,
		&emp.LastName,
		&emp.PersonalEmail,
		&emp.MobileNo,
		&relations,
	)
	if err != nil {
		return nil, err
	}

	if relations != nil {
		if err := json.Unmarshal(relations, &emp.Relations); err != nil {
			return nil, err
		}
	}

	return &emp, nil
}

func (r *OnboardingRepository) UpdatePrimaryDetails(id int, req model.PrimaryDetailsRequest) error {

	query := `
	UPDATE employees
	SET first_name=$1, middle_name=$2, last_name=$3, display_name=$4,
	    gender=$5, dob=$6, marital_status=$7, blood_group=$8, nationality=$9
	WHERE id=$10
	`

	_, err := r.DB.Exec(context.Background(), query,
		req.FirstName, req.MiddleName, req.LastName, req.DisplayName,
		req.Gender, req.DOB, req.MaritalStatus, req.BloodGroup, req.Nationality, id,
	)

	return err
}

func (r *OnboardingRepository) UpdateContactDetails(id int, req model.ContactRequest) error {

	query := `
	UPDATE employees
	SET personal_email=$1, mobile_no=$2, work_no=$3, residence_no=$4
	WHERE id=$5
	`

	_, err := r.DB.Exec(context.Background(), query,
		req.PersonalEmail, req.MobileNo, req.WorkNo, req.ResidenceNo, id,
	)

	return err
}

func (r *OnboardingRepository) UpdateRelations(id int, relations map[string]interface{}) error {

	data, err := json.Marshal(relations)
	if err != nil {
		return err
	}

	_, err = r.DB.Exec(context.Background(),
		`UPDATE employees SET relations=$1 WHERE id=$2`, data, id)

	return err
}

func (r *OnboardingRepository) GetFullProfile(id int) (*model.OnboardingProfileDTO, error) {

	emp, err := r.GetEmployee(id)
	if err != nil {
		return nil, err
	}

	education, err := r.GetEducation(id)
	if err != nil {
		return nil, err
	}

	experience, err := r.GetExperience(id)
	if err != nil {
		return nil, err
	}

	
	addrRows, err := r.DB.Query(context.Background(), `
	SELECT id,employee_id,address_type,line1,line2,city,state,pin_code,country,ownership_type
	FROM employee_addresses WHERE employee_id=$1`, id)
	if err != nil {
		return nil, err
	}
	defer addrRows.Close()

	var addresses []model.Address
	for addrRows.Next() {
		var addr model.Address
		err := addrRows.Scan(
			&addr.ID, &addr.EmployeeID, &addr.AddressType,
			&addr.Line1, &addr.Line2, &addr.City, &addr.State,
			&addr.PinCode, &addr.Country, &addr.OwnershipType,
		)
		if err != nil {
			return nil, err
		}
		addresses = append(addresses, addr)
	}

	
	docRows, err := r.DB.Query(context.Background(), `
	SELECT id,employee_id,doc_category,file_name,s3_url,file_size_kb,mime_type,
	       verification_status,verified_by,verified_at,rejection_note,uploaded_at
	FROM employee_documents WHERE employee_id=$1`, id)
	if err != nil {
		return nil, err
	}
	defer docRows.Close()

	var documents []model.EmployeeDocument
	for docRows.Next() {
		var doc model.EmployeeDocument
		err := docRows.Scan(
			&doc.ID, &doc.EmployeeID, &doc.DocCategory, &doc.FileName,
			&doc.S3URL, &doc.FileSizeKB, &doc.MimeType,
			&doc.VerificationStatus, &doc.VerifiedBy, &doc.VerifiedAt,
			&doc.RejectionNote, &doc.UploadedAt,
		)
		if err != nil {
			return nil, err
		}
		documents = append(documents, doc)
	}

	
	assetRows, err := r.DB.Query(context.Background(), `
	SELECT id,employee_id,asset_type,asset_name,asset_category,serial_no,
	       assigned_on,acknowledgement_status,acknowledged_at,condition,
	       assigned_by,notes,returned_on
	FROM employee_assets WHERE employee_id=$1 AND is_active=true`, id)
	if err != nil {
		return nil, err
	}
	defer assetRows.Close()

	var assets []model.EmployeeAsset
	for assetRows.Next() {
		var a model.EmployeeAsset
		err := assetRows.Scan(
			&a.ID, &a.EmployeeID, &a.AssetType, &a.AssetName,
			&a.AssetCategory, &a.SerialNo, &a.AssignedOn,
			&a.AcknowledgementStatus, &a.AcknowledgedAt,
			&a.Condition, &a.AssignedBy, &a.Notes, &a.ReturnedOn,
		)
		if err != nil {
			return nil, err
		}
		assets = append(assets, a)
	}

	
	idRows, err := r.DB.Query(context.Background(), `
	SELECT id,employee_id,doc_type,doc_number,name_on_doc,
	       issue_date,expiry_date,extra_info,created_at
	FROM employee_identity_documents WHERE employee_id=$1`, id)
	if err != nil {
		return nil, err
	}
	defer idRows.Close()

	var identity []model.IdentityDocument
	for idRows.Next() {
		var doc model.IdentityDocument
		var extra []byte

		err := idRows.Scan(
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

		if doc.DocType == "AADHAAR" || doc.DocType == "PAN" {
			decrypted, err := utils.Decrypt(doc.DocNumber, []byte(os.Getenv("AES_KEY")))
			if err != nil {
				return nil, err
			}
			if doc.DocType == "AADHAAR" {
				doc.DocNumber = utils.MaskAadhaar(decrypted)
			} else {
				doc.DocNumber = utils.MaskPAN(decrypted)
			}
		}

		identity = append(identity, doc)
	}

	relations := emp.Relations
	if relations == nil {
		relations = make(map[string]interface{})
	}

	return &model.OnboardingProfileDTO{
		Employee:   *emp,
		Education:  education,
		Experience: experience,
		Addresses:  addresses,
		Documents:  documents,
		Assets:     assets,
		Identity:   identity,
		Relations:  relations,
	}, nil
}