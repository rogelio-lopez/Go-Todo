package main

import (
	"os"
)

func main() {
	var db DB = getFileJSON()

	// Use function to assign current list to a list value
	listIndex := db.ThisList.Index
	var currentList List = db.ListArr[listIndex]
	var entry Entry

	args := os.Args
	if len(args) > 1 {
		switch args[1] {

		//DB Commands
		case "db-name": // Change DB name
			db.changeDbName()
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
