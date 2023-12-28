package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"time"
)

var funcMap = template.FuncMap{
	"mod":               mod,
	"add":               add,
	"sortEggsByGeckoID": sortEggsByGeckoID,
	"N":                 N,
}

func homepage(w http.ResponseWriter, r *http.Request) {
	var incubatingEggs []Egg
	for _, egg := range eggs {
		if egg.SlotID != 0 {
			incubatingEggs = append(incubatingEggs, egg)
		}
	}
	data := TemplateData{
		Eggs:          incubatingEggs,
		NextLayDate:   GetNextLayDateInfo(),
		NextHatchDate: GetNextHatchDateInfo(),
		TotalSales:    TotalSales(),
	}

	tpl := template.Must(template.New("home.html").Funcs(funcMap).ParseFiles("assets/home.html"))
	tpl.Execute(w, data)
}

func newGecko(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("assets/new_gecko.html"))
	if r.Method != http.MethodPost {
		tpl.Execute(w, nil)
		return
	}

	geckoId, _ := strconv.Atoi(r.FormValue("geckoId"))
	_, err := AddGecko(geckoId, r.FormValue("description"))
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func newEgg(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.New("new_egg.html").Funcs(funcMap).ParseFiles("assets/new_egg.html"))
	if r.Method != http.MethodPost {
		var availableGeckos []int
		for _, gecko := range geckos {
			availableGeckos = append(availableGeckos, gecko.ID)
		}
		err := tpl.Execute(w, map[string]interface{}{
			"AvailableGeckos": availableGeckos,
			"TodaysDate":      time.Now().Format("2006-01-02"),
			"IncubatorSize":   incubatorSize,
		})
		if err != nil {
			log.Println(err)
		}
		return
	}

	slotId, _ := strconv.Atoi(r.FormValue("slotId"))
	geckoId, _ := strconv.Atoi(r.FormValue("gecko"))
	eggCount, _ := strconv.Atoi(r.FormValue("eggCount"))
	date, _ := time.Parse("2006-01-02", r.FormValue("date"))
	AddEgg(slotId, geckoId, eggCount, date.Format("02/01/2006"), "")

	// tpl.Execute(w, struct{ Success bool }{true})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func newSale(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("assets/new_sale.html"))
	if r.Method != http.MethodPost {
		tpl.Execute(w, map[string]interface{}{
			"AvailableSources": availableSources,
			"TodaysDate":       time.Now().Format("2006-01-02"),
		})
		return
	}

	male, _ := strconv.Atoi(r.FormValue("male"))
	female, _ := strconv.Atoi(r.FormValue("female"))
	baby, _ := strconv.Atoi(r.FormValue("baby"))
	price, _ := strconv.Atoi(r.FormValue("price"))
	date, _ := time.Parse("2006-01-02", r.FormValue("date"))
	AddSale(r.FormValue("buyer"), r.FormValue("source"), male, female, baby, price, date.Format("02/01/2006"))

	// tpl.Execute(w, struct{ Success bool }{true})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
