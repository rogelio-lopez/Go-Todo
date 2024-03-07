package main

import (
	"fmt"
)

func printInstructions() {
	fmt.Println("Go-do Usage:")
	fmt.Println("./godo [add] [del] [shw]")
}

func displayList(list []Todo) {
	for i, v := range list {
		fmt.Printf("%-*d | %-*s | %-*d | %-*s |\n", 1, i, 30, v.Item, 1, v.Priority, 8, v.Date)
	}
}
