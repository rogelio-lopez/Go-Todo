package main

import (
	"fmt"
	"strings"
)

/* Prints */
//need to account for "current" list
//param: index of current list
//if shw -a then show all list lists
func (list *List) shw() {
	longestEntry, longestDate := 0, 0
	for _, e := range list.List {
		if len(e.Entry) > longestEntry {
			longestEntry = len(e.Entry)
		}
		if len(e.Date) > longestDate {
			longestDate = len(e.Date)
		}
	}

	if len(list.List) < 1 {
		fmt.Println("Your todo list is empty ¯\\_(ツ)_/¯")
		return
	}

	if list.List_name != "" {
		fmt.Printf("%s's Todo\n", list.List_name)
	} else {
		fmt.Println("My Todo")
	}
	fmt.Println(strings.Repeat("-", 11+longestEntry+longestDate))
	for i, v := range list.List {
		fmt.Printf("| %-*d | %-*s | %-*s |\n", 1, i+1, longestEntry, v.Entry, longestDate, v.Date)
	}
	fmt.Println(strings.Repeat("-", 11+longestEntry+longestDate))
}

func printInstructions() {
	fmt.Println("Go-do Usage: ./godo [add] [del] [shw]")
}
