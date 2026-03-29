package service

import "hrms/internal/onboarding/model"

func (s *OnboardingService) ComputeCompletion(employeeID int) (*model.CompletionDTO, error) {

	sections := make(map[string]bool)

	// Profile → always true if employee exists
	sections["profile"] = true

	// Contact → check mobile/email
	emp, err := s.Repo.GetEmployee(employeeID)
	if err != nil {
		return nil, err
	}
	sections["contact"] = emp.PersonalEmail != "" || emp.MobileNo != ""

	education, err := s.Repo.HasEducation(employeeID)
if err != nil {
    return nil, err
}
experience, err := s.Repo.HasExperience(employeeID)
if err != nil {
    return nil, err
}
address, err := s.Repo.HasAddress(employeeID)
if err != nil {
    return nil, err
}
documents, err := s.Repo.HasDocuments(employeeID)
if err != nil {
    return nil, err
}
identity, err := s.Repo.HasIdentity(employeeID)
if err != nil {
    return nil, err
}
assets, err := s.Repo.HasAssets(employeeID)
if err != nil {
    return nil, err
}

	sections["education"] = education
	sections["experience"] = experience
	sections["addresses"] = address
	sections["documents"] = documents
	sections["identity"] = identity
	sections["assets"] = assets

	// Relations (from JSONB)
	sections["relations"] = emp.Relations != nil

	// Calculate percentage
	total := len(sections)
	completed := 0

	for _, v := range sections {
		if v {
			completed++
		}
	}

	percentage := (completed * 100) / total

	return &model.CompletionDTO{
		Sections:   sections,
		Percentage: percentage,
	}, nil
}