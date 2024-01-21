package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"time"
)

type Gecko struct {
	ID                   int    `json:"id"`
	Description          string `json:"description"`
	TankID               int    `json:"tankId"`
	Gender               string `json:"gender"` // male, female, baby
	DateOfBirth          time.Time
	FormattedDateOfBirth string `json:"formattedDateOfBirth"`
	Age                  string
	Deleted              bool `json:"deleted"`
}

func AddGecko(description string, tankId int, gender string, dateOfBirth string) (*Gecko, error) {
	// // Check if a gecko with the same ID already exists
	// for _, existingGecko := range geckos {
	// 	if existingGecko.ID == id {
	// 		return nil, errors.New("Gecko with the same ID already exists")
	// 	}
	// }

	var gecko Gecko
	gecko.ID = len(geckos) + 1
	gecko.Description = description
	gecko.TankID = tankId
	gecko.Gender = gender
	DateOfBirth, err := time.Parse("02/01/2006", dateOfBirth)
	if err != nil {
		log.Println(err)
	}
	gecko.DateOfBirth = DateOfBirth
	gecko.FormattedDateOfBirth = dateOfBirth
	gecko.Deleted = false
	geckos = append(geckos, &gecko)

	log.Println("Added gecko '" + description + "' (id: " + strconv.Itoa(gecko.ID) + ")")

	WriteToDB("gecko", gecko, Incubator{}, Egg{}, Sale{}, Tank{}, Notifications{})

	return &gecko, nil
}

func LoadGecko(id int, description string, tankId int, gender string, dateOfBirth string, deleted bool) (*Gecko, error) {
	var gecko Gecko
	gecko.ID = id
	gecko.Description = description
	gecko.TankID = tankId
	gecko.Gender = gender
	DateOfBirth, err := time.Parse("02/01/2006", dateOfBirth)
	if err != nil {
		log.Println(err)
	}
	gecko.DateOfBirth = DateOfBirth
	gecko.FormattedDateOfBirth = dateOfBirth
	gecko.Deleted = deleted
	geckos = append(geckos, &gecko)

	log.Println("Loaded gecko '" + description + "' (id: " + strconv.Itoa(id) + ")")

	return &gecko, nil
}

func (gecko *Gecko) Update() {
	UpdateDB("gecko", *gecko, Incubator{}, Egg{}, Sale{}, Tank{}, Notifications{})
	log.Print("Updated gecko '" + gecko.Description + "' (id: " + strconv.Itoa(gecko.ID) + ")")
}

func (gecko *Gecko) Delete() {
	if !gecko.Deleted {
		gecko.Deleted = true
		UpdateDB("gecko", *gecko, Incubator{}, Egg{}, Sale{}, Tank{}, Notifications{})
		log.Print("Marked gecko '" + gecko.Description + "' (id: " + strconv.Itoa(gecko.ID) + ") as deleted")
	}
}

func GetGecko(id int) (*Gecko, error) {
	for _, gecko := range geckos {
		if gecko.ID == id {
			return gecko, nil
		}
	}
	log.Println("Gecko not found: " + strconv.Itoa(id))
	return &Gecko{}, fmt.Errorf("Gecko not found with ID: %d", id)
}

func (gecko Gecko) GetLastLayDate() time.Time {
	eggsLaid := gecko.GetLayHistory()
	var lastLayDate time.Time
	for _, egg := range eggsLaid {
		if egg.LayDate.After(lastLayDate) {
			lastLayDate = egg.LayDate
		}
	}

	return lastLayDate
}

func (gecko Gecko) GetLayETA() time.Time {
	return gecko.GetLastLayDate().Add(gecko.GetAverageLayTimeDuration())
}

func (gecko Gecko) GetLayHistory() []Egg {
	var eggsLaid []Egg
	for _, egg := range eggs {
		if egg.GeckoID == gecko.ID {
			eggsLaid = append(eggsLaid, *egg)
		}
	}

	return eggsLaid
}

func (gecko Gecko) GetAverageLayTimeString() string {
	ret, _, _, _ := gecko.GetAverageLayInfo()
	return ret
}

func (gecko Gecko) GetAverageLayTimeDays() int {
	_, ret, _, _ := gecko.GetAverageLayInfo()
	return ret
}

func (gecko Gecko) GetAverageLayTimeDuration() time.Duration {
	_, _, ret, _ := gecko.GetAverageLayInfo()
	return ret
}

func (gecko Gecko) GetNextLayDate() string {
	_, _, _, ret := gecko.GetAverageLayInfo()
	return ret
}

func (gecko Gecko) GetAverageLayInfo() (string, int, time.Duration, string) {
	var eggsLaid []*Egg
	var skipFirst bool = false
	var lastLayDate time.Time
	var LayTimeSum, LayTotal float64

	// get eggs laid by this gecko
	for _, egg := range eggs {
		if egg.GeckoID == gecko.ID {
			eggsLaid = append(eggsLaid, egg)
		}
	}

	if len(eggsLaid) == 0 {
		// never laid any eggs before so use default lay time
		nextLayDate := lastLayDate.Add(time.Duration(config.LayTime)).Format("02/01/2006")
		return fmt.Sprintf("%.0f days", config.LayTime.Hours()/24), int(config.LayTime.Hours() / 24), config.LayTime, nextLayDate
	}

	// sort eggs by lay date
	sort.Slice(eggsLaid, func(i, j int) bool {
		return eggsLaid[i].LayDate.Before(eggsLaid[j].LayDate)
	})

	// get difference between lay dates
	for _, egg := range eggsLaid {
		// they are in order ignore if this isnt true
		if !skipFirst {
			skipFirst = true
			lastLayDate = egg.LayDate
			continue
		}
		LayTimeSum += egg.LayDate.Sub(lastLayDate).Hours() / 24
		LayTotal += 1
		lastLayDate = egg.LayDate
	}

	layTimeAverage := LayTimeSum / LayTotal

	nextLayDate := lastLayDate.Add(time.Duration(time.Hour * 24 * time.Duration(layTimeAverage))).Format("02/01/2006")

	return fmt.Sprintf("%.0f days", layTimeAverage), int(layTimeAverage), time.Duration(time.Hour * 24 * time.Duration(layTimeAverage)), nextLayDate
}
