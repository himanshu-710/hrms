

package handler

import "hrms/internal/onboarding/service"

type OnboardingHandler struct {
	Service *service.OnboardingService
}

func NewOnboardingHandler(service *service.OnboardingService) *OnboardingHandler {
	return &OnboardingHandler{
		Service: service,
	}
}