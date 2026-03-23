package service

import "hrms/internal/onboarding/model"

func (s *OnboardingService) SaveAddresses(req model.AddressesRequest) error {

	if req.CopyFromCurrent {
		req.Permanent = req.Current
	}

	return s.Repo.SaveAddresses(req)
}