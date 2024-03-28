package main

import (
	"bufio"
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
	var entry string

	if len(args) > 2 {
		switch args[2] {
		case "-i":
			entry = "* "
		}
	}

	fmt.Printf("~ ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	err := scanner.Err()
	if err != nil {
		log.Fatalf("Scanner Error: %s\n", err)
	}

	t.Entry = entry + scanner.Text()
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
	var listLen int = len(t.List)
	var index int = 0

	i := 0
	for i < listLen {
		entry := t.List[index].Entry

		if i == 0 && entry[0:1] != "*" {
			el := t.List[0]
			t.List = t.List[1:]
			t.List = append(t.List, el)

		} else if entry[0:1] != "*" {
			el := t.List[index]
			newList := append(t.List[0:index], t.List[index+1:]...)
			t.List = append(newList, el)
		} else if entry[0:1] == "*" {
			index++
		}
		i++
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
