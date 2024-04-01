package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

// Create/Switch current list

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
