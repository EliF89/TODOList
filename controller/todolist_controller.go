package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/efreddo/v1/todolist/model"
	"github.com/efreddo/v1/todolist/logutils"
	"github.com/julienschmidt/httprouter"
)

const (
	TODOLIST_BADREQUEST = 10;
	TODOLIST_OPERATION_ERROR = 11;
)

/* 
	request type: POST
	url: /lists/ {"Name": "New ToDo list"}
	The request body must contain a JSON object with a Name field

	Examples:

	   req: POST /lists/ {"Name": ""}
	   res: 400 empty name

	   req: POST /create/ {"name": "New ToDo List"}
	   res: 200
*/
func CreateToDoList(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	req := struct{ Name string }{}
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Name == "" {
		todolistBadRequestError(w, "CreateToDoList", err)		
		return		
	}

	toDoList, err :=  model.CreateToDoList(req.Name)
	if err != nil {				
		todolistOperationError(w, "CreateToDoList", req.Name, err)
		return
	}

	logutils.Info.Println(fmt.Sprintf(
		"CreateToDoList:: new ToDo '%s' list created", toDoList.Name ))
	json.NewEncoder(w).Encode(toDoList)
}	

/* 
	request type: DELETE
	url: /lists/:list/ 

	Examples:

	   req: DELETE /lists//
	   res: 400 empty name
	   
	   req: DELETE /lists/wronglist/
	   res: 404 ToDo list not found

	   req: POST /lists/oklist/ 
	   res: 200
*/
func DeleteToDoList(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	key := param.ByName("list")
	if key == "" {
		todolistBadRequestError(w, "DeleteToDoList", errors.New("Missing mandatory information: todolist name."))	
		return
	}

	list, err :=  model.DeleteToDoList(key)
	if err != nil {
		todolistOperationError(w, "DeleteToDoList", key, err)
		return
	}

	logutils.Info.Println(fmt.Sprintf(
		"DeleteToDoList:: ToDo list '%s' deleted", list.Name ))
	json.NewEncoder(w).Encode(list)	
}	


/* 
	request type: PUT
	url: /lists/:list/
	The request body must contain a JSON object with a Name field

	Examples:

	   req: PUT /lists//  {"Name": "New ToDo list"}
	   res: 400 wrong name
	   
	   req: PUT /lists/oklist/  {"Name": ""}
	   res: 400 wrong name

	   req: PUT /lists/wrongname/ 	{"Name": "New name"}
	   res: 404 ToDo list not found

	   req: PUT /lists/okname/ 	{"Name": "New name"}
	   res: 200

*/
func UpdateToDoList(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	key := param.ByName("list")
	req := struct{ Name string }{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || req.Name == "" || key == "" {
		todolistBadRequestError(w, "UpdateToDoList", err)	
		return
	}

	list, err :=  model.UpdateToDoList(key, req.Name)
	if err != nil {
		todolistOperationError(w, "UpdateToDoList", key, err)
		return
	}
	logutils.Info.Println(fmt.Sprintf(
		"UpdateToDoList:: Retrieved ToDoList '%s'. Number of task={%d}",key, list.TaskNumber ))
	json.NewEncoder(w).Encode(list)
}	

/* 
	request type: GET
	url: /lists/

	Examples:

	   req: GET /lists/
	   res: 404 Error while retrieving lists

	   req: GET /lists/
	   res: 200
	   
*/
func GetAllToDoList(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	todoList, err :=  model.GetAllToDoList()
	if err != nil {
		todolistOperationError(w, "GetAllToDoList", "all", err)
		return
	}	
	logutils.Info.Println(fmt.Sprintf(
		"GetAllToDoList:: retrieved %d todo list", len(todoList) ))
	json.NewEncoder(w).Encode(todoList)
}	

/* 
	request type: GET
	url: /lists/:list/

	Examples:

	   req: GET /lists// 
	   res: 400 wrong name
	   
	   req: GET /lists/wrongname/ 
	   res: 404 ToDo list not found

	   req: GET /lists/okname/ 
	   res: 200
*/
func GetToDoList(w http.ResponseWriter, r *http.Request, param httprouter.Params) {
	key := param.ByName("list")

	if key == "" {
		todolistBadRequestError(w, "GetToDoList", errors.New("Missing mandatory information: todolist name."))	
		return
	}

	list, err :=  model.GetToDoList(key)
	if err != nil {
		todolistOperationError(w, "GetToDoList", key, err)
		return
	}

	logutils.Info.Println(fmt.Sprintf(
		"GetToDoList:: Retrieved ToDoList '%s'. Number of task={%d}",key, list.TaskNumber ))
	json.NewEncoder(w).Encode(list)
}	

func todolistBadRequestError(w http.ResponseWriter, caller string, err error){
	HandleError(w, http.StatusBadRequest, TODOLIST_BADREQUEST, caller,
		"Missing ToDo list name",  
		fmt.Sprintf("Bad request received: Missing mandatory parameters list name. %v", err))
}

func todolistOperationError(w http.ResponseWriter, caller, todolist string, err error){
	HandleError(w, http.StatusNotFound, TODOLIST_OPERATION_ERROR, caller,
		fmt.Sprintf("Error while performing operation on ToDo list = {%s}", todolist),  
		fmt.Sprintf("%v",err))
}