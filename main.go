package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
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

	args := os.Args
	if len(args) > 1 {
		switch args[1] {

		case "shw":
			displayList(todoList)

		case "add":
			todo.addTodoItem(args, timestamp)
			todoList = append(todoList, todo)
			pushFileJSON(todoList)

		case "del":
			fmt.Println(args[2])

		default:
			printInstructions()
		}
	} else {
		//havent done
		displayList(todoList)
	}
}

func (t *Todo) addTodoItem(args []string, timestamp time.Time) {
	for i := 1; i < len(args); i++ {
		t.Item += args[i] + " "
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
