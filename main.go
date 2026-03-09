package main

import (
	"nilchanpub/notificationSender"
	"nilchanpub/notificationSender/notification"
	"nilchanpub/notificationSender/notification/Link"
	"strconv"
)

func main() {
	/*apartment := structs.NewApartment("HKJHKJHS", 25.6, 10, 500)
	fmt.Println(apartment)
	apartment.ChangePrice(12931)
	fmt.Println(apartment)*/
	//Вид канала свяи
	link1 := Link.Email{}
	//Переменная для количесвта уведомлений
	n := 100
	//Иницализация массива уведомлений
	var notificationArray []Notification.Notification
	//Заполняю массив уведомлений
	for i := 0; i < n; i++ {
		//Создаю уведомление
		notification := Notification.Notification{
			Text: strconv.Itoa(i),
		}
		notificationArray = append(notificationArray, notification)
	}
	notificationSender := NotificationSender.NewNotificationSender(link1)
	notificationSender.GetAndSend(notificationArray)
}
