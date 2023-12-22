package main

import (
	"log"
	"strconv"
	"time"
)

type Sale struct {
	Buyer      string
	Source     string
	Male       int
	Female     int
	Baby       int
	TotalPrice int
	Date       time.Time
}

func AddSale(buyer string, source string, male int, female int, baby int, totalPrice int, date string) *Sale {
	var sale Sale
	sale.Buyer = buyer
	sale.Source = source
	sale.Male = male
	sale.Female = female
	sale.Baby = baby
	sale.TotalPrice = totalPrice
	SaleDate, err := time.Parse("02/01/2006", date)
	if err != nil {
		log.Println(err)
	}
	sale.Date = SaleDate

	sales = append(sales, sale)

	log.Println("Added new sale £" + strconv.Itoa(sale.TotalPrice))

	return &sale
}

func TotalSales() string {
	total := 0
	for _, sale := range sales {
		total += sale.TotalPrice
	}

	return "£" + strconv.Itoa(total)
}
