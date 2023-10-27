package db

import (
	"Day01/internal/entity"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

const (
	jsonFormat string = ".json"
	xmlFormat  string = ".xml"
)

type DBReadWriter interface {
	DBReader
	DBWriter
}

type DBReader interface {
	Read(data []byte) (entity.Recipe, error)
}

type DBWriter interface {
	Write(Recipe entity.Recipe) (string, error)
}

func ReadFile(filename string) ([]byte, error) {
	// Check flag with filename
	if len(filename) == 0 {
		return nil, entity.EmptyFilename
	}

	// Open file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read data
	data, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetDB(filename string) (DBReadWriter, entity.Recipe, error) {
	var rw DBReadWriter
	// Get file's format
	extension := filepath.Ext(filename)

	// Define format type of file
	switch extension {
	case jsonFormat:
		rw = NewDBjson()
	case xmlFormat:
		rw = NewDBxml()
	default:
		return nil, entity.Recipe{}, entity.UnexpectedFormat
	}

	// Read data from file
	data, err := ReadFile(filename)
	if err != nil {
		return nil, entity.Recipe{}, err
	}

	// Decode from XML/JSON
	recipe, err := Decode(rw, data)
	if err != nil {
		return nil, entity.Recipe{}, err
	}

	return rw, recipe, nil
}

func PrintDB(rw DBReadWriter, recipe entity.Recipe) error {
	// Encode to JSON/XML
	output, err := Encode(rw, recipe)
	if err != nil {
		return err
	}
	
	fmt.Println("\tEncode data:\n", output)
	return nil
}

func Decode(r DBReader, data []byte) (entity.Recipe, error) {
	return r.Read(data)
}

func Encode(w DBWriter, recipe entity.Recipe) (string, error) {
	return w.Write(recipe)
}
