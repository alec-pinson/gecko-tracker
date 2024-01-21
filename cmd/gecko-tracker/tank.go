package main

import (
	"log"
	"sort"
	"strconv"
)

type Tank struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TankContents struct {
	ID     int
	Name   string
	Geckos []*Gecko
}

func AddTank(name string) *Tank {
	var tank Tank
	tank.ID = len(tanks) + 1
	tank.Name = name
	tanks = append(tanks, &tank)

	log.Println("Added new tank, " + name + " (" + strconv.Itoa(tank.ID) + ")")

	WriteToDB("tank", Gecko{}, Incubator{}, Egg{}, Sale{}, tank, Notifications{})

	return &tank
}

func LoadTank(id int, name string) *Tank {
	var tank Tank
	tank.ID = id
	tank.Name = name
	tanks = append(tanks, &tank)

	log.Println("Loaded tank, " + name + " (" + strconv.Itoa(tank.ID) + ")")

	return &tank
}

func (tank Tank) GetTankContents() TankContents {
	var tankContents TankContents
	var tankContentsGeckos []*Gecko
	tankContents.ID = tank.ID
	tankContents.Name = tank.Name
	for _, gecko := range geckos {
		if gecko.TankID == tank.ID && !gecko.Deleted {
			gecko.Age = GetAge(gecko.DateOfBirth)
			tankContentsGeckos = append(tankContentsGeckos, gecko)
		}
	}

	// sort geckos by dob
	sort.Slice(tankContentsGeckos, func(i, j int) bool {
		return tankContentsGeckos[i].DateOfBirth.Before(tankContentsGeckos[j].DateOfBirth)
	})

	tankContents.Geckos = tankContentsGeckos

	return tankContents
}
