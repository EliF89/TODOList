package structutils

import "fmt"

var data map[string]*ToDoList

// ToDoList manages a list of tasks in memory
type ToDoList struct {
	Name string			// ToDo list name
	TaskNumber int		// Number of task in the list
	Tasks  []*Task		// list of Task associated
	lastID int64		// last inserted Task ID	
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
	newToDoList := &ToDoList{}
	newToDoList.Name = name
	//newToDoList.TaskNumber = 0

	data[name] = newToDoList
	return data[name], nil
}

func GetToDoList(name string) (*ToDoList, error) {
	if name == "" || data[name] == nil {
		return nil, fmt.Errorf("ToDo list not found")
	}
	return data[name], nil
}

func GetAllToDoList() ([]ToDoList) {
	allToDoList := []ToDoList{}
    for _, value := range data {
        allToDoList = append(allToDoList, *value)
    }
	return allToDoList
}

func DeleteToDoList(name string) ( error)  {
	if name == "" || data[name] == nil {
		return  fmt.Errorf("ToDo list not found, list not deleted")
	}
	delete(data, name)
	return nil
}

func GetToDoListsInfo() (string){
	keys := make([]string, 0, len(data))
    for k := range data {
        keys = append(keys, fmt.Sprintf(" '%s' ", k))
	}
	return fmt.Sprintf("#ToDoList = %d. Names = %v",  len(data), keys)
}


// Add the given Task in the ToDoList.
func (m *ToDoList) AddTask(taskTitle string) (*Task, error) {
	if taskTitle == "" {
		return nil, fmt.Errorf("empty title")
	}
	m.lastID++
	task := &Task {	ID: 	m.lastID,
					Title: 	taskTitle,
					Done:	false} 

	m.Tasks = append(m.Tasks, cloneTask(task))
	m.TaskNumber = len(m.Tasks)
	return task, nil
}

// Add the given Task in the ToDoList.
func (m *ToDoList) UpdateTask(taskTitle string, done bool) (*Task, error) {
	if taskTitle == "" {
		return nil, fmt.Errorf("empty title")
	}

	for i, t := range m.Tasks {
		if t.Title == taskTitle {
			t.Done = done
			m.Tasks[i] = cloneTask(t)
			return m.Tasks[i], nil
		}
	}
	return nil, fmt.Errorf("unknown task")
}

// Remove the given Task from the ToDo list
func (m *ToDoList) RemoveTask(taskName string) (*Task, error) {

	for i, t := range m.Tasks {
		if t.Title == taskName {
			m.Tasks = append(m.Tasks[:i], m.Tasks[i+1:]...)
			m.TaskNumber = len(m.Tasks)
			return t, nil
		}
	}
	return nil, fmt.Errorf("unknown task")
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
func (m *ToDoList) GetTask(title string) (*Task, bool) {
	for _, t := range m.Tasks {
		if t.Title == title {
			return t, true
		}
	}
	return nil, false
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