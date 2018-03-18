package handler

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/efreddo/todolist/logutils"
	"github.com/efreddo/todolist/structutils"
	"github.com/julienschmidt/httprouter"
)


/* 
	request type: POST
	url: /create/ {"Name": "New ToDo list"}
	The request body must contain a JSON object with a Name field

	Examples:

	   req: POST /create/ {"Name": ""}
	   res: 400 empty name

	   req: POST /create/ {"name": "New ToDo List"}
	   res: 200
*/
func CreateToDoList(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	req := struct{ Name string }{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Name == "" {
		logutils.Error.Println("CreateToDoList:: Bad request received", err)
		http.Error(w, "Missing ToDo list name", http.StatusBadRequest)
		return
	}

	_, err := structutils.CreateToDoList(req.Name)
	if err != nil {				
		logutils.Error.Println(fmt.Sprintf(
			"CreateToDoList:: Error while creating ToDo list '%s'. error={%v}",
			req.Name, err))
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
	}
	toDoListInfo := structutils.GetToDoListsInfo()
	logutils.Info.Println(fmt.Sprintf(
		"CreateToDoList:: new ToDo '%s' list created. Set of ToDoList: {%s}", req.Name, toDoListInfo ))
	fmt.Fprintf(w, toDoListInfo)
}	

/* 
	request type: DELETE
	url: /delete/ {"Name": "ToDo list"}
	The request body must contain a JSON object with a Name field

	Examples:

	   req: DELETE /delete/ {"Name": ""}
	   res: 400 empty name
	   
	   req: DELETE /delete/ {"Name": "wrong"}
	   res: 404 ToDo list not found

	   req: POST /delete/ {"name": "ToDo List"}
	   res: 200
*/
func DeleteToDoList(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	req := struct{ Name string }{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Name == "" {
		logutils.Error.Println("DeleteToDoList:: Bad request received", err)
		http.Error(w, "Missing ToDo list name", http.StatusBadRequest)
		return
	}

	err := structutils.DeleteToDoList(req.Name)
	if err != nil {
		logutils.Error.Println(fmt.Sprintf(
			"DeleteToDoList:: Error while deleting ToDo list '%s'. error={%v}",
			req.Name, err))
			http.Error(w, err.Error(), http.StatusNotFound)
			return
	}
	
	toDoListInfo := structutils.GetToDoListsInfo()
	logutils.Info.Println(fmt.Sprintf(
		"DeleteToDoList:: ToDo list '%s' deleted. Set of ToDoList: {%s}",req.Name, toDoListInfo ))
	fmt.Fprintf(w, toDoListInfo)
}	


/* 
	request type: GET
	url: /showall/

	Examples:

	   req: GET /showall/ 
	   res: 200
	   
*/
func ShowAllToDoList(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	todoList := structutils.GetAllToDoList()
	logutils.Info.Println(fmt.Sprintf(
		"ShowAllToDoList:: retrieved %d todo list", len(todoList) ))
	json.NewEncoder(w).Encode(todoList)
}	

/* 
	request type: GET
	url: /show/:list/

	Examples:

	   req: GET /show// 
	   res: 400 wrong name
	   
	   req: GET /show// 
	   res: 404 ToDo list not found

	   req: GET /show/okname/ 
	   res: 200
*/
func ShowToDoList(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	key := param.ByName("list")

	if key == "" {
		logutils.Error.Println("ShowToDoList:: Bad request received. No list name provided")
		http.Error(w, "Missing ToDo list name", http.StatusBadRequest)
		return
	}

	list, err := structutils.GetToDoList(key)
	if err != nil {
		logutils.Error.Println(fmt.Sprintf(
			"ShowToDoList:: Error while retrieving ToDo list '%s'. error={%v}",
			key, err))
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	logutils.Info.Println(fmt.Sprintf(
		"ShowToDoList:: Retrieved ToDoList '%s': tasks={%v}",key, list.GetAll() ))
	json.NewEncoder(w).Encode(list)
}	
