package main

import (
	"fmt"
)

/*--- Checkout COntrols ---*/
func (db *DB) checkout(args []string) {
	if len(args) >= 3 {
		switch args[2] {
		case "-d":
			db.delList(args)
		default:
			db.addList(args)
		}
	} else {
		shwDB(*db)
	}
}

func (db *DB) delList(args []string) {
	if len(args) >= 4 {
		var listName = argText(args[3:])

		for i, l := range db.ListArr {
			if listName == l.ListName {

				if len(db.ListArr) < 2 {
					db.ListArr = []List{}

				} else if i == 0 {
					db.ListArr = db.ListArr[1:]

				} else if i == len(db.ListArr) {
					db.ListArr = db.ListArr[0 : len(db.ListArr)-1]

				} else {
					first := db.ListArr[0:i]
					second := db.ListArr[i+1:]
					db.ListArr = append(first, second...)
				}

			}
		}
	} else {
		fmt.Println("Ya need to give me the name of the list to delete")
	}

	if getSelectedListIndex(db.ListArr) < 0 && len(db.ListArr) > 0 {
		db.ListArr[0].SelectedList = true
	}
	db.refreshIndexes()
	db.pushFileJSON()
}

func (db *DB) addList(args []string) {
	var listName string = argText(args[2:])

	// Create new list if name != a list name
	if db.changeThisList(listName) < 0 {
		var newList List
		newList.Index = uint(len(db.ListArr))
		newList.ListName = listName
		newList.LastModified = timestamp()
		newList.List = []Entry{}
		newList.SelectedList = true

		db.ListArr = append(db.ListArr, newList)
		db.pushFileJSON()
	}
}

func (db *DB) changeThisList(listName string) int {
	var found int = -1
	for i, l := range db.ListArr {
		if l.ListName == listName {
			l.SelectedList = true
			found = i
			db.pushFileJSON()

		} else {
			l.SelectedList = false
		}
	}
	return found
}
