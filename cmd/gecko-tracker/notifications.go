package main

import (
	"log"
	"time"

	"github.com/gregdel/pushover"
)

type Notifications struct {
	Configured bool
	Pushover   Pushover
}

func NotificationTimer() {
	for {
		// run at 9am
		if time.Now().Format("15:04") == "09:00" {
			LayDateNotifications()
			HatchDateNotifications()
		}
		time.Sleep(1 * time.Minute)
	}
}

type Pushover struct {
	Device    string
	Sound     string
	UserToken string
	APIToken  string
}

func PushoverNotification(n Pushover, Message string) {
	app := pushover.New(n.UserToken)
	recipient := pushover.NewRecipient(n.APIToken)

	message := &pushover.Message{
		Message:    Message,
		DeviceName: n.Device,
		// Sound:      n.Sound,
	}

	_, err := app.SendMessage(message, recipient)
	if err != nil {
		log.Println(err)
	}
}

func SendNotification(Message string) {
	log.Println("Sending notification: " + Message)

	if notifications.Pushover.Device != "" && notifications.Pushover.UserToken != "" && notifications.Pushover.APIToken != "" {
		PushoverNotification(notifications.Pushover, Message)
	}
}

func LayDateNotifications() {
	layETA := GetNextLayDateInfo().Date

	// current date is lay date
	if time.Now().Format("02-01-2006") == layETA.Format("02-01-2006") {
		SendNotification("Gecko due to lay eggs today")
		return
	}

	// current date is after lay date
	if time.Now().After(layETA) {
		SendNotification("Gecko should have laid an egg by now")
		return
	}

	// current date is 1 day before lay date
	tomorrow := time.Now().Add(24 * time.Hour)
	if tomorrow.Format("02-01-2006") == layETA.Format("02-01-2006") {
		SendNotification("Gecko due to lay an egg tomorrow")
		return
	}
}

func HatchDateNotifications() {
	layETA := GetNextHatchDateInfo().Date

	// current date is lay date
	if time.Now().Format("02-01-2006") == layETA.Format("02-01-2006") {
		SendNotification("Egg due to hatch today")
		return
	}

	// current date is after lay date
	if time.Now().After(layETA) {
		SendNotification("Egg should have hatched by now")
		return
	}

	// current date is 1 day before lay date
	tomorrow := time.Now().Add(24 * time.Hour)
	if tomorrow.Format("02-01-2006") == layETA.Format("02-01-2006") {
		SendNotification("Egg due to hatch tomorrow")
		return
	}
}
