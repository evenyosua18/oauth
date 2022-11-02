package migration

import "log"

func Migrate() {
	//Do Migration
	log.Println("run migration...")

	//Migrate main database
	log.Println("migrate init database")
	if err := InitDatabase(); err != nil {
		panic(err)
	} else {
		log.Println("success migrate init database")
	}
}
