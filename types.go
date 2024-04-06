package main

type DB struct {
	DbName   string   `json:"DbName"`
	ThisList ThisList `json:"ThisList"`
	ListArr  []List   `json:"ListArr"`
}
type List struct {
	Index        uint    `json:"Index"`
	ListName     string  `json:"ListName"`
	LastModified string  `json:"LastModified"`
	List         []Entry `json:"List"`
}
type Entry struct {
	Index     uint   `json:"index"`
	Entry     string `json:"todoItem"`
	Date      string `json:"dateAdded"`
	Important bool   `json:"important"`
}
type ThisList struct {
	Index uint   `json:"index"`
	Name  string `json:"name"`
}
