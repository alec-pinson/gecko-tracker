package main

import (
	"sort"
	"strconv"
	"time"
)

func GetNextLayDateInfo() string {
	var LayETA time.Time = time.Now().Add(time.Hour * 99999)
	var geckoId int = 0
	for _, gecko := range geckos {
		if gecko.GetLayETA().Before(LayETA) {
			LayETA = gecko.GetLayETA()
			geckoId = gecko.ID
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
