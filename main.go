package main

import (
	"os"
)

type DB struct {
	Db_name        string `json:"Db_name"`
	DB_CurrentList string `json:"DB_CurrentList"`
	Db_lists       []List `json:"Db_lists"`
}
type List struct {
	Index         uint    `json:"index"`
	List_name     string  `json:"List_name"`
	Last_modified string  `json:"Last_modified"`
	List          []Entry `json:"List"`
}
type Entry struct {
	Index     uint   `json:"index"`
	Entry     string `json:"todoItem"`
	Date      string `json:"dateAdded"`
	Important bool   `json:"important"`
}

func main() {
	var db DB
	var currentList List //this is how I'll keep track of the current list I'm working on - use index
	var entry Entry

	db.getFileJSON()

	args := os.Args
	if len(args) > 1 {
		switch args[1] {

		//DB Commands
		case "db-name": // Change DB name
			db.changeDbName()
		case "db-shw": // Show table of lists
			db.shwDB()
		case "checkout": // Create/Switch lists
			db.checkout(args)

		//List Commands
		case "shw":
			currentList.shw()
		case "add":
			entry.add(args)
			currentList.List = append(currentList.List, entry)
			db.pushFileJSON()
		case "del":
			currentList.delTodoItem(args)
			db.pushFileJSON()
		case "ordr":
			currentList.orderBy()
			db.pushFileJSON()
		default:
			printInstructions()
		}
	} else {
		printInstructions()
	}
}
