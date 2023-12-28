package main

import (
	"net/http"
)

type Grid struct {
	Rows    int
	Columns int
	Row     int
	Column  int
}

var geckos []Gecko
var eggs []Egg
var sales []Sale
var availableSources = []string{"Preloved", "Facebook"}
var incubatorSize Grid = Grid{
	Rows:    3,
	Columns: 2,
}

func main() {
	http.HandleFunc("/styles.css", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "assets/styles.css") })
	http.HandleFunc("/", homepage)
	http.HandleFunc("/newEgg", newEgg)
	http.HandleFunc("/newGecko", newGecko)
	http.HandleFunc("/newSale", newSale)
	http.ListenAndServe(":8080", nil)
}
