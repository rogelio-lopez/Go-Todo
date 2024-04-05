package main

type DB struct {
	Db_name        string      `json:"Db_name"`
	DB_CurrentList CurrentList `json:"DB_CurrentList"`
	Db_lists       []List      `json:"Db_lists"`
}
type List struct {
	Index         uint    `json:"index"`
	List_name     string  `json:"List_name"`
	Last_modified string  `json:"Last_modified"`
	List          []Entry `json:"List"`
}
type Entry struct {
	Index     uint   `json:"index"`
	Entry     string `json:"todoItem"`
	Date      string `json:"dateAdded"`
	Important bool   `json:"important"`
}

type CurrentList struct {
	Index uint   `json:"index"`
	Name  string `json:"name"`
}
