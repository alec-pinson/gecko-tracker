package main

import (
	"html/template"
	"net/http"
	"strconv"
	"time"
)

type EggSlot struct {
	Gecko    int
	EggCount int
	LayDate  string
	HatchEta string
}

var geckos []Gecko
var eggs []Egg
var sales []Sale

func main() {
	http.HandleFunc("/styles.css", func(w http.ResponseWriter, r *http.Request) { http.ServeFile(w, r, "assets/styles.css") })
	http.HandleFunc("/", homepage)
	http.HandleFunc("/newEgg", newEgg)
	http.HandleFunc("/newGecko", newGecko)
	http.HandleFunc("/newSale", newSale)
	http.ListenAndServe(":8080", nil)
}

func homepage(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("assets/home.html"))
	tpl.Execute(w, struct {
		TotalSales    string
		NextLayDate   string
		NextHatchDate string
		Slot1         EggSlot
		Slot2         EggSlot
		Slot3         EggSlot
		Slot4         EggSlot
		Slot5         EggSlot
		Slot6         EggSlot
	}{
		TotalSales:    TotalSales(),
		NextLayDate:   GetNextLayDateInfo(),
		NextHatchDate: GetNextHatchDateInfo(),
		Slot1: EggSlot{
			Gecko:    GetEgg(1).GeckoID,
			EggCount: GetEgg(1).Count,
			LayDate:  GetEgg(1).GetLayDateString(),
			HatchEta: GetEgg(1).GetHatchETAString(),
		},
		Slot2: EggSlot{
			Gecko:    GetEgg(2).GeckoID,
			EggCount: GetEgg(2).Count,
			LayDate:  GetEgg(2).GetLayDateString(),
			HatchEta: GetEgg(2).GetHatchETAString(),
		},
		Slot3: EggSlot{
			Gecko:    GetEgg(3).GeckoID,
			EggCount: GetEgg(3).Count,
			LayDate:  GetEgg(3).GetLayDateString(),
			HatchEta: GetEgg(3).GetHatchETAString(),
		},
		Slot4: EggSlot{
			Gecko:    GetEgg(4).GeckoID,
			EggCount: GetEgg(4).Count,
			LayDate:  GetEgg(4).GetLayDateString(),
			HatchEta: GetEgg(4).GetHatchETAString(),
		},
		Slot5: EggSlot{
			Gecko:    GetEgg(5).GeckoID,
			EggCount: GetEgg(5).Count,
			LayDate:  GetEgg(5).GetLayDateString(),
			HatchEta: GetEgg(5).GetHatchETAString(),
		},
		Slot6: EggSlot{
			Gecko:    GetEgg(6).GeckoID,
			EggCount: GetEgg(6).Count,
			LayDate:  GetEgg(6).GetLayDateString(),
			HatchEta: GetEgg(6).GetHatchETAString(),
		},
	})
}

func newGecko(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("assets/new_gecko.html"))
	if r.Method != http.MethodPost {
		tpl.Execute(w, nil)
		return
	}

	geckoId, _ := strconv.Atoi(r.FormValue("geckoId"))
	AddGecko(geckoId, r.FormValue("description"))

	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func newEgg(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("assets/new_egg.html"))
	if r.Method != http.MethodPost {
		tpl.Execute(w, nil)
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
	var slotId int = 0
	for _, egg := range eggs {
		if egg.SlotId != 0 {
			// only eggs that are incubating
			if egg.GetHatchETA().Before(HatchETA) {
				HatchETA = egg.GetHatchETA()
				slotId = egg.SlotId
			}
		}
	}

	return HatchETA.Format("02-01-2006") + " (slot " + strconv.Itoa(slotId) + ")"
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
