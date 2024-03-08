package main

import (
	"encoding/json"
	"log"
	"os"
)

func getFileJSON() []Todo {
	var list []Todo

	fileData, err := os.ReadFile("todo-list.json")
	if err != nil {
		log.Fatalf("ReadFile Error: %s\n", err)
	}

	if len(fileData) < 1 {
		return []Todo{}
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
