package main

import (
	"fmt"
	"strings"
)

func printInstructions() {
	fmt.Println("Go-do Usage:")
	fmt.Println("./godo [add] [del] [shw]")
}

func printList(todoList []TodoItem, username string) {

	if len(todoList) < 1 {
		fmt.Println("Your todo list is empty ¯\\_(ツ)_/¯")
		return
	}

	longestEntry, longestDate := 0, 0
	for _, t := range todoList {
		if len(t.Entry) > longestEntry {
			longestEntry = len(t.Entry)
		}
		if len(t.Date) > longestDate {
			longestDate = len(t.Date)
		}
	}

	if username != "" {
		fmt.Printf("%s Todo", username)
	} else {
		fmt.Println("My Todo")
	}
	fmt.Println(strings.Repeat("-", 15+longestEntry+longestDate))
	for i, v := range todoList {
		fmt.Printf("| %-*d | %-*s | %-*d | %-*s |\n", 1, i+1, longestEntry, v.Entry, 1, v.Priority, longestDate, v.Date)
	}
	fmt.Println(strings.Repeat("-", 15+longestEntry+longestDate))
}
