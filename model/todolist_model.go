package model

import "fmt"

var data map[string]*ToDoList

// ToDoList manages a list of tasks in memory
type ToDoList struct {
	Name string			
	Tasks  []*Task	
	TaskNumber int	
}


func CreateToDoList(name string) (*ToDoList, error) {
	if name == "" {
		return nil, fmt.Errorf("empty ToDo list name")
	}
	
	if list, _ := GetToDoList(name); list != nil {
		return nil, fmt.Errorf("list already present")
	}
	if data == nil {
		data = make(map[string]*ToDoList, 100)
	}
	newToDoList := &ToDoList{}
	newToDoList.Name = name
	newToDoList.TaskNumber = 0

	data[name] = newToDoList
	return data[name], nil
}

func  GetToDoList(name string) (*ToDoList, error) {
	if name == "" || data == nil || data[name] == nil {
		return nil, fmt.Errorf("ToDo list not found")
	}
	return data[name], nil
}

func GetAllToDoList() ([]ToDoList, error) {
	allToDoList := []ToDoList{}
	if data != nil {
		for _, value := range data {
			allToDoList = append(allToDoList, *value)
		}
	}
	return allToDoList, nil
}

func DeleteToDoList(name string) (*ToDoList, error) {
	if name == "" || data == nil || data[name] == nil {
		return  nil, fmt.Errorf("ToDo list not found, list not deleted")
	}
	list := data[name]
	delete(data, name)
	return list, nil
}

func UpdateToDoList(name string, newName string)(*ToDoList, error) {
	if name == "" || newName == "" || data == nil || data[name] == nil {
		return  nil, fmt.Errorf("ToDo list not found, list not deleted")
	}
	list := data[name]
	list.Name = newName
	delete(data, name)
	data[newName] = list
	return list, nil
}