package main

import (
	"os"
)

type DB struct {
	Db_name  string `json:"Db_name"`
	Db_lists []List `json:"Db_lists"`
}
type List struct {
	List_name string  `json:"List_name"`
	List      []Entry `json:"List"`
}
type Entry struct {
	Entry     string `json:"todoItem"`
	Date      string `json:"dateAdded"`
	Important bool   `json:"important"`
}

func main() {
	var db DB
	var l List
	var e Entry

	db.getFileJSON()

	args := os.Args
	if len(args) > 1 {
		switch args[1] {
		case "shw":
			l.printList()
		case "add":
			e.addTodoItem(args)
			l.List = append(l.List, e)
			db.pushFileJSON()
		case "del":
			l.delTodoItem(args)
			db.pushFileJSON()
			/*
				case "usr":
					l.addUser(args)
					db.pushFileJSON()
			*/
		case "ordr":
			l.orderBy()
			db.pushFileJSON()
		default:
			printInstructions()
		}
	} else {
		printInstructions()
		//todo.printList()
	}
}
