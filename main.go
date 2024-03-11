package main

import (
	"os"
)

type Todo struct {
	UserName string     `json:"username"`
	List     []TodoItem `json:"list"`
}
type TodoItem struct {
	Entry string `json:"todoItem"`
	Date  string `json:"dateAdded"`
}

func main() {
	var todo Todo
	var todoItem TodoItem

	todo.getFileJSON()

	args := os.Args
	if len(args) > 1 {
		switch args[1] {

		case "shw":
			todo.printList()

		case "add":
			todoItem.addTodoItem(args)
			todo.List = append(todo.List, todoItem)
			todo.pushFileJSON()

		case "del":
			todo.delTodoItem(args)
			todo.pushFileJSON()

		case "usr":
			todo.addUser(args)
			todo.pushFileJSON()

		case "ordr":
			todo.orderBy()
			todo.pushFileJSON()

		default:
			printInstructions()
		}
	} else {
		printInstructions()
		todo.printList()
	}
}
