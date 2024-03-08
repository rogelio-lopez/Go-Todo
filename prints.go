package main

import (
	"fmt"
)

func printInstructions() {
	fmt.Println("Go-do Usage:")
	fmt.Println("./godo [add] [del] [shw]")
}

func displayList(todoList []Todo) {

	if len(todoList) < 1 {
		fmt.Println("Your todo list is empty ¯\\_(ツ)_/¯")
		return
	}

	var largestItem int = 0
	var largestDate int = 0

	for _, t := range todoList {
		if len(t.Item) > largestItem {
			largestItem = len(t.Item)
		}
		if len(t.Date) > largestDate {
			largestDate = len(t.Date)
		}
	}
	for i, v := range todoList {
		fmt.Printf("%-*d | %-*s | %-*d | %-*s |\n", 1, i+1, largestItem, v.Item, 1, v.Priority, largestDate, v.Date)
	}
}
