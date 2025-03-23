package models

import "time"

type Notification struct {
	NotificationId int       `json:"notificationId"`
	Title          string    `json:"title"`
	Status         string    `json:"status"`
	Content        string    `json:"content"`
	Priority       string    `json:"priority"`
	UserId         int       `json:"userId"`
	Timestamp      time.Time `json:"timestamp"`
	Reminder       time.Time `json:"reminder"`
	IsReminder     bool      `json:"isReminder"`
}
