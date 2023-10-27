package app

import (
	"Day01/internal/db"
	"Day01/internal/entity"
	"flag"
	"log"
)

func RunEx00() {
	var filename string

	// Parsing flag with filename
	flag.StringVar(&filename, "f", "", "filename")
	flag.Parse()

	rw, recipe, err := db.GetDB(filename)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.PrintDB(rw, recipe); err != nil {
		log.Fatal(err)
	}
}

func RunEx01() {
	var (
		oldFilename string
		newFilename string
	)

	// Parsing flag with filename
	flag.StringVar(&oldFilename, "old", "", "old filename")
	flag.StringVar(&newFilename, "new", "", "new filename")
	flag.Parse()

	_, oldDB, err := db.GetDB(oldFilename)
	if err != nil {
		log.Fatal(err)
	}

	_, newDB, err := db.GetDB(newFilename)
	if err != nil {
		log.Fatal(err)
	}

	// Compare old and new database
	db.CompareDB(oldDB, newDB)
}

func RunEx02() {
	var oldBackUp, newBackUp string

	flag.StringVar(&oldBackUp, "old", "", "file with filepath for all system files")
	flag.StringVar(&newBackUp, "new", "", "file with filepath for all system files")
	flag.Parse()

	if oldBackUp == "" || newBackUp == "" {
		log.Fatal(entity.FileNotFound)
	}

	db.CompareFileSystem(oldBackUp, newBackUp)
}
