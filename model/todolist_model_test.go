package model

import "testing"

/*******************************
	CREATE ToDo list
*******************************/

func TestCreateToDoList_noName_Error(t *testing.T) {
	_ , err := CreateToDoList("")
	if err == nil {
		t.Errorf("expected empty ToDo list name error, got nil")
	}
}

func TestCreateToDoList_list1_ok(t *testing.T) {
	list, err := CreateToDoList("List1")
	if err != nil {
		t.Errorf("no error expected, got %v", err)
	}
	if (list == nil){
		t.Errorf("expected a ToDoList to be retrieved, got nil")
	}
	if (list.Name != "List1"){
		t.Errorf("expected ToDoList returned with name List1, got %s", list.Name)
	}
}

func TestCreateToDoList_alreadyExisting_error(t *testing.T) {
	_, err := CreateToDoList("List1")
	if err == nil {
		t.Errorf("expected ToDo alredy present error, got nil ")
	}	
}

func TestCreateToDoList_list2_ok(t *testing.T) {
	list, err := CreateToDoList("List2")
	if err != nil {
		t.Errorf("no error expected, got %v", err)
	}
	if (list == nil){
		t.Errorf("expected a ToDoList to be retrieved, got nil")
	}
	if (list.Name != "List2"){
		t.Errorf("expected ToDoList returned with name List2, got %s", list.Name)
	}
}

/*******************************
	GET ToDo list
*******************************/

func TestGetToDoList_invalidName_error(t *testing.T) {
	_, err := GetToDoList("invalid")
	if err == nil {
		t.Errorf("Expected error list not found, got nil")
	}
}


func TestGetToDoList_nullName_error(t *testing.T) {
	_, err := GetToDoList("")
	if err == nil {
		t.Errorf("Expected error invalid list name, got nil")
	}
}

func TestGetToDoList_ok(t *testing.T) {
	list, err := GetToDoList("List2")
	if err != nil {
		t.Errorf("no error expected, got %v", err)
	}
	if (list == nil){
		t.Errorf("expected ToDoList List2 to be retrieved, got nil")
	}
	if (list.Name != "List2"){
		t.Errorf("expected ToDoList returned with name List2, got %s ", list.Name)
	}
}



/*******************************
	UPDATE ToDo list
*******************************/

func TestUpdateToDoList_invalidName_error(t *testing.T) {
	_, err := UpdateToDoList("invalid", "")
	if err == nil {
		t.Errorf("Expected error list not found, got nil")
	}
}


func TestUpdateToDoList_nullName_error(t *testing.T) {
	_, err := UpdateToDoList("", "")
	if err == nil {
		t.Errorf("Expected error invalid list name, got nil")
	}
}

func TestUpdateToDoList_nullNewName_error(t *testing.T) {
	_, err := UpdateToDoList("List2", "")
	if err == nil {
		t.Errorf("Expected error invalid list name, got nil")
	}
}

func TestUpdateToDoList_ok(t *testing.T) {
	list, err := UpdateToDoList("List2", "List2New")
	if err != nil {
		t.Errorf("no error expected, got %v", err)
	}
	if (list == nil){
		t.Errorf("expected ToDoList List2New to be returned, got nil")
	}
	if (list.Name != "List2New"){
		t.Errorf("expected ToDoList returned with name List2New, got %s ", list.Name)
	}

}

/*******************************
	DELETE ToDo list
*******************************/

func TestDeleteToDoList_invalidName_error(t *testing.T) {
	_, err := DeleteToDoList("invalid")
	if err == nil {
		t.Errorf("Expected error list not found, got nil")
	}
}


func TestDeleteToDoList_nullName_error(t *testing.T) {
	_, err := DeleteToDoList("")
	if err == nil {
		t.Errorf("Expected error invalid list name, got nil")
	}
}

func TestDeleteToDoList_ok(t *testing.T) {
	list, err := DeleteToDoList("List2New")
	if err != nil {
		t.Errorf("no error expected, got %v", err)
	}
	if (list == nil){
		t.Errorf("expected ToDoList List2New to be returned, got nil")
	}
	if (list.Name != "List2New"){
		t.Errorf("expected ToDoList returned with name List2New, got %s ", list.Name)
	}

	_, err = GetToDoList("List2")
	if err == nil {
		t.Errorf("expected error ToDo list not found, got nil")
	}
}
