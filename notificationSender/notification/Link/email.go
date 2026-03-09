package Link

import (
	"fmt"
	Notification "nilchanpub/notificationSender/notification"
)

type Email struct{}

func (e Email) Send(notifications []Notification.Notification) {
	notificationsCount := len(notifications)
	fmt.Println("Оправил", notificationsCount, "уведомлений через email")
}
func (e Email) Get(notifications []Notification.Notification) {
	notificationsCount := len(notifications)
	fmt.Println("Получил", notificationsCount, "уведомлений через email")
}
