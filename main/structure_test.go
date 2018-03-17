package main

import "testing"

func newTaskOrFatal(t *testing.T, title string) *Task {
	task, err := CreateTask(title)
	if err != nil {
		t.Fatalf("new task: %v", err)
	}
	return task
}

func TestNewTask(t *testing.T) {
	title := "learn Go"
	task := newTaskOrFatal(t, title)
	if task.Title != title {
		t.Errorf("expected title %q, got %q", title, task.Title)
	}
	if task.Done {
		t.Errorf("new task is done")
	}
}

func TestNewTaskEmptyTitle(t *testing.T) {
	_, err := CreateTask("")
	if err == nil {
		t.Errorf("expected 'empty title' error, got nil")
	}
}

func TestCreateToDoList(t *testing.T) {
	_ , err := CreateToDoList("")
	if err == nil {
		t.Errorf("expected 'empty ToDo list name' error, got nil")
	}
}

func TestCreateToDoListAndRetrieve(t *testing.T) {
	CreateToDoList("TODO List title")
	
	mRet , _ := GetToDoList("TODO List title")

	if (mRet == nil){
		t.Errorf("expected a ToDoList to be retrieved, got nil")
	}
}

func TestRetrieveInvalidToDoList(t *testing.T) {
	CreateToDoList("TODO List title")
	
	_ , err := GetToDoList("Invalid title")

	if err == nil {
		t.Errorf("expected 'ToDo list not found' error, got nil")
	}
}

func TestAddTaskAndRetrieve(t *testing.T) {
	task := newTaskOrFatal(t, "learn Go")
	
	m, _ := CreateToDoList("TODO List title")
	m.Add(task)

	all := m.GetAll()
	if len(all) != 1 {
		t.Errorf("expected 1 task, got %v", len(all))
	}
	if *all[0] != *task {
		t.Errorf("expected %v, got %v", task, all[0])
	}
}

func TestAddAndRetrieveTwoTasks(t *testing.T) {
	learnGo := newTaskOrFatal(t, "learn Go")
	learnTDD := newTaskOrFatal(t, "learn TDD")

	m, _ := CreateToDoList("TODO List title")
	m.Add(learnGo)
	m.Add(learnTDD)

	all := m.GetAll()
	if len(all) != 2 {
		t.Errorf("expected 2 tasks, got %v", len(all))
	}
	if *all[0] != *learnGo && *all[1] != *learnGo {
		t.Errorf("missing task: %v", learnGo)
	}
	if *all[0] != *learnTDD && *all[1] != *learnTDD {
		t.Errorf("missing task: %v", learnTDD)
	}
}

func TestAddModifyAndRetrieve(t *testing.T) {
	task := newTaskOrFatal(t, "learn Go")
	m, _ := CreateToDoList("TODO List title")
	m.Add(task)

	task.Done = true
	if m.GetAll()[0].Done {
		t.Errorf("saved task wasn't done")
	}
}

func TestSaveTwiceAndRetrieve(t *testing.T) {
	task := newTaskOrFatal(t, "learn Go")
	m, _ := CreateToDoList("TODO List title")
	m.Add(task)
	m.Add(task)

	all := m.GetAll()
	if len(all) != 1 {
		t.Errorf("expected 1 task, got %v", len(all))
	}
	if *all[0] != *task {
		t.Errorf("expected task %v, got %v", task, all[0])
	}
}

func TestSaveAndFind(t *testing.T) {
	task := newTaskOrFatal(t, "learn Go")
	m, _ := CreateToDoList("TODO List title")
	m.Add(task)

	nt, ok := m.GetTask(task.ID)
	if !ok {
		t.Errorf("didn't find task")
	}
	if *task != *nt {
		t.Errorf("expected %v, got %v", task, nt)
	}
}