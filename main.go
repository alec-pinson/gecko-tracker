package main

import (
	"fmt"
	"html/template"
	"net/http"
	"sort"
	"strconv"
	"time"
)

type EggSlot struct {
	Gecko    int
	EggCount int
	LayDate  string
	HatchEta string
}
type TemplateData struct {
	Eggs          []Egg
	NextLayDate   string
	NextHatchDate string
	TotalSales    string
}

var geckos []Gecko
var eggs []Egg
var sales []Sale

func mod(i, j int) int {
	return i % j
}
func add(i, j int) int {
	return i + j
}
func sortEggsByGeckoID(eggs []Egg) []Egg {
	sort.SliceStable(eggs, func(i, j int) bool {
		return eggs[i].GeckoID < eggs[j].GeckoID
	})
	return eggs
}

func main() {
	http.HandleFunc("/styles.css", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "assets/styles.css") })
	http.HandleFunc("/", homepage)
	http.HandleFunc("/newEgg", newEgg)
	http.HandleFunc("/newGecko", newGecko)
	http.HandleFunc("/newSale", newSale)
	http.ListenAndServe(":8080", nil)

}

func homepage(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{
		"mod":               mod,
		"add":               add,
		"sortEggsByGeckoID": sortEggsByGeckoID,
	}
	for i := range eggs {
		eggs[i].FormattedLayDate = eggs[i].LayDate.Format("02-01-2006")
		eggs[i].FormattedHatchDateETA = GetHatchETAString(&eggs[i]) // Assuming HatchDate is the ETA
	}
	data := TemplateData{
		Eggs:          eggs,
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
	tpl := template.Must(template.ParseFiles("assets/new_egg.html"))
	if r.Method != http.MethodPost {
		var availableGeckos []int
		for _, gecko := range geckos {
			availableGeckos = append(availableGeckos, gecko.Id)
		}
		tpl.Execute(w, map[string]interface{}{
			"AvailableGeckos": availableGeckos,
		})
		return
	}
	geckoId, _ := strconv.Atoi(r.FormValue("gecko"))
	eggCount, _ := strconv.Atoi(r.FormValue("eggCount"))
	date, _ := time.Parse("2006-01-02", r.FormValue("date"))
	AddEgg(geckoId, eggCount, date.Format("02/01/2006"), "")

	// tpl.Execute(w, struct{ Success bool }{true})
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func GetNextLayDateInfo() string {
	var LayETA time.Time = time.Now().Add(time.Hour * 99999)
	var geckoId int = 0
	for _, gecko := range geckos {
		if gecko.GetLayETA().Before(LayETA) {
			LayETA = gecko.GetLayETA()
			geckoId = gecko.Id
		}
	}

	return LayETA.Format("02-01-2006") + " (gecko " + strconv.Itoa(geckoId) + ")"
}

func GetNextHatchDateInfo() string {
	var HatchETA time.Time = time.Now().Add(time.Hour * 99999)
	for _, egg := range eggs {
		if egg.GetHatchETA().Before(HatchETA) {
			HatchETA = egg.GetHatchETA()
		}

	}

	return HatchETA.Format("02-01-2006")
}

func newSale(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("assets/new_sale.html"))
	if r.Method != http.MethodPost {
		tpl.Execute(w, nil)
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
