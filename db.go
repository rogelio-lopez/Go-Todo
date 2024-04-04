package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Create/Switch current list
func (db *DB) checkout(args []string) {
	if len(args) >= 3 {
		switch args[2] {
		case "-d":
			db.delList(args)
		default:
			db.addList(args)
		}
	} else {
		fmt.Printf("You're on list: %s\n", db.DB_CurrentList)
	}
}

func (db *DB) delList(args []string) {
	if len(args) >= 4 {

	} else {
		fmt.Println("Ya need to give me the name of the list to delte")
	}
}

func (db *DB) addList(args []string) {
	var name string = ""

	for i, s := range args[2:] {
		if i < 1 {
			name = s
		} else {
			name = name + " " + s
		}
	}

	if db.changeCurrentList(name) < 0 {
		//create the list
		var newList List
		newList.Index = 2
		newList.List_name = name
		newList.Last_modified = timestamp()
		newList.List = []Entry{}

		db.Db_lists = append(db.Db_lists, newList)
		db.pushFileJSON()
	} else {
		db.changeCurrentList(name)
	}
}

/* -------> working on this */
func (db *DB) changeCurrentList(listName string) int {
	for i, l := range db.Db_lists {
		if l.List_name == listName {
			db.DB_CurrentList = listName
			return i
		}
	
	return -1
}

// Delete List

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

// Show a table of available lists
func (db *DB) shwDB() {
	longestName, longestDate := 0, 0
	for _, l := range db.Db_lists {
		if len(l.List_name) > longestName {
			longestName = len(l.List_name)
		}
		if len(l.Last_modified) > longestDate {
			longestDate = len(l.Last_modified)
		}
	}

	if len(db.Db_lists) < 1 {
		fmt.Println("Your todo DB is empty ¯\\_(ツ)_/¯")
		return
	}

	if db.Db_name != "" {
		fmt.Printf("%s\n", db.Db_name)
	} else {
		fmt.Println("My Lists")
	}

	fmt.Println(strings.Repeat("-", 11+longestName+longestDate))
	for _, l := range db.Db_lists {
		fmt.Printf("| %-*d | %-*s | %-*s |\n", 1, l.Index, longestName, l.List_name, longestDate, l.Last_modified)
	}
	fmt.Println(strings.Repeat("-", 11+longestName+longestDate))
}
