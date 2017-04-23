package main

import "errors"

var dal *Dal;

//TodoDal interface wraps 
type TodoDal interface{
	GetTodoItems(user string) (TodoList, error)
	CreateTodoList(user string, list TodoList) (error)
	AddItem(user string, item TodoItem) (error)
	//RemoveItem(user string, item TodoItem) (error)
	//GetItem(user string, item TodoItem) (TodoItem, error)
}


func GetDal() (TodoDal){
	if dal == nil {
		dal = &Dal {}
		dal.lists = make(map[string]TodoList)
	}
	return dal
}


//Dal - the main Data access layer
type Dal struct {
	lists map[string]TodoList
}

func (dal *Dal) AddItem(user string, item TodoItem) (error) {

	list, ok := dal.lists[user]
	if ok {
		_ = append(list.items, item)

		return nil
	}else{
		return errors.New("TodoList for user not found")
	}

}

func (dal *Dal) CreateTodoList(user string, list TodoList) error {

	//check if todolist exists in memory store
	_, ok := dal.lists[user]
	if !ok {
		dal.lists[user] = list
		return nil
	}else{
		return errors.New("TodoList already found in list")
	}

}

//GetTodoItems returns the current TODO list
func (dal *Dal) GetTodoItems(user string) (TodoList, error) {

	list := TodoList{}

	return list, nil

}