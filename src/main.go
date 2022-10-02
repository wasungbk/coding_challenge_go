package main

import (
	"encoding/csv"
	"log"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	// verify file path exists in arguments
	if len(os.Args) != 2 {
		log.Fatal("must contain file path as argument")
	}

	//open csv file
	filePath := os.Args[1]
	err := readFile(filePath)
	if err != nil {
		log.Fatal("unable to read file")
	}
}

func readFile(filePath string) error {
	f, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer f.Close()
	reader := csv.NewReader(f)
	headers, err := reader.Read()
	if err != nil {
		log.Fatal("could not read csv file")
	}
	if checkHeaders(headers, []string{"timestamp", "username", "operation", "size"}) {
		log.Fatal("unexpected headers in csv. expected: timestamp,username,operation,size")
	}

	//read until error from csv
	for {

		line, err := reader.Read()
		if err == io.EOF {
			break
		}

	}

	// if err := scanner.Err(); err != nil {
	// 	log.Fatal(err)
	// }
	return nil
}

func checkHeaders(headers, expected []string) bool {
	if len(headers) != len(expected) {
		return false
	}
	for i := range headers {
		if headers[i] != expected[i] {
			return false
		}
	}
	return true
}
