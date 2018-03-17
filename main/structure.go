package main

import "fmt"

var data map[string]*ToDoList

// ToDoList manages a list of tasks in memory
type ToDoList struct {
	Tasks  []*Task		// list of Task associated
	LastID int64		// last inserted Task ID
}


type Task struct {
	ID    int64  // Unique identifier
	Title string // Task description
	Done  bool   // Task status: done/not done
}


// returns an empty ToDoList with the name provided.
func CreateToDoList(name string) (*ToDoList, error) {
	if name == "" {
		return nil, fmt.Errorf("empty ToDo list name")
	}
	if data == nil {
		data = make(map[string]*ToDoList, 100)
	}

	data[name] = &ToDoList{}
	return data[name], nil
}

func GetToDoList(name string) (*ToDoList, error) {
	if name == "" || data[name] == nil {
		return nil, fmt.Errorf("ToDo list not found")
	}
	return data[name], nil
}

// add a new task
// constraint: the title cannot be empty
func CreateTask(title string) (*Task, error) {
	if title == "" {
		return nil, fmt.Errorf("empty title")
	}
	return &Task{	ID:0, 
					Title: title,
					Done: false} , nil
}

// Add the given Task in the ToDoList.
func (m *ToDoList) Add(task *Task) error {
	if task.ID == 0 {
		m.LastID++
		task.ID = m.LastID
		m.Tasks = append(m.Tasks, cloneTask(task))
		return nil
	}

	for i, t := range m.Tasks {
		if t.ID == task.ID {
			m.Tasks[i] = cloneTask(task)
			return nil
		}
	}
	return fmt.Errorf("unknown task")
}

// cloneTask creates and returns a deep copy of the given Task.
func cloneTask(t *Task) *Task {
	c := *t
	return &c
}

// Returns the list of all the Tasks in the ToDoList.
func (m *ToDoList) GetAll() []*Task {
	return m.Tasks
}

/* 
	Search for the Task with the given id.
	Returns:
	- the Task if found,  null otherwise
	- a bool indicating if the Task was found
*/
func (m *ToDoList) GetTask(ID int64) (*Task, bool) {
	for _, t := range m.Tasks {
		if t.ID == ID {
			return t, true
		}
	}
	return nil, false
}
