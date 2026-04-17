package scheduler

import (
	"context"
	"log"
	"time"

	"github.com/robfig/cron/v3"
)

const OnboardingReminderCronSpec = "0 9 * * *"

type onboardingReminderService interface {
	SendOnboardingReminders(ctx context.Context) error
}

func StartOnboardingReminderCron(ctx context.Context, svc onboardingReminderService) *cron.Cron {
	c := cron.New(cron.WithLocation(time.Local))

	_, err := c.AddFunc(OnboardingReminderCronSpec, func() {
		if err := svc.SendOnboardingReminders(ctx); err != nil {
			log.Printf("onboarding reminder job failed: %v", err)
		}
	})
	if err != nil {
		log.Printf("failed to register onboarding reminder cron: %v", err)
		return c
	}

	c.Start()
	return c
}
