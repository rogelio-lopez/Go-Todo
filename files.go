package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Files
func getFileJSON() DB {
	var fileDb DB

	if !fileExists("todo-db.json") {
		var fileDb DB = createDB()
		fileDb.pushFileJSON()
	}

	fileData, err := os.ReadFile("todo-db.json")
	if err != nil {
		log.Fatalf("ReadFile Error: %s\n", err)
	}

	unmarshErr := json.Unmarshal(fileData, &fileDb)
	if unmarshErr != nil {
		log.Fatalf("Unmarshal Error: %s\n", unmarshErr)
	}

	return fileDb
}

func (db *DB) pushFileJSON() {
	listAsByte, err := json.MarshalIndent(db, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("todo-db.json", listAsByte, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

func fileExists(path string) bool {
	_, err := os.Open(path)
	return err == nil
}

func createDB() DB {
	file, err := os.Create("todo-db.json")
	if err != nil {
		log.Fatalf("Creeate file error: %s\n", err)
	}
	defer file.Close()

	fmt.Println("Creating Todo DB")

	fmt.Print("DB Name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if scanner.Err() != nil {
		log.Fatalln("Scanning Error in createDB()")
	}
	return DB{
		Db_name: scanner.Text(),
		DB_CurrentList: CurrentList{
			Name:  "New List",
			Index: 0,
		},
		Db_lists: []List{
			{
				Index:         0,
				List_name:     "New List",
				Last_modified: timestamp(),
				List:          []Entry{},
			},
		},
	}
}
