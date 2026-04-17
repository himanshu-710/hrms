package repository

import (
	"context"
	"encoding/json"

	"hrms/internal/onboarding/model"
)

func (r *OnboardingRepository) GetEmployeesForOnboardingReminders(ctx context.Context) ([]model.ReminderCandidate, error) {
	query := `
	SELECT id, first_name, last_name, work_email, date_of_joining, (CURRENT_DATE - date_of_joining) AS reminder_day
	FROM employees
	WHERE is_active = true
	  AND date_of_joining IS NOT NULL
	  AND (CURRENT_DATE - date_of_joining) IN (3, 7, 14)
	`

	rows, err := r.DB.Query(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var list []model.ReminderCandidate
	for rows.Next() {
		var candidate model.ReminderCandidate
		if err := rows.Scan(
			&candidate.Employee.ID,
			&candidate.Employee.FirstName,
			&candidate.Employee.LastName,
			&candidate.Employee.WorkEmail,
			&candidate.Employee.DateOfJoining,
			&candidate.ReminderDay,
		); err != nil {
			return nil, err
		}
		list = append(list, candidate)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return list, nil
}

func (r *OnboardingRepository) CreateNotification(ctx context.Context, notification model.Notification) error {
	var metadata []byte
	var err error

	if notification.Metadata != nil {
		metadata, err = json.Marshal(notification.Metadata)
		if err != nil {
			return err
		}
	}

	query := `
	INSERT INTO notifications (
		employee_id,
		notification_type,
		title,
		message,
		reminder_day,
		metadata
	) VALUES ($1, $2, $3, $4, $5, $6)
	ON CONFLICT (employee_id, notification_type, reminder_day) DO NOTHING
	`

	_, err = r.DB.Exec(
		ctx,
		query,
		notification.EmployeeID,
		notification.NotificationType,
		notification.Title,
		notification.Message,
		notification.ReminderDay,
		metadata,
	)
	return err
}

func (r *OnboardingRepository) GetNotifications(ctx context.Context, employeeID int) ([]model.Notification, error) {
	query := `
	SELECT id, employee_id, notification_type, title, message, reminder_day, metadata, is_read, created_at, read_at
	FROM notifications
	WHERE employee_id = $1
	ORDER BY created_at DESC
	`

	rows, err := r.DB.Query(ctx, query, employeeID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notifications []model.Notification
	for rows.Next() {
		var item model.Notification
		var metadata []byte

		if err := rows.Scan(
			&item.ID,
			&item.EmployeeID,
			&item.NotificationType,
			&item.Title,
			&item.Message,
			&item.ReminderDay,
			&metadata,
			&item.IsRead,
			&item.CreatedAt,
			&item.ReadAt,
		); err != nil {
			return nil, err
		}

		if len(metadata) > 0 {
			if err := json.Unmarshal(metadata, &item.Metadata); err != nil {
				return nil, err
			}
		}

		notifications = append(notifications, item)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notifications, nil
}

func (r *OnboardingRepository) MarkNotificationRead(ctx context.Context, employeeID int, notificationID int) error {
	_, err := r.DB.Exec(
		ctx,
		`UPDATE notifications
		 SET is_read = true, read_at = CURRENT_TIMESTAMP
		 WHERE id = $1 AND employee_id = $2`,
		notificationID,
		employeeID,
	)
	return err
}
