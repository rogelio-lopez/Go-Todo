package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func (t *TodoItem) addTodoItem(args []string, timestamp time.Time) {
	for i := 1; i < len(args); i++ {
		t.Entry += args[i] + " "
	}
	t.Date = timestamp.Format(time.Stamp)

	fmt.Println("Priority [1-5]")

	reader := bufio.NewReader(os.Stdin)
	priority, err := reader.ReadString('\n')
	if err != nil {
		log.Fatalf("Reader error: %s\n", err)
	}

	t.Priority, err = strconv.Atoi(priority[0:1])
	if err != nil {
		log.Fatalf("Strconv Error: %s\n", err)
	}
}

func (t *Todo) delTodoItem(args []string) {
	//if no number entered in args, delete first
	if len(args) < 3 {
		t.List = t.List[1:len(t.List)]
	} else {
		//Delete the entry number entered in args
		n, err := strconv.Atoi(args[2])
		if err != nil {
			log.Fatalf("Strconv Error: %s\n", err)
		}
		fmt.Println(n)
	}
}

func (t *Todo) addUser(args []string) {
	for i := 1; i < len(args); i++ {
		t.UserName += args[i] + " "
	}
}

// Files
func getFileJSON() []TodoItem {
	var list []TodoItem

	fileData, err := os.ReadFile("todo-list.json")
	if err != nil {
		log.Fatalf("ReadFile Error: %s\n", err)
	}

	if len(fileData) < 1 {
		return []TodoItem{}
	}

	unmarshErr := json.Unmarshal(fileData, &list)
	if unmarshErr != nil {
		log.Fatalf("Unmarshal Error: %s\n", unmarshErr)
	}

	return list
}

func pushFileJSON(list []TodoItem) {
	listAsByte, err := json.MarshalIndent(list, "", "\t")
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile("todo-list.json", listAsByte, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
