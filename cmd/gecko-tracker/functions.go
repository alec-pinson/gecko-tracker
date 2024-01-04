package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"
)

type NextLayDate struct {
	Value  string
	Colour string
}

func GetNextLayDateInfo() NextLayDate {
	var LayETA time.Time = time.Now().Add(time.Hour * 999999)
	var nextLayDate NextLayDate
	nextLayDate.Colour = "000000" // black
	var geckoId int = 0
	var geckoDescription string = ""
	for _, gecko := range geckos {
		if gecko.Gender == "female" && !gecko.Deleted {
			if gecko.GetLayETA().Before(LayETA) {
				LayETA = gecko.GetLayETA()
				geckoId = gecko.ID
				geckoDescription = gecko.Description
			}
		}
	}

	if len(eggs) == 0 {
		nextLayDate.Value = "No gecko eggs have been recorded"
		return nextLayDate
	}

	if geckoId == 0 {
		// no geckos found
		nextLayDate.Value = "No female geckos found"
		return nextLayDate
	}

	if time.Now().After(LayETA) {
		// if lay date is after current date
		// make it red
		nextLayDate.Colour = "#FF0000"
	} else if time.Now().Add(time.Hour * (2 * 24)).After(LayETA) {
		// if lay date is in the next 2 days
		// make it amber
		nextLayDate.Colour = "#FF9B00"
	} else {
		// normal colour
		nextLayDate.Colour = "000000" // black
	}

	if geckoDescription == "" {
		nextLayDate.Value = LayETA.Format("02-01-2006") + " (gecko " + strconv.Itoa(geckoId) + ")"
	} else {
		nextLayDate.Value = LayETA.Format("02-01-2006") + " (gecko " + strconv.Itoa(geckoId) + " - " + geckoDescription + ")"
	}

	return nextLayDate
}

type NextHatchDate struct {
	Value  string
	Colour string
}

func GetNextHatchDateInfo() NextHatchDate {
	var HatchETA time.Time = time.Now().Add(time.Hour * 999999)
	var nextHatchDate NextHatchDate
	nextHatchDate.Colour = "000000" // black
	var unhatchedEggsFound = false
	for _, egg := range eggs {
		if egg.GetHatchETA().Before(HatchETA) && !egg.HasHatched {
			HatchETA = egg.GetHatchETA()
		}
		if !egg.HasHatched {
			unhatchedEggsFound = true
		}
	}

	if len(eggs) == 0 {
		nextHatchDate.Value = "No gecko eggs have been recorded"
		return nextHatchDate
	}

	// check if any eggs waiting to hatch
	if !unhatchedEggsFound {
		nextHatchDate.Value = "No eggs waiting to hatch"
		return nextHatchDate
	}

	if time.Now().After(HatchETA) {
		// if hatch date is after current date
		// make it red
		nextHatchDate.Colour = "#FF0000"
	} else if time.Now().Add(time.Hour * (2 * 24)).After(HatchETA) {
		// if hatch date is in the next 2 days
		// make it amber
		nextHatchDate.Colour = "#FF9B00"
	} else {
		// normal colour
		nextHatchDate.Colour = "000000" // black
	}

	nextHatchDate.Value = HatchETA.Format("02-01-2006")

	return nextHatchDate
}

func GetAverageHatchTimeString() string {
	ret, _, _ := GetAverageHatchTimeInfo()
	return ret
}

func GetAverageHatchTimeDays() int {
	_, ret, _ := GetAverageHatchTimeInfo()
	return ret
}

func GetAverageHatchTimeDuration() time.Duration {
	_, _, ret := GetAverageHatchTimeInfo()
	return ret
}

func GetAverageHatchTimeInfo() (string, int, time.Duration) {
	var hatchTimeSum, hatchedEggTotal float64
	for _, egg := range eggs {
		if egg.HasHatched {
			hatchTimeSum += egg.HatchDate.Sub(egg.LayDate).Hours() / 24
			hatchedEggTotal += 1
		}
	}

	if hatchedEggTotal == 0 {
		return fmt.Sprintf("%.0f days", config.HatchTime.Hours()/24), int(config.HatchTime.Hours() / 24), config.HatchTime
	}

	hatchTimeAverage := hatchTimeSum / hatchedEggTotal

	return fmt.Sprintf("%.0f days", hatchTimeAverage), int(hatchTimeAverage), time.Duration(time.Hour * 24 * time.Duration(hatchTimeAverage))
}

func N(start, end int) (stream chan int) {
	stream = make(chan int)
	go func() {
		for i := start; i <= end; i++ {
			stream <- i
		}
		close(stream)
	}()
	return
}

func sortEggsIntoGrid(eggs []Egg) map[string]Egg {
	eggMap := make(map[string]Egg)

	for _, egg := range eggs {
		eggMap[strconv.Itoa(egg.IncubatorID)+","+strconv.Itoa(egg.Incubator.Row)+","+strconv.Itoa(egg.Incubator.Column)] = egg
	}

	return eggMap
}

func toSlotID(incubatorId int, row int, column int) string {
	return fmt.Sprintf("%s,%s,%s", strconv.Itoa(incubatorId), strconv.Itoa(row), strconv.Itoa(column))
}

func Title(str string) string {
	return strings.Title(str)
}

func DaysToHours(days string) string {
	iDays, err := strconv.Atoi(strings.TrimSuffix(days, "d"))
	if err != nil {
		log.Println(err)
		return ""
	}
	iHours := iDays * 24

	return strconv.Itoa(iHours) + "h"
}

func GetEggRAG(egg Egg) string {
	if time.Now().After(egg.GetHatchETA()) {
		// if hatch date is after current date
		// make it red
		return "#FF0000"
	} else if time.Now().Add(time.Hour * (2 * 24)).After(egg.GetHatchETA()) {
		// if hatch date is in the next 2 days
		// make it amber
		return "#FF9B00"
	} else {
		// normal colour
		return "#008721" // green
	}
}

func TimeDiff(a, b time.Time) (year, month, day, hour, min, sec int) {
	if a.Location() != b.Location() {
		b = b.In(a.Location())
	}
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	h1, m1, s1 := a.Clock()
	h2, m2, s2 := b.Clock()

	year = int(y2 - y1)
	month = int(M2 - M1)
	day = int(d2 - d1)
	hour = int(h2 - h1)
	min = int(m2 - m1)
	sec = int(s2 - s1)

	// Normalize negative values
	if sec < 0 {
		sec += 60
		min--
	}
	if min < 0 {
		min += 60
		hour--
	}
	if hour < 0 {
		hour += 24
		day--
	}
	if day < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		day += 32 - t.Day()
		month--
	}
	if month < 0 {
		month += 12
		year--
	}

	return
}

func GetAge(d time.Time) string {
	year, month, day, hour, min, _ := TimeDiff(time.Now(), d)

	switch true {
	// 1 minute old
	case year == 0 && month == 0 && day == 0 && hour == 0 && min == 1:
		return "1 minute"
	// x minutes old
	case year == 0 && month == 0 && day == 0 && hour == 0 && min != 1:
		return fmt.Sprintf("%d minutes", min)
	// 1 hour old
	case year == 0 && month == 0 && day == 0 && hour == 1:
		return "1 hour"
	// x hours old
	case year == 0 && month == 0 && day == 0 && hour != 1:
		return fmt.Sprintf("%d hours", hour)
	// 1 day old
	case year == 0 && month == 0 && day == 1:
		return "1 day"
	// x days old
	case year == 0 && month == 0 && day != 1:
		return fmt.Sprintf("%d days", day)
	// 1 month old
	case year == 0 && month == 1:
		return "1 month"
	// x months old
	case year == 0 && month != 1:
		return fmt.Sprintf("%d months", month)
	// 1 year old
	case year == 1 && month == 0:
		return "1 year"
	// 1 year, 1 month old
	case year == 1 && month == 1:
		return "1 year, 1 month"
	// 1 year, x months old
	case year == 1 && month != 1:
		return fmt.Sprintf("1 year, %d months", month)
	// x years old
	case year > 1 && month == 0:
		return fmt.Sprintf("%d years", year)
	// x years old, 1 month old
	case year > 1 && month == 1:
		return fmt.Sprintf("%d years, 1 month", year)
	// x years old, x months old
	case year > 1 && month != 1:
		return fmt.Sprintf("%d years, %d months", year, month)
	}

	return ""
}
