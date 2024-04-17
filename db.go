package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

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

		//Handle current list backup if deleted

		for i, l := range db.ListArr {
			if listName == l.ListName {

				if len(db.ListArr) < 2 {
					db.ListArr = []List{}

				} else if i == 0 {
					db.ListArr = db.ListArr[1:]

				} else if i == len(db.ListArr) {
					db.ListArr = db.ListArr[0 : len(db.ListArr)-1]

				} else {
					first := db.ListArr[0:i]
					second := db.ListArr[i+1:]
					db.ListArr = append(first, second...)
				}

			}
		}
	} else {
		fmt.Println("Ya need to give me the name of the list to delete")
	}
	db.refreshIndexes()
	db.pushFileJSON()
}

func (db *DB) addList(args []string) {
	var listName string = argText(args[2:])

	// Create new list if name != a list name
	if db.changeCurrentList(listName) < 0 {
		var newList List
		newList.Index = uint(len(db.ListArr))
		newList.ListName = listName
		newList.LastModified = timestamp()
		newList.List = []Entry{}

		db.ListArr = append(db.ListArr, newList)
		db.pushFileJSON()
	}
}

func (db *DB) changeCurrentList(listName string) int {
	for i, l := range db.ListArr {
		if l.ListName == listName {
			db.ThisList.Name = listName
			db.ThisList.Index = uint(i)
			db.pushFileJSON()
			return i
		}
	}
	return -1
}

// Show a table of available lists
func shwDB(db DB) {
	longestName, longestDate := 0, 0
	for i, l := range db.ListArr {
		if len(l.ListName) > longestName {
			if i == int(db.ThisList.Index) {
				longestName = len(l.ListName) + 3
			} else {
				longestName = len(l.ListName)
			}
		}
		if len(l.LastModified) > longestDate {
			longestDate = len(l.LastModified)
		}
	}

	if len(db.ListArr) < 1 {
		fmt.Println("Your DB is empty ¯\\_(ツ)_/¯")
		return
	}

	if db.DbName != "" {
		fmt.Printf("%s\n", db.DbName)
	} else {
		fmt.Println("My Lists")
	}

	fmt.Println(strings.Repeat("-", 11+longestName+longestDate))
	for i, l := range db.ListArr {
		if i == int(db.ThisList.Index) {
			fmt.Printf("| %-*d | %-*s | %-*s |\n", 1, l.Index, longestName, "-> "+l.ListName, longestDate, l.LastModified)
		} else {
			fmt.Printf("| %-*d | %-*s | %-*s |\n", 1, l.Index, longestName, l.ListName, longestDate, l.LastModified)
		}
	}
	fmt.Println(strings.Repeat("-", 11+longestName+longestDate))
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
