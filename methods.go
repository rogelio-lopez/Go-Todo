package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

/* Moodify List */
func (t *TodoItem) addTodoItem(args []string) {
	var timestamp time.Time = time.Now()
	var entryArr []string = args[2:]
	var entry string

	if args[2] == "-i" {
		entryArr = args[3:]
		entry = "* "
	}
	for _, e := range entryArr {
		entry += e + " "
	}
	entry = entry[0 : len(entry)-1]

	t.Entry = entry
	t.Date = timestamp.Format(time.Stamp)
}

func (t *Todo) delTodoItem(args []string) {
	if len(args) < 3 {
		t.List = t.List[1:len(t.List)]
	} else {
		n, err := strconv.Atoi(args[2])
		if err != nil {
			log.Fatalf("Strconv Error: %s\n", err)
		}

		var l int = len(t.List)
		if n == len(t.List) {
			t.List = t.List[0 : len(t.List)-1]
		} else if n == 1 {
			t.List = t.List[1:len(t.List)]
		} else {
			firstHalf, secondHalf := t.List[0:n-1], t.List[n:l]
			t.List = append(firstHalf, secondHalf...)
		}
	}
}

func (t *Todo) addUser(args []string) {
	t.UserName = ""
	for i := 2; i < len(args); i++ {
		t.UserName += args[i] + " "
	}
	t.UserName = t.UserName[0 : len(t.UserName)-1]
}

func (t *Todo) orderBy() {
	for i := 0; i < len(t.List); i++ {
		//Something that keeps the position of top pointer / the important ones
		entry := t.List[0].Entry
		if entry[0:1] != "*" {
		}
	}
}

/* Prints */
func (t *Todo) printList() {
	longestEntry, longestDate := 0, 0
	for _, t := range t.List {
		if len(t.Entry) > longestEntry {
			longestEntry = len(t.Entry)
		}
		if len(t.Date) > longestDate {
			longestDate = len(t.Date)
		}
	}

	if len(t.List) < 1 {
		fmt.Println("Your todo list is empty ¯\\_(ツ)_/¯")
		return
	}

	if t.UserName != "" {
		fmt.Printf("%s's Todo\n", t.UserName)
	} else {
		fmt.Println("My Todo")
	}
	fmt.Println(strings.Repeat("-", 11+longestEntry+longestDate))
	for i, v := range t.List {
		fmt.Printf("| %-*d | %-*s | %-*s |\n", 1, i+1, longestEntry, v.Entry, longestDate, v.Date)
	}
	fmt.Println(strings.Repeat("-", 11+longestEntry+longestDate))
}

func printInstructions() {
	fmt.Println("Go-do Usage: ./godo [add] [del] [shw]")
}

// Files
func (t *Todo) getFileJSON() {
	var info Todo

	fileData, err := os.ReadFile("todo-list.json")
	if err != nil {
		log.Fatalf("ReadFile Error: %s\n", err)
	}

	unmarshErr := json.Unmarshal(fileData, &info)
	if unmarshErr != nil {
		log.Fatalf("Unmarshal Error: %s\n", unmarshErr)
	}

	t.UserName = info.UserName
	t.List = info.List
}

func (t *Todo) pushFileJSON() {
	listAsByte, err := json.MarshalIndent(t, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("todo-list.json", listAsByte, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
