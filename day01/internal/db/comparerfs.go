package db

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
)

func CompareFileSystem(oldBackUp, newBackUp string) {
	oldData, err := getFilesMap(oldBackUp)
	if err != nil {
		log.Fatal(err)
	}

	newData, err := getFilesMap(newBackUp)
	if err != nil {
		log.Fatal(err)
	}

	for key, _ := range newData {
		if _, ok := oldData[key]; !ok {
			fmt.Printf("ADDED: %s\n", key)
		}
	}

	for key, _ := range oldData {
		if _, ok := newData[key]; !ok {
			fmt.Printf("REMOVED: %s\n", key)
		}
	}
}

func getFilesMap(filename string) (map[string]bool, error) {
	data, err := ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	dataMap := make(map[string]bool)

	// Parse []byte to strings
	bytesReader := bytes.NewReader(data)
	reader := bufio.NewReader(bytesReader)
	for {
		line, _, err := reader.ReadLine()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
		dataMap[string(line)] = true
	}

	return dataMap, nil
}
