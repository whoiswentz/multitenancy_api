package main

import (
	"mongodb-test/db"
)

func main() {
	anon := db.New("mongodb://localhost:27017")
	sec := db.New("mongodb://localhost:27018")
	noc := db.New("mongodb://localhost:27019")

	cancel := anon.Connect()
	cancel = sec.Connect()
	cancel = noc.Connect()
	
	cancel()
}
