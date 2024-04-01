package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

/* Moodify List */
func (t *Entry) add(args []string) {
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
	t.Date = timestamp()
}

func (t *List) delTodoItem(args []string) {
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

func (t *List) orderBy() {
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
