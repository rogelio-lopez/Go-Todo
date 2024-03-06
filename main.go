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
	var todo Todo
	var timestamp time.Time = time.Now()
	var todoList []Todo = getFileJSON()

	//Print commands for executable
	printInstructions()

	args := os.Args
	if len(args) > 1 {
		switch args[1] {
		case "shw":
			displayList()
		case "add":
			fmt.Println(args[2])
		case "del":
			fmt.Println(args[2])
		}
	} else {
		//use this to be the default
		//which would to just display the list
		//displayList()
		fmt.Println("Enter a todo list item:")

		reader := bufio.NewReader(os.Stdin)
		strInput, err := reader.ReadString('\n')
		if err != nil {
			log.Fatalf("Reader error: %s\n", err)
		}

		todo.Item = strInput
		todo.Date = timestamp.Format(time.Stamp)
		todo.Priority = 69

		todoList = append(todoList, todo)

		pushFileJSON(todoList)

		for _, t := range todoList {
			fmt.Println(t.Item)
		}
	}
}

func getFileJSON() []Todo {
	var list []Todo

	fileData, err := os.ReadFile("todo-list.json")
	if err != nil {
		log.Fatalf("ReadFile Error: %s\n", err)
	}

	unmarshErr := json.Unmarshal(fileData, &list)
	if unmarshErr != nil {
		log.Fatalf("Unmarshal Error: %s\n", unmarshErr)
	}

	return list
}

func pushFileJSON(list []Todo) {
	listAsByte, err := json.MarshalIndent(list, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("todo-list.json", listAsByte, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
