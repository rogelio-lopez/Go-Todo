package main

import (
	"os"
	"time"
)

type TodoItem struct {
	Entry    string `json:"todoItem"`
	Date     string `json:"dateAdded"`
	Priority int    `json:"priority"`
}

type Todo struct {
	UserName string
	List     []TodoItem
}

func main() {
	var item TodoItem
	var timestamp time.Time = time.Now()
	var todo Todo

	todo.UserName = ""
	todo.List = getFileJSON()

	args := os.Args
	if len(args) > 1 {
		switch args[1] {

		case "shw":
			printList(todo.List, todo.UserName)

		case "add":
			item.addTodoItem(args, timestamp)
			todo.List = append(todo.List, item)
			pushFileJSON(todo.List)

		case "del":
			todo.delTodoItem(args)
			pushFileJSON(todo.List)

		case "usr":
			todo.addUser(args)
		default:
			printInstructions()
		}
	} else {
		//havent done
		printList(todo.List, todo.UserName)
	}
}
