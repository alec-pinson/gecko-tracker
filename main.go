package main

import (
	"net/http"
)

type TemplateData struct {
	Eggs          []Egg
	NextLayDate   string
	NextHatchDate string
	TotalSales    string
}

var geckos []Gecko
var eggs []Egg
var sales []Sale
var availableSources = []string{"Preloved", "Facebook"}

func main() {
	http.HandleFunc("/styles.css", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "assets/styles.css") })
	http.HandleFunc("/", homepage)
	http.HandleFunc("/newEgg", newEgg)
	http.HandleFunc("/newGecko", newGecko)
	http.HandleFunc("/newSale", newSale)
	http.ListenAndServe(":8080", nil)

}
