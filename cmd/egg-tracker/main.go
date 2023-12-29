package main

import (
	"net/http"
)

var geckos []*Gecko
var incubators []*Incubator
var eggs []*Egg
var sales []*Sale

func main() {
	LoadConfiguration()
	LoadFromDB()

	http.HandleFunc("/styles.css", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "assets/styles.css") })
	http.HandleFunc("/images/favicon.ico", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "assets/images/favicon.ico") })
	http.HandleFunc("/", homepage)
	http.HandleFunc("/newIncubator", newIncubator)
	http.HandleFunc("/newEgg", newEgg)
	http.HandleFunc("/newGecko", newGecko)
	http.HandleFunc("/newSale", newSale)
	http.HandleFunc("/hasHatched", hasHatched)
	http.ListenAndServe(":8080", nil)
}
