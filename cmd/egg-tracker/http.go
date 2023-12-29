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
	"sortEggsIntoGrid": sortEggsIntoGrid,
	"toSlotID":         toSlotID,
	"N":                N,
}

type TemplateData struct {
	Eggs          []Egg
	NextLayDate   string
	NextHatchDate string
	TotalSales    string
	Incubators    []*Incubator
}

func homepage(w http.ResponseWriter, r *http.Request) {
	var incubatingEggs []Egg
	for _, egg := range eggs {
		if !egg.HasHatched {
			incubatingEggs = append(incubatingEggs, *egg)
		}
	}
	data := TemplateData{
		Eggs:          incubatingEggs,
		NextLayDate:   GetNextLayDateInfo(),
		NextHatchDate: GetNextHatchDateInfo(),
		TotalSales:    TotalSales(),
		Incubators:    incubators,
	}

	tpl := template.Must(template.New("home.html").Funcs(funcMap).ParseFiles("assets/home.html"))
	err := tpl.Execute(w, data)
	if err != nil {
		log.Println(err)
	}
}

func newGecko(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("assets/new_gecko.html"))
	if r.Method != http.MethodPost {
		err := tpl.Execute(w, nil)
		if err != nil {
			log.Println(err)
		}
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
		var availableIncubators map[int]Incubator = make(map[int]Incubator)
		for _, incubator := range incubators {
			availableIncubators[incubator.ID] = *incubator
		}
		geckoId, _ := strconv.Atoi(r.FormValue("gecko"))
		if geckoId == 0 {
			geckoId = 1
		}
		incubatorId, _ := strconv.Atoi(r.FormValue("incubator"))
		if incubatorId == 0 {
			incubatorId = 1
		}
		err := tpl.Execute(w, map[string]interface{}{
			"AvailableGeckos":     availableGeckos,
			"AvailableIncubators": availableIncubators,
			"TodaysDate":          time.Now().Format("2006-01-02"),
			"SelectedGecko":       geckoId,
			"SelectedIncubator":   incubatorId,
			"Row":                 r.FormValue("row"),
			"Column":              r.FormValue("column"),
		})
		if err != nil {
			log.Println(err)
		}
		return
	}

	incubatorId, _ := strconv.Atoi(r.FormValue("incubator"))
	row, _ := strconv.Atoi(r.FormValue("row"))
	column, _ := strconv.Atoi(r.FormValue("column"))
	geckoId, _ := strconv.Atoi(r.FormValue("gecko"))
	eggCount, _ := strconv.Atoi(r.FormValue("eggCount"))
	date, _ := time.Parse("2006-01-02", r.FormValue("date"))
	AddEgg(incubatorId, row, column, geckoId, eggCount, date.Format("02/01/2006"))

	// tpl.Execute(w, struct{ Success bool }{true})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func newSale(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("assets/new_sale.html"))
	if r.Method != http.MethodPost {
		err := tpl.Execute(w, map[string]interface{}{
			"AvailableSources": config.Sources,
			"TodaysDate":       time.Now().Format("2006-01-02"),
		})
		if err != nil {
			log.Println(err)
		}
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

func newIncubator(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("assets/new_incubator.html"))
	if r.Method != http.MethodPost {
		err := tpl.Execute(w, map[string]interface{}{})
		if err != nil {
			log.Println(err)
		}
		return
	}

	rows, _ := strconv.Atoi(r.FormValue("rows"))
	columns, _ := strconv.Atoi(r.FormValue("columns"))
	AddIncubator(rows, columns)

	// tpl.Execute(w, struct{ Success bool }{true})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func hasHatched(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("assets/has_hatched.html"))
	if r.Method != http.MethodPost {
		err := tpl.Execute(w, map[string]interface{}{})
		if err != nil {
			log.Println(err)
		}
		return
	}

	eggId := r.FormValue("eggId")
	egg := GetEgg(eggId)
	egg.Hatched()

	// tpl.Execute(w, struct{ Success bool }{true})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
