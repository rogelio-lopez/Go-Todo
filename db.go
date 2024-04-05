package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Change the name of the DB
func (db *DB) changeDbName() {
	fmt.Print("New DB name: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if scanner.Err() != nil {
		log.Fatalf("Change DB name error: %s", scanner.Err())
	}

	db.Db_name = scanner.Text()
	db.pushFileJSON()
}

/*--- Checkout ---*/
func (db *DB) checkout(args []string) {
	if len(args) >= 3 {
		switch args[2] {
		case "-d":
			db.delList(args)
		default:
			db.addList(args)
		}
	} else {
		shwDB(*db)
	}
}

func (db *DB) delList(args []string) {
	if len(args) >= 4 {
		var listName = argText(args[3:])
		// <-------- WORKING ON THIS NEXT
	} else {
		fmt.Println("Ya need to give me the name of the list to delte")
	}
}

func (db *DB) addList(args []string) {
	var listName string = argText(args[2:])

	// Create new list if name != a list name
	if db.changeCurrentList(listName) < 0 {
		var newList List
		newList.Index = uint(len(db.Db_lists))
		newList.List_name = listName
		newList.Last_modified = timestamp()
		newList.List = []Entry{}

		db.Db_lists = append(db.Db_lists, newList)
		db.pushFileJSON()
	}
}

func (db *DB) changeCurrentList(listName string) int {
	for i, l := range db.Db_lists {
		if l.List_name == listName {
			db.DB_CurrentList.Name = listName
			db.DB_CurrentList.Index = uint(i)
			db.pushFileJSON()
			return i
		}
	}
	return -1
}

// Show a table of available lists
func shwDB(db DB) {
	longestName, longestDate := 0, 0
	for i, l := range db.Db_lists {
		if len(l.List_name) > longestName {
			if i == int(db.DB_CurrentList.Index) {
				longestName = len(l.List_name) + 3
			} else {
				longestName = len(l.List_name)
			}
		}
		if len(l.Last_modified) > longestDate {
			longestDate = len(l.Last_modified)
		}
	}

	if len(db.Db_lists) < 1 {
		fmt.Println("Your DB is empty ¯\\_(ツ)_/¯")
		return
	}

	if db.Db_name != "" {
		fmt.Printf("%s\n", db.Db_name)
	} else {
		fmt.Println("My Lists")
	}

	fmt.Println(strings.Repeat("-", 11+longestName+longestDate))
	for i, l := range db.Db_lists {
		if i == int(db.DB_CurrentList.Index) {
			fmt.Printf("| %-*d | %-*s | %-*s |\n", 1, l.Index, longestName, "-> "+l.List_name, longestDate, l.Last_modified)
		} else {
			fmt.Printf("| %-*d | %-*s | %-*s |\n", 1, l.Index, longestName, l.List_name, longestDate, l.Last_modified)
		}
	}
	fmt.Println(strings.Repeat("-", 11+longestName+longestDate))
}
