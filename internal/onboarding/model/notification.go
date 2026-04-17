package model

import "time"

const EventOnboardingReminder = "ONBOARDING_REMINDER"

type Notification struct {
	ID               int                    `json:"id"`
	EmployeeID       int                    `json:"employee_id"`
	NotificationType string                 `json:"notification_type"`
	Title            string                 `json:"title"`
	Message          string                 `json:"message"`
	ReminderDay      int                    `json:"reminder_day"`
	Metadata         map[string]interface{} `json:"metadata"`
	IsRead           bool                   `json:"is_read"`
	CreatedAt        time.Time              `json:"created_at"`
	ReadAt           *time.Time             `json:"read_at,omitempty"`
}

type ReminderCandidate struct {
	Employee    Employee
	ReminderDay int
}
