package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/efreddo/todolist/logutils"
	"github.com/efreddo/todolist/model"
	"github.com/julienschmidt/httprouter"
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
		logutils.Error.Println("AddTask:: Bad request received", err)
		http.Error(w, "Missing ToDo list name or task title", http.StatusBadRequest)
		return
	}

	task, err :=  model.AddTask(key, req.Title)
	if err != nil {
		logutils.Error.Println(fmt.Sprintf(
			"AddTask:: Error while creating task = {%s} in ToDo list = {%s}. error={%v}",
			req.Title, key, err))
		http.Error(w, fmt.Sprintf("Error while creating task '%s' in ToDo list '%s'", req.Title, key), 
				http.StatusNotFound)
		return
	}

	logutils.Info.Println(fmt.Sprintf(
		"AddTask:: new task added to ToDoList '%s': task={title: %s, done=%t}",key, task.Title, task.Done ))
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
	req := struct{ 
		Title string
		Done  bool }{}
	if  key == "" || title == ""  {
		logutils.Error.Println("RemoveTask:: Bad request received. Null list name or task title")
		http.Error(w, "Missing ToDo list name or task title", http.StatusBadRequest)
		return
	}
	task, err :=  model.RemoveTask(key, title)
	if err != nil {
		logutils.Error.Println(fmt.Sprintf(
			"RemoveTask:: Error while removing task = {%s} in ToDo list = {%s}. error={%v}",
			req.Title, key, err))
			http.Error(w, fmt.Sprintf("Error while deleting task '%s' in ToDo list '%s'", title, key),  
				http.StatusNotFound)
			return
	}

	logutils.Info.Println(fmt.Sprintf(
		"RemoveTask:: task removed from  ToDoList '%s': task={title: %s, done=%t}",key, task.Title, task.Done ))
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
		logutils.Error.Println(fmt.Sprintf("ShowTask:: Bad request received, list = '%s', task='%s'", key, title))
		http.Error(w, "Missing ToDo list name or task title", http.StatusBadRequest)
		return
	}
	
	task, err :=  model.GetTask(key, title)
	if err != nil {
		logutils.Error.Println(fmt.Sprintf(
			"ShowTask:: Task = {%s} not found in ToDo list = {%s}",
			title, key))
		http.Error(w, fmt.Sprintf("Error while retrieving task '%s' from ToDo list '%s'", title, key), http.StatusNotFound)
		return	
	}
	
	logutils.Info.Println(fmt.Sprintf(
		"ShowTask:: task retrieved from ToDoList '%s': task={title: %s, done=%t}",key, task.Title, task.Done ))
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
		logutils.Error.Println("UpdateTask:: Bad request received", err)
		http.Error(w, "Missing ToDo list name or task title", http.StatusBadRequest)
		return
	}
	if req.Title == "" {
		req.Title = title
	}
	task, err :=  model.UpdateTask(key, title, req.Title, req.Done)
	if err != nil {
		logutils.Error.Println(fmt.Sprintf(
			"UpdateTask:: Error while updating task = {%s} in ToDo list = {%s}. error={%v}",
			title, key, err))
		http.Error(w, fmt.Sprintf("Error while updating task '%s' in ToDo list '%s'", title, key), http.StatusNotFound)
		return	
	}

	logutils.Info.Println(fmt.Sprintf(
		"UpdateTask:: task updated in  ToDoList '%s': task={title: %s, done=%t}",key, task.Title, task.Done ))
	json.NewEncoder(w).Encode(task)
}

