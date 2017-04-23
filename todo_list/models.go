package main

//TodoItem is the main structure for todo list items
type TodoItem struct {
	title string
	dateCreated int64
	notes string
}

//TodoList is the main list of TODO's
type TodoList struct {
	items []TodoItem
}