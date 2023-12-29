package main

import (
	"log"
	"net/url"

	"github.com/mitchellh/mapstructure"
	"github.com/zemirco/couchdb"
)

// create your own document
type CouchDBDocument struct {
	couchdb.Document
	Type      string
	Gecko     Gecko
	Incubator Incubator
	Egg       Egg
	Sale      Sale
}

func WriteToDB(dataType string, gecko Gecko, incubator Incubator, egg Egg, sale Sale) {
	u, err := url.Parse(config.Database.Url)
	if err != nil {
		panic(err)
	}

	// connect
	client, err := couchdb.NewAuthClient(config.Database.Username, config.Database.Password, u)
	if err != nil {
		panic(err)
	}

	// use database and create a document
	db := client.Use(config.Database.Name)
	doc := &CouchDBDocument{
		Type:      dataType,
		Gecko:     gecko,
		Incubator: incubator,
		Egg:       egg,
		Sale:      sale,
	}
	result, err := db.Post(doc)
	if err != nil {
		panic(err)
	}

	// get id and current revision.
	if err := db.Get(doc, result.ID); err != nil {
		panic(err)
	}

	// // delete document
	// if _, err = db.Delete(doc); err != nil {
	// 	panic(err)
	// }

	// // and finally delete the database
	// if _, err = client.Delete("dummy"); err != nil {
	// 	panic(err)
	// }

}

func LoadFromDB() {
	u, err := url.Parse(config.Database.Url)
	if err != nil {
		panic(err)
	}

	// connect
	client, err := couchdb.NewAuthClient(config.Database.Username, config.Database.Password, u)
	if err != nil {
		panic(err)
	}

	// create the database if it doesnt exist
	client.Create(config.Database.Name)

	log.Printf("Connected to Database %s%s with username '%s'", config.Database.Url, config.Database.Name, config.Database.Username)

	db := client.Use(config.Database.Name)
	result, _ := db.AllDocs(&couchdb.QueryParameters{IncludeDocs: &[]bool{true}[0]})
	var data CouchDBDocument
	for _, row := range result.Rows {
		mapstructure.Decode(row.Doc, &data)

		switch dataType := data.Type; {
		case dataType == "egg":
			LoadEgg(data.Egg.ID, data.Egg.IncubatorID, data.Egg.Incubator.Row, data.Egg.Incubator.Column, data.Egg.GeckoID, data.Egg.Count, data.Egg.FormattedLayDate, data.Egg.FormattedHatchDateETA, data.Egg.HasHatched)
		case dataType == "gecko":
			LoadGecko(data.Gecko.ID, data.Gecko.Description)
		case dataType == "incubator":
			LoadIncubator(data.Incubator.ID, data.Incubator.Rows, data.Incubator.Columns)
		case dataType == "sale":
			LoadSale(data.Sale.Buyer, data.Sale.Source, data.Sale.Male, data.Sale.Female, data.Sale.Baby, data.Sale.TotalPrice, data.Sale.Date)
		default:
			log.Println("Unknown data type: " + dataType)
		}
	}
}
