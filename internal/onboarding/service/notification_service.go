package service

import (
	"context"
	"fmt"
	"strings"

	"hrms/internal/onboarding/model"
	"hrms/internal/onboarding/repository"
)

type NotificationDispatcher interface {
	Dispatch(ctx context.Context, eventType string, employee model.Employee, notification model.Notification) error
}

type onboardingNotificationDispatcher struct {
	repo *repository.OnboardingRepository
}

func NewOnboardingNotificationDispatcher(repo *repository.OnboardingRepository) NotificationDispatcher {
	return &onboardingNotificationDispatcher{repo: repo}
}

func (d *onboardingNotificationDispatcher) Dispatch(ctx context.Context, eventType string, employee model.Employee, notification model.Notification) error {
	notification.EmployeeID = employee.ID
	notification.NotificationType = eventType
	return d.repo.CreateNotification(ctx, notification)
}

func (s *OnboardingService) SendOnboardingReminders(ctx context.Context) error {
	candidates, err := s.Repo.GetEmployeesForOnboardingReminders(ctx)
	if err != nil {
		return err
	}

	for _, candidate := range candidates {
		completion, err := s.ComputeCompletion(candidate.Employee.ID)
		if err != nil {
			return err
		}

		if completion.Percentage >= 100 {
			continue
		}

		incomplete := incompleteSections(completion.Sections)
		notification := model.Notification{
			Title:       "Complete your onboarding profile",
			Message:     fmt.Sprintf("You still have pending onboarding sections: %s.", strings.Join(incomplete, ", ")),
			ReminderDay: candidate.ReminderDay,
			Metadata: map[string]interface{}{
				"completion_percentage": completion.Percentage,
				"incomplete_sections":   incomplete,
			},
		}

		if err := s.Dispatcher.Dispatch(ctx, model.EventOnboardingReminder, candidate.Employee, notification); err != nil {
			return err
		}
	}

	return nil
}

func (s *OnboardingService) GetNotifications(ctx context.Context, employeeID int) ([]model.Notification, error) {
	return s.Repo.GetNotifications(ctx, employeeID)
}

func (s *OnboardingService) MarkNotificationRead(ctx context.Context, employeeID int, notificationID int) error {
	return s.Repo.MarkNotificationRead(ctx, employeeID, notificationID)
}

func incompleteSections(sections map[string]bool) []string {
	list := make([]string, 0, len(sections))
	for section, complete := range sections {
		if !complete {
			list = append(list, section)
		}
	}
	return list
}
