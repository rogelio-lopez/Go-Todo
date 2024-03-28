package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Files
func (t *Todo) getFileJSON() {
	var info Todo

	if !fileExists("todo-db.json") {
		fmt.Print("hello \n")
		//createDB()
	}

	fileData, err := os.ReadFile("todo-list.json")
	if err != nil {
		log.Fatalf("ReadFile Error: %s\n", err)
	}

	unmarshErr := json.Unmarshal(fileData, &info)
	if unmarshErr != nil {
		log.Fatalf("Unmarshal Error: %s\n", unmarshErr)
	}

	t.UserName = info.UserName
	t.List = info.List
}

func (t *Todo) pushFileJSON() {
	listAsByte, err := json.MarshalIndent(t, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("todo-list.json", listAsByte, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func fileExists(path string) bool {
	_, err := os.Open(path)
	return err == nil
}

func createDB() {
	fmt.Println("Creating Todo DB")

	fmt.Print("DB Name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if scanner.Err() == nil {
		log.Fatalln("Scanning Error in createDB()")
	}
}
