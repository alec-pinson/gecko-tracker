package main

import (
	"log"
	"net/url"

	"github.com/zemirco/couchdb"
)

// create your own document
type CouchDBDocument struct {
	couchdb.Document
	EggData Egg
}

func AddEggToDB(eggData Egg) {
	u, err := url.Parse("https://couchdb.xxxxx.com/")
	if err != nil {
		panic(err)
	}

	// connect
	client, err := couchdb.NewAuthClient("username", "password", u)
	if err != nil {
		panic(err)
	}

	// create a database
	client.Create("eggs")

	// use database and create a document
	db := client.Use("eggs")
	doc := &CouchDBDocument{
		EggData: eggData,
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

func GetEggFromDB() {
	u, err := url.Parse("https://couchdb.xxxxx.com/")
	if err != nil {
		panic(err)
	}

	// connect
	client, err := couchdb.NewAuthClient("username", "password", u)
	if err != nil {
		panic(err)
	}

	var doc couchdb.CouchDoc

	db := client.Use("eggs")
	result, _ := db.AllDocs(&couchdb.QueryParameters{})
	for _, row := range result.Rows {
		log.Println(row.Value)
		db.Get(doc, row.ID)
		log.Println(doc)
		log.Println(db.View("EggData"))
	}
}
