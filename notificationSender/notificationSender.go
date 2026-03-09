package NotificationSender

import "nilchanpub/notificationSender/notification"

// Link канал связи отправки уведомлений
type Link interface {
	Get(notifications []Notification.Notification)
	Send(notifications []Notification.Notification)
}

// Модуль отправки уведомлений
// Зависит от канала связи
type NotificationSender struct {
	Link Link
}

func NewNotificationSender(Link Link) *NotificationSender {
	return &NotificationSender{
		Link: Link}
}

// Метод GetAndSend() принмает и отправляет массив уведомлений
func (nS NotificationSender) GetAndSend(notifications []Notification.Notification) {
	nS.Link.Get(notifications)
	nS.Link.Send(notifications)
}
