package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	var db DB = getFileJSON()

	// Use function to assign current list to a list value
	var currentList *List = &db.ListArr[getSelectedListIndex(db.ListArr)]

	var entry Entry

	args := os.Args
	if len(args) > 1 {
		switch args[1] {

		//DB Commands
		case "db-name": // Change DB name
			db.changeDbName()

		//DB List Commands
		case "checkout": // Create/Switch lists
			db.checkout(args)

		//List Entry Commands
		case "shw":
			currentList.shwEntries()
		case "add":
			entry.addEntry(args)
			currentList.List = append(currentList.List, entry)
			db.pushFileJSON()
		case "del":
			currentList.delEntry(args)
			db.pushFileJSON()
		case "ordr":
			currentList.orderEntries()
			db.pushFileJSON()
		default:
			printInstructions()
		}
		//db.pushFileJSON()

	} else {
		printInstructions()
	}
}

// Change the name of the DB
func (db *DB) changeDbName() {
	fmt.Print("New DB name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if scanner.Err() != nil {
		log.Fatalf("Change DB name error: %s", scanner.Err())
	}

	db.DbName = scanner.Text()
	db.pushFileJSON()
}
