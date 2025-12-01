package models

type NotificationType int

const (
	LowStock NotificationType = iota + 1
	Update
	Advertisement
)

type Notification struct {
	Type  NotificationType `json:"notification_type"`
	Title string           `json:"title"`
	Body  string           `json:"body"`
}

func NewNotification(t NotificationType, title, body string) Notification {
	return Notification{Type: t, Title: title, Body: body}
}
