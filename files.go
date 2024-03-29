package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Files
func (db *DB) getFileJSON() {
	var fileDb DB

	if !fileExists("todo-db.json") {
		fileDb = createDB()
		fmt.Print(fileDb)
	}

	fileData, err := os.ReadFile("todo-db.json")
	if err != nil {
		log.Fatalf("ReadFile Error: %s\n", err)
	}

	unmarshErr := json.Unmarshal(fileData, &fileDb)
	if unmarshErr != nil {
		log.Fatalf("Unmarshal Error: %s\n", unmarshErr)
	}

	db.Db_name = fileDb.Db_name
	db.Db_lists = fileDb.Db_lists
}

func (db *DB) pushFileJSON() {
	listAsByte, err := json.MarshalIndent(db, "", "\t")
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
		Db_name:  scanner.Text(),
		Db_lists: []List{},
	}
}
