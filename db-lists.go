package main

import (
	"fmt"
)

// Checkout Controls
// add / del / switch lists
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

// Delete List
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
		fmt.Println("Ya need to give me the name of a list to delete")
	}

	if getSelectedListIndex(db.ListArr) < 0 && len(db.ListArr) > 0 {
		db.ListArr[0].SelectedList = true
	}
	db.refreshIndexes()
	db.pushFileJSON()
}

// Add List
func (db *DB) addList(args []string) {
	var listName string = argText(args[2:])

	// Create new list if name != a list name
	if db.changeThisList(listName) < 0 {
		db.allSelectedToFalse()
		var newList List = List{
			Index:        uint(len(db.ListArr)),
			ListName:     listName,
			LastModified: timestamp(),
			List:         []Entry{},
			SelectedList: true,
		}
		db.ListArr = append(db.ListArr, newList)
		db.pushFileJSON()
	}
}

// Switch to different list if name exists
func (db *DB) changeThisList(listName string) int {
	var found int = -1
	for i := 0; i < len(db.ListArr); i++ {
		if db.ListArr[i].ListName == listName {
			db.allSelectedToFalse()
			db.ListArr[i].SelectedList = true
			db.pushFileJSON()
			found = i
		} else {
			db.ListArr[i].SelectedList = false
		}
	}
	return found
}

// Update all "selectedList" to false
func (db *DB) allSelectedToFalse() {
	for i := 0; i < len(db.ListArr); i++ {
		db.ListArr[i].SelectedList = false
	}
}
