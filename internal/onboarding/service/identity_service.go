package service

import (
	"os"
	"hrms/internal/onboarding/model"
	"hrms/pkg/utils"
)

func (s *OnboardingService) SaveIdentity(req model.IdentityRequest) error {

	key := []byte(os.Getenv("AES_KEY"))

	// Encrypt only sensitive docs
	if req.DocType == "AADHAAR" || req.DocType == "PAN" {
		encrypted, err := utils.Encrypt(req.DocNumber, key)
		if err != nil {
			return err
		}
		req.DocNumber = encrypted
	}

	return s.Repo.SaveIdentity(req)
}

func (s *OnboardingService) GetIdentity(employeeID int) ([]model.IdentityDocument, error) {

	key := []byte(os.Getenv("AES_KEY"))

	list, err := s.Repo.GetIdentity(employeeID)
	if err != nil {
		return nil, err
	}

	for i := range list {

		doc := &list[i]

		if doc.DocType == "AADHAAR" || doc.DocType == "PAN" {

			decrypted, err := utils.Decrypt(doc.DocNumber, key)
			if err != nil {
				return nil, err
			}

			// Apply masking
			if doc.DocType == "AADHAAR" {
				doc.DocNumber = utils.MaskAadhaar(decrypted)
			} else if doc.DocType == "PAN" {
				doc.DocNumber = utils.MaskPAN(decrypted)
			}
		}
	}

	return list, nil
}