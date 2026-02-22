package main

import (
	"fmt"
)

type Notifier interface {
	Notify(message string) error
}

type EmailNotifier struct {
	To string
}

type SlackNotifier struct {
	Channel string
}

func (e *EmailNotifier) Notify(message string) error {
	fmt.Printf("Sending email to %s: %s\n", e.To, message)
	return nil
}

func (s *SlackNotifier) Notify(message string) error {
	fmt.Printf("Sending Slack message to channel %s: %s\n", s.Channel, message)
	return nil
}

func SendAll(notifiers []Notifier, message string) []error {
	var errors []error
	for _, notifier := range notifiers {
		if err := notifier.Notify(message); err != nil {
			errors = append(errors, err)
		}
	}
	return errors
}	