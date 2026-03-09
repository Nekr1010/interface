package Link

import (
	"fmt"
	Notification "nilchanpub/notificationSender/notification"
)

type SMS struct{}

func (e SMS) Send(notifications []Notification.Notification) {
	notificationsCount := len(notifications)
	fmt.Println("Оправил", notificationsCount, "уведомлений через SMS")
}
func (e SMS) Get(notifications []Notification.Notification) {
	notificationsCount := len(notifications)
	fmt.Println("Получил", notificationsCount, "уведомлений через SMS")
}
