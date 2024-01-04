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
	"Title":            Title,
}

type TemplateData struct {
	Eggs             []Egg
	NextLayDate      NextLayDate
	NextHatchDate    NextHatchDate
	AverageHatchTime string
	TotalSales       string
	Incubators       []*Incubator
	Geckos           []*Gecko
	Tanks            []TankContents
}

func homepage(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path == "/deleteGecko" {
		geckoId, _ := strconv.Atoi(r.FormValue("geckoId"))
		gecko, _ := GetGecko(geckoId)
		gecko.Delete()
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	var incubatingEggs []Egg
	for _, egg := range eggs {
		if !egg.HasHatched {
			egg.Colour = GetEggRAG(*egg)
			incubatingEggs = append(incubatingEggs, *egg)
		}
	}
	var tankContents []TankContents
	for _, tank := range tanks {
		tankContents = append(tankContents, tank.GetTankContents())
	}

	data := TemplateData{
		Eggs:             incubatingEggs,
		NextLayDate:      GetNextLayDateInfo(),
		NextHatchDate:    GetNextHatchDateInfo(),
		AverageHatchTime: GetAverageHatchTime(),
		TotalSales:       TotalSales(),
		Incubators:       incubators,
		Tanks:            tankContents,
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
		err := tpl.Execute(w, map[string]interface{}{
			"TodaysDate": time.Now().Format("2006-01-02"),
			"Tanks":      tanks,
		})
		if err != nil {
			log.Println(err)
		}
		return
	}

	tankId, _ := strconv.Atoi(r.FormValue("tankId"))
	dateOfBirth, _ := time.Parse("2006-01-02", r.FormValue("dob"))
	_, err := AddGecko(r.FormValue("description"), tankId, r.FormValue("gender"), dateOfBirth.Format("02/01/2006"))
	if err != nil {
		fmt.Println(err)
	}

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func editGecko(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.New("new_gecko.html").Funcs(funcMap).ParseFiles("assets/new_gecko.html"))

	geckoId, _ := strconv.Atoi(r.FormValue("geckoId"))
	gecko, _ := GetGecko(geckoId)

	if r.Method != http.MethodPost {
		err := tpl.Execute(w, map[string]interface{}{
			"TodaysDate":    time.Now().Format("2006-01-02"),
			"Tanks":         tanks,
			"EditGecko":     gecko,
			"EditGeckoDate": gecko.DateOfBirth.Format("2006-01-02"),
		})
		if err != nil {
			log.Println(err)
		}
		return
	}

	gecko.TankID, _ = strconv.Atoi(r.FormValue("tankId"))
	gecko.Description = r.FormValue("description")
	gecko.Gender = r.FormValue("gender")
	gecko.DateOfBirth, _ = time.Parse("2006-01-02", r.FormValue("dob"))
	gecko.FormattedDateOfBirth = gecko.DateOfBirth.Format("02/01/2006")
	gecko.Update()

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func newEgg(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.New("new_egg.html").Funcs(funcMap).ParseFiles("assets/new_egg.html"))
	if r.Method != http.MethodPost {
		var availableIncubators map[int]Incubator = make(map[int]Incubator)
		for _, incubator := range incubators {
			availableIncubators[incubator.ID] = *incubator
		}
		var femaleGeckos []*Gecko
		for _, gecko := range geckos {
			if gecko.Gender == "female" && !gecko.Deleted {
				femaleGeckos = append(femaleGeckos, gecko)
			}
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
			"AvailableGeckos":     femaleGeckos,
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

	if incubatorId == 0 || row == 0 || column == 0 {
		http.Redirect(w, r, "/newEgg", http.StatusSeeOther)
		return
	}
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

func newTank(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("assets/new_tank.html"))
	if r.Method != http.MethodPost {
		err := tpl.Execute(w, map[string]interface{}{})
		if err != nil {
			log.Println(err)
		}
		return
	}

	AddTank(r.FormValue("name"))

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
