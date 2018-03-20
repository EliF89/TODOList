package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/efreddo/todolist/model"
	"github.com/efreddo/todolist/logutils"
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

	toDoList, err :=  model.CreateToDoList(req.Name)
	if err != nil {				
		logutils.Error.Println(fmt.Sprintf(
			"CreateToDoList:: Error while creating ToDo list '%s'. error={%v}",
			req.Name, err))
			http.Error(w, fmt.Sprintf("Error while creating ToDo list '%s'", req.Name), http.StatusBadRequest)
			return
	}

	logutils.Info.Println(fmt.Sprintf(
		"CreateToDoList:: new ToDo '%s' list created", toDoList.Name ))
	json.NewEncoder(w).Encode(toDoList)
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

	list, err :=  model.DeleteToDoList(req.Name)
	if err != nil {
		logutils.Error.Println(fmt.Sprintf(
			"DeleteToDoList:: Error while deleting ToDo list '%s'. error={%v}",
			req.Name, err))
			http.Error(w, fmt.Sprintf("Error while deleting ToDo list %s",req.Name) , http.StatusNotFound)
			return
	}

	logutils.Info.Println(fmt.Sprintf(
		"DeleteToDoList:: ToDo list '%s' deleted", list.Name ))
	json.NewEncoder(w).Encode(list)	
}	


/* 
	request type: GET
	url: /showall/

	Examples:

	   req: GET /showall/ 
	   res: 200
	   
*/
func GetAllToDoList(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	todoList, err :=  model.GetAllToDoList()
	if err != nil {
		logutils.Error.Println(fmt.Sprintf(
			"ShowAllToDoList:: Error while retrieving ToDo list from DB. error={%v}",  err))
		http.Error(w, "Error while retrieving ToDo list", http.StatusUnprocessableEntity)
		return
	}	
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
func GetToDoList(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	key := param.ByName("list")

	if key == "" {
		logutils.Error.Println("ShowToDoList:: Bad request received. No list name provided")
		http.Error(w, "Missing ToDo list name", http.StatusBadRequest)
		return
	}

	list, err :=  model.GetToDoList(key)
	if err != nil {
		logutils.Error.Println(fmt.Sprintf(
			"ShowToDoList:: Error while retrieving ToDo list '%s'. error={%v}",
			key, err))
		http.Error(w, fmt.Sprintf("ToDo list not found"), http.StatusNotFound)
		return
	}
	logutils.Info.Println(fmt.Sprintf(
		"ShowToDoList:: Retrieved ToDoList '%s'. Number of task={%d}",key, list.TaskNumber ))
	json.NewEncoder(w).Encode(list)
}	
