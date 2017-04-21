package main


//TodoDal interface wraps 
type TodoDal interface{
	GetTodoItems(user string) (TodoList, error)
	CreateTodoList(user string, list TodoList) (error)
	AddItem(user string, item TodoItem) (error)
	RemoveItem(user string, item TodoItem) (error)
	GetItem(user string, item TodoItem) (TodoItem, error)
}

//Dal - the main Data access layer
type Dal struct {
	lists map[string]TodoList
}


//GetTodoItems returns the current TODO list
func (dal *Dal) GetTodoItems(user string) (TodoList, error) {

	list := TodoList{}

	return list, nil

}