package main

import (
	"net/http"
)

var geckos []*Gecko
var incubators []*Incubator
var tanks []*Tank
var eggs []*Egg
var sales []*Sale

func main() {
	LoadConfiguration()
	LoadFromDB()

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
	http.ListenAndServe(":8080", nil)
}
