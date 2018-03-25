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
	TASK_BADREQUEST = 20;
	TASK_OPERATION_ERROR = 21;
)

/* 
	request type: POST
	url: /lists/:list/tasks {"Title": "New Task"}
	The request body must contain a JSON object with a Title field

	Examples:

	   req: POST /lists/oklist/tasks {"Title": ""}
	   res: 400 empty title or ToDo list info

	   req: POST /lists/wronglist/tasks {"Title": "Task Title"}
	   res: 404 ToDo list not found

	   req: POST /lists/:list/tasks {"Title": "Task already inserted"}
	   res: 404 Task already present in List

	   req: POST /lists/oklist/tasks {"Title": "New Task"}
	   res: 200
*/	   
func CreateTask(w http.ResponseWriter, r *http.Request, param httprouter.Params)  {
	key := param.ByName("list")
	req := struct{ Title string }{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil  || key == "" || req.Title == "" {
		taskBadRequestError(w, "CreateTask", err)		
		return		
	}
	
	task, err :=  model.AddTask(key, req.Title)
	if err != nil {
		taskOperationError(w, "CreateTask", req.Title, key, err)
		return
	}

	logutils.Info.Println(fmt.Sprintf(
		"CreateTask:: new task added to ToDoList '%s': task={title: %s, done=%t}",key, task.Title, task.Done ))
	json.NewEncoder(w).Encode(task)
}


/* 
	request type: DELETE
	url: /lists/:list/tasks/:task

	Examples:

	   req: DELETE /lists/oklist/tasks//
	   res: 400 empty title
	   
	   req: POST /lists/wronglist/tasks/oktask
	   res: 404 ToDo list not found

	   req: DELETE /lists/oklist/tasks/oktask
	   res: 200
*/	   
func DeleteTask(w http.ResponseWriter, r *http.Request, param httprouter.Params)  {
	key := param.ByName("list")
	title := param.ByName("task")

	if  key == "" || title == ""  {
		taskBadRequestError(w, "DeleteTask", errors.New("Missing mandatory information: todolist name or task title"))		
		return	
	}
	task, err :=  model.RemoveTask(key, title)
	if err != nil {
		taskOperationError(w, "DeleteTask", title, key, err)
		return
	}

	logutils.Info.Println(fmt.Sprintf(
		"DeleteTask:: task removed from  ToDoList '%s': task={title: %s, done=%t}",key, task.Title, task.Done ))
	json.NewEncoder(w).Encode(task)
}


/* 
	request type: GET
	url: /lists/:list/tasks/:task

	Examples:

	   req: GET /lists/oklist/tasks//
	   res: 400 empty title
	   
	   req: POST /lists/wronglist/tasks/oktitle
	   res: 404 ToDo list not found

	   req:  GET /lists/oklist/tasks/oktitle
	   res: 200
*/	   
func GetTask(w http.ResponseWriter, r *http.Request, param httprouter.Params)  {
	key := param.ByName("list")
	title := param.ByName("task")
	if  key == "" || title == "" {
		taskBadRequestError(w, "GetTask", errors.New("Missing mandatory information: todolist name or task title"))		
		return
	}
	
	task, err :=  model.GetTask(key, title)
	if err != nil {
		taskOperationError(w, "GetTask", title, key, err)
		return
	}
	
	logutils.Info.Println(fmt.Sprintf(
		"GetTask:: task retrieved from ToDoList '%s': task={title: %s, done=%t}",key, task.Title, task.Done ))
	json.NewEncoder(w).Encode(task)
}

/* 
	request type: PUT
	url: /lists/:list/tasks/:task {"Title": "New Title", "Done": true}
	The request body must contain a JSON object with  Done fields and optional a Title with the new title

	Examples:

	   req: POST /oklist/addtask/ {"Title": "", "Done": true}
	   res: 400 empty title

	   req: POST /addtask/badlist/ {"Title": "", "Done": true}
	   res: 404 ToDo list not found

	   req: POST /addtask/oklist/ {"Title": "Task already inserted", "Done": true}
	   res: 406 Task unknown. Use addtask request to add the task

	   req: POST /oklist/show/: {"Title": "New Task", "Done": true}
	   res: 200
*/	   
func UpdateTask(w http.ResponseWriter, r *http.Request, param httprouter.Params)  {
	key := param.ByName("list")	
	title := param.ByName("task")
	req := struct{ 
		Title string
		Done  bool }{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || key == "" || title == "" {
		taskBadRequestError(w, "UpdateTask", err)		
		return
	}
	if req.Title == "" {
		req.Title = title
	}
	task, err :=  model.UpdateTask(key, title, req.Title, req.Done)
	if err != nil {
		taskOperationError(w, "UpdateTask", title, key, err)
		return
	}

	logutils.Info.Println(fmt.Sprintf(
		"UpdateTask:: task updated in  ToDoList '%s': task={title: %s, done=%t}",key, task.Title, task.Done ))
	json.NewEncoder(w).Encode(task)
}

func taskBadRequestError(w http.ResponseWriter, caller string, err error){
	HandleError(w, http.StatusBadRequest, TASK_BADREQUEST, caller,
		"Missing ToDo list name or task title",  
		fmt.Sprintf("Bad request received: Missing mandatory parameters list or title. %v", err))
}

func taskOperationError(w http.ResponseWriter, caller, task, todolist string, err error){
	HandleError(w, http.StatusNotFound, TASK_OPERATION_ERROR, caller,
		fmt.Sprintf("Error while performing operation on task = {%s}, ToDo list = {%s}", task, todolist),  
		fmt.Sprintf("%v",err))
}