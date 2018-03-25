package model

import "testing"

/*******************************
	CREATE Task
*******************************/
func TestCreateTask_noListName_Error(t *testing.T) {
	_, err := AddTask("", "")
	if err == nil {
		t.Errorf("expected empty ToDo list name or task title error, got nil")
	}
}

func TestCreateTask_noTaskTitle_Error(t *testing.T) {
	_, err := AddTask("List1", "")
	if err == nil {
		t.Errorf("expected empty ToDo list name or task title error, got nil")
	}
}


func TestCreateTask_invalidList_Error(t *testing.T) {
	_, err := AddTask("Invalid", "")
	if err == nil {
		t.Errorf("expected invalid ToDo list error, got nil")
	}
}

func TestCreateTask_ListTask1_task1_ok(t *testing.T) {	
	
	list, err := CreateToDoList("ListTask1")
	if err != nil {
		t.Errorf("no error expected, got %v", err)
	}
	
	if err != nil {
		t.Errorf("no error expected, got %v", err)
	}
	if (list == nil){
		t.Errorf("expected ToDoList List2 to be retrieved, got nil")
	}
	if len(list.Tasks) != 0 {		
		t.Errorf("expected 0 tasks in ToDoList ListTask1, got %d", len(list.Tasks))
	}

	task, err := AddTask("ListTask1", "Task1")
	if err != nil {
		t.Errorf("no error expected, got %v", err)
	}
	if (task == nil){
		t.Errorf("expected a Task to be retrieved, got nil")
	}
	if (task.Title != "Task1"){
		t.Errorf("expected task title eq Task1, got %s", task.Title)
	}
	if (task.ToDoList != "ListTask1"){
		t.Errorf("expected Task1 associated to ListTask1, got %s", task.Title)
	}	
	if (task.Done != false){
		t.Errorf("expected Task1 done = false, got %t", task.Done)
	}

	list, err = GetToDoList("ListTask1")
	if err != nil {
		t.Errorf("no error expected, got %v", err)
	}
	if (list == nil){
		t.Errorf("expected ToDoList List2 to be retrieved, got nil")
	}
	if len(list.Tasks) != 1 {		
		t.Errorf("expected 1 tasks in ToDoList ListTask1, got %d", len(list.Tasks))
	}

}

func TestCreateTask_alreadyExisting_error(t *testing.T) {
	_, err := AddTask("ListTask1", "Task1")
	if err == nil {
		t.Errorf("expected Task alredy present error, got nil ")
	}	
}

func TestCreateTask_ListTask1_task2_ok(t *testing.T) {
	task, err := AddTask("ListTask1", "Task2")
	if err != nil {
		t.Errorf("no error expected, got %v", err)
	}
	if (task == nil){
		t.Errorf("expected a Task to be retrieved, got nil")
	}
	if (task.Title != "Task2"){
		t.Errorf("expected task title eq Task2, got %s", task.Title)
	}
	if (task.ToDoList != "ListTask1"){
		t.Errorf("expected Task2 associated to ListTask1, got %s", task.Title)
	}	
	if (task.Done != false){
		t.Errorf("expected Task2 done = false, got %t", task.Done)
	}

	
	list, err := GetToDoList("ListTask1")
	if err != nil {
		t.Errorf("no error expected, got %v", err)
	}
	if (list == nil){
		t.Errorf("expected ToDoList List2 to be retrieved, got nil")
	}
	if len(list.Tasks) != 2 {		
		t.Errorf("expected 2 tasks in ToDoList ListTask1, got %d", len(list.Tasks))
	}

}



/*******************************
	GET Task
*******************************/
func TestGetTaks_invalidListName_error(t *testing.T) {
	_, err := GetTask("invalid", "Task1")
	if err == nil {
		t.Errorf("Expected error list not found, got nil")
	}
}


func TestGetTask_nullName_error(t *testing.T) {
	_, err := GetTask("", "Task1")
	if err == nil {
		t.Errorf("Expected error null list name, got nil")
	}
}


func TestGetTask_nullTask_error(t *testing.T) {
	_, err := GetTask("ListTask1", "")
	if err == nil {
		t.Errorf("Expected error null task title, got nil")
	}
}


func TestGetTask_invalidTask_error(t *testing.T) {
	_, err := GetTask("ListTask1", "invalid")
	if err == nil {
		t.Errorf("Expected error task not found, got nil")
	}
}

func TestGetTask_ok(t *testing.T) {
	task, err := GetTask("ListTask1", "Task1")
	if err != nil {
		t.Errorf("no error expected, got %v", err)
	}
	if (task == nil){
		t.Errorf("expected Task1 in ListTask1 to be retrieved, got nil")
	}
	if (task.Title != "Task1"){
		t.Errorf("expected task title eq Task1, got %s", task.Title)
	}
	if (task.ToDoList != "ListTask1"){
		t.Errorf("expected Task1 associated to ListTask1, got %s", task.Title)
	}	
	if (task.Done != false){
		t.Errorf("expected Task1 done = false, got %t", task.Done)
	}
}


/*******************************
	UPDATE Task
*******************************/
func TestUpdateTaks_invalidListName_error(t *testing.T) {
	_, err := UpdateTask("invalid", "Task1", "new name", false)
	if err == nil {
		t.Errorf("Expected error list not found, got nil")
	}
}


func TestUpdateTask_nullName_error(t *testing.T) {
	_, err := UpdateTask("", "Task1", "new name", false)
	if err == nil {
		t.Errorf("Expected error null list name, got nil")
	}
}


func TestUpdateTask_nullTask_error(t *testing.T) {
	_, err := UpdateTask("ListTask1", "", "new name", false)
	if err == nil {
		t.Errorf("Expected error null task title, got nil")
	}
}


func TestUpdateTask_invalidTask_error(t *testing.T) {
	_, err := UpdateTask("ListTask1", "invalid", "new name", false)
	if err == nil {
		t.Errorf("Expected error task not found, got nil")
	}
}

func TestUpdateTask_newName_ok(t *testing.T) {
	task, err := UpdateTask("ListTask1", "Task1", "Task1New", false)
	if err != nil {
		t.Errorf("no error expected, got %v", err)
	}
	if (task == nil){
		t.Errorf("expected Task1New in ListTask1 to be retrieved, got nil")
	}
	if (task.Title != "Task1New"){
		t.Errorf("expected task title eq Task1New, got %s", task.Title)
	}
	if (task.ToDoList != "ListTask1"){
		t.Errorf("expected Task1New associated to ListTask1, got %s", task.Title)
	}	
	if (task.Done != false){
		t.Errorf("expected Task1New done = false, got %t", task.Done)
	}
}


func TestUpdateTask_newNameAndStatus_ok(t *testing.T) {
	task, err := UpdateTask("ListTask1", "Task2", "Task2New", true)
	if err != nil {
		t.Errorf("no error expected, got %v", err)
	}
	if (task == nil){
		t.Errorf("expected Task2New in ListTask1 to be retrieved, got nil")
	}
	if (task.Title != "Task2New"){
		t.Errorf("expected task title eq Task2New, got %s", task.Title)
	}
	if (task.ToDoList != "ListTask1"){
		t.Errorf("expected Task2New associated to ListTask1, got %s", task.Title)
	}	
	if (task.Done != true){
		t.Errorf("expected Task2New done = true, got %t", task.Done)
	}
}

/*******************************
	DELETE Task
*******************************/
func TestRemoveTaks_invalidListName_error(t *testing.T) {
	_, err := RemoveTask("invalid", "Task1New")
	if err == nil {
		t.Errorf("Expected error list not found, got nil")
	}
}


func TestRemoveTaskTask_nullName_error(t *testing.T) {
	_, err := RemoveTask("", "Task1New")
	if err == nil {
		t.Errorf("Expected error null list name, got nil")
	}
}


func TestRemoveTask_nullTask_error(t *testing.T) {
	_, err := RemoveTask("ListTask1", "")
	if err == nil {
		t.Errorf("Expected error null task title, got nil")
	}
}


func TestRemoveTask_invalidTask_error(t *testing.T) {
	_, err := RemoveTask("ListTask1", "invalid")
	if err == nil {
		t.Errorf("Expected error task not found, got nil")
	}
}

func TestRemoveTask_newName_ok(t *testing.T) {
	task, err := RemoveTask("ListTask1", "Task1New")
	if err != nil {
		t.Errorf("no error expected, got %v", err)
	}
	if (task == nil){
		t.Errorf("expected Task1New in ListTask1 to be retrieved, got nil")
	}
	if (task.Title != "Task1New"){
		t.Errorf("expected task title eq Task1New, got %s", task.Title)
	}
	if (task.ToDoList != "ListTask1"){
		t.Errorf("expected Task1New associated to ListTask1, got %s", task.Title)
	}	
	if (task.Done != false){
		t.Errorf("expected Task1New done = false, got %t", task.Done)
	}

	
	list, err := GetToDoList("ListTask1")
	if err != nil {
		t.Errorf("no error expected, got %v", err)
	}
	if (list == nil){
		t.Errorf("expected ToDoList List2 to be retrieved, got nil")
	}
	if len(list.Tasks) != 1 {		
		t.Errorf("expected 1 tasks in ToDoList ListTask1, got %d", len(list.Tasks))
	}
}
