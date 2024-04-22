package main

import (
	"fmt"
	"strings"
	"time"
)

// Get timestamp for now
func timestamp() string {
	var timestamp time.Time = time.Now()
	return timestamp.Format(time.Stamp)
}

// Create text from arguments (array)
func argText(args []string) string {
	var txt string = ""
	for i, s := range args {
		if i < 1 {
			txt = s
		} else {
			txt = txt + " " + s
		}
	}
	return txt
}

// Check if db has a selected list
func getSelectedListIndex(dbLists []List) int {
	for i, l := range dbLists {
		if l.SelectedList {
			return i
		}
	}
	return -1
}

// Show a table of available lists in db
func shwDB(db DB) {
	longestName, longestDate := 0, 0
	for _, l := range db.ListArr {
		if len(l.ListName) > longestName {
			if l.SelectedList {
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
	for _, l := range db.ListArr {
		if l.SelectedList {
			fmt.Printf("| %-*d | %-*s | %-*s |\n", 1, l.Index, longestName, "-> "+l.ListName, longestDate, l.LastModified)
		} else {
			fmt.Printf("| %-*d | %-*s | %-*s |\n", 1, l.Index, longestName, l.ListName, longestDate, l.LastModified)
		}
	}
	fmt.Println(strings.Repeat("-", 11+longestName+longestDate))
}

// Print instructions
func printInstructions() {
	fmt.Println("Lists: ./godo [checkout] [-d (delete)] [listname]")
	fmt.Println("List Entries: ./godo [add] [ordr] [del] [shw]")
}
