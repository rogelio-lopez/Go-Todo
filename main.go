package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"
)

type Todo struct {
	Item     string `json:"todoItem"`
	Date     string `json:"dateAdded"`
	Priority int    `json:"priority"`
}

func main() {
	data, err := os.ReadFile("todo-list.json")
	if err != nil {
		fmt.Println("Could not read content from file")
		return
	}

	var todoInfo []Todo
	unmarshErr := json.Unmarshal(data, &todoInfo)
	if unmarshErr != nil {
		log.Fatal(unmarshErr)
	}
	var newItem Todo
	var t time.Time = time.Now()

	fmt.Println("Enter a todo list item:")

	reader := bufio.NewReader(os.Stdin)
	strInput, err := reader.ReadString('\n')
	if err != nil {
		fmt.Printf("Reader error: %s", err)
	}

	newItem.Item = strInput
	newItem.Date = t.Format(time.Stamp)
	newItem.Priority = 69

	todoInfo = append(todoInfo, newItem)

	byteInfo, err := json.MarshalIndent(todoInfo, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("todo-list.json", byteInfo, 0644)
	if err != nil {
		log.Fatal(err)
	}

	for _, t := range todoInfo {
		fmt.Println(t.Item)
	}
}
