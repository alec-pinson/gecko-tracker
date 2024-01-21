package main

import (
	"log"
	"net/http"
)

var geckos []*Gecko
var incubators []*Incubator
var tanks []*Tank
var eggs []*Egg
var sales []*Sale
var notifications Notifications

func main() {
	LoadConfiguration()
	LoadFromDB()

	log.Println(notifications.Pushover.APIToken)
	log.Println(notifications)

	go NotificationTimer()

	http.HandleFunc("/styles.css", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "assets/styles.css") })
	http.HandleFunc("/images/favicon.ico", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "assets/images/favicon.ico") })
	http.HandleFunc("/images/edit.png", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "assets/images/edit.png") })
	http.HandleFunc("/", homepage)
	http.HandleFunc("/newIncubator", newIncubator)
	http.HandleFunc("/newTank", newTank)
	http.HandleFunc("/newEgg", newEgg)
	http.HandleFunc("/newGecko", newGecko)
	http.HandleFunc("/editGecko", editGecko)
	http.HandleFunc("/newSale", newSale)
	http.HandleFunc("/hasHatched", hasHatched)
	http.HandleFunc("/notifications", notificationSetup)
	http.ListenAndServe(":8080", nil)
}
