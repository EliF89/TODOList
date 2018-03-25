package model

import "fmt"

type Task struct {
	ToDoList string
	Title string 
	Done  bool   
}

func AddTask(todoListName string, taskTitle string) (*Task, error) {
	if taskTitle == "" || todoListName == "" {
		return nil, fmt.Errorf("empty mandatory parameters")
	}
	
	if task, _ := GetTask(todoListName, taskTitle); task != nil {
		return nil, fmt.Errorf("task already present")
	}

	list, err := GetToDoList(todoListName)

	if err != nil{
		return nil, err
	}

	task := &Task {	ToDoList: todoListName,
					Title: 	taskTitle,
					Done:	false} 

	list.Tasks = append(list.Tasks, cloneTask(task))
	list.TaskNumber = list.TaskNumber + 1 
	return task, nil
}

func GetTask(todoListName string, taskTitle string) (*Task, error) {
	if taskTitle == "" || todoListName == "" {
		return nil, fmt.Errorf("empty mandatory parameters")
	}
	
	list, err := GetToDoList(todoListName)

	if err != nil{
		return nil, err
	}

	for _, t := range list.Tasks {
		if t.Title == taskTitle {
			return t, nil
		}
	}
	return nil, fmt.Errorf("Task not found")
}

func UpdateTask(todoListName string, taskTitle string, newTitle string, done bool) (*Task, error) {
	if taskTitle == "" || todoListName == "" {
		return nil, fmt.Errorf("empty mandatory parameters")
	}
	
	list, err := GetToDoList(todoListName)

	if err != nil{
		return nil, err
	}

	for _, t := range list.Tasks {
		if t.Title == taskTitle {
			t.Done = done
			t.Title = newTitle
			return t, nil
		}
	}
	return nil, fmt.Errorf("Task not found")
}

func RemoveTask(todoListName string, taskTitle string) (*Task, error) {
	if taskTitle == "" || todoListName == "" {
		return nil, fmt.Errorf("empty mandatory parameters")
	}
	
	list, err := GetToDoList(todoListName)

	if err != nil{
		return nil, err
	}

	for i, t := range list.Tasks {
		if t.Title == taskTitle {
			list.Tasks = append(list.Tasks[:i], list.Tasks[i+1:]...)
			list.TaskNumber = list.TaskNumber - 1 
			return t, nil
		}
	}
	return nil, fmt.Errorf("Task not found")
}

// cloneTask creates and returns a deep copy of the given Task.
func cloneTask(t *Task) *Task {
	c := *t
	return &c
}
