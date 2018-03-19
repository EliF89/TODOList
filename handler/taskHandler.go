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
	url: /addtask/:list/ {"Title": "New Task"}
	The request body must contain a JSON object with a Title field

	Examples:

	   req: POST /addtask/:list/ {"Title": ""}
	   res: 400 empty title or ToDo list info

	   req: POST /addtask/:list/ {"Title": ""}
	   res: 404 ToDo list not found

	   req: POST /addtask/:list/ {"Title": "Task already inserted"}
	   res: 406 Task already present in List. Use updatetask request to modfy the task

	   req: POST /addtask/:list/ {"Title": "New Task"}
	   res: 200
*/	   
func AddTask(w http.ResponseWriter, r *http.Request, param httprouter.Params)  {
	key := param.ByName("list")
	req := struct{ Title string }{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil  || key == "" || req.Title == "" {
		logutils.Error.Println("AddTask:: Bad request received", err)
		http.Error(w, "Missing ToDo list name or task title", http.StatusBadRequest)
		return
	}

	list, err := structutils.GetToDoList(key)
	if err != nil {
		logutils.Error.Println(fmt.Sprintf(
			"AddTask:: Error while retrieving ToDo list %s. error={%v}", key, err))
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	// Check if task already exist in the list
	_, found := list.GetTask(req.Title)
	if found {		
		logutils.Warning.Println(fmt.Sprintf(
			"AddTask:: task {%s} already present in ToDo list {%s}. Task not added.",
			req.Title, key))
		http.Error(w, "Task already present in List. Use updatetask to modfy", http.StatusNotAcceptable)
		return	
	}

	task, err := list.AddTask(req.Title)
	if err != nil {
		logutils.Error.Println(fmt.Sprintf(
			"AddTask:: Error while creating task = {%s} in ToDo list = {%s}. error={%v}",
			req.Title, key, err))
	}
	json.Marshal(list.GetAll()) 
	logutils.Info.Println(fmt.Sprintf(
		"AddTask:: new task added to ToDoList '%s': task={title: %s, done=%t}",key, task.Title, task.Done ))
	json.NewEncoder(w).Encode(list.GetAll())
}



/* 
	request type: PUT
	url: /updatetask/:list/ {"Title": "New Task", "Done": true}
	The request body must contain a JSON object with a Title field

	Examples:

	   req: POST /oklist/addtask/ {"Title": "", "Done": true}
	   res: 400 empty title

	   req: POST /addtask/badlist/ {"Title": ""}
	   res: 404 ToDo list not found

	   req: POST /addtask/oklist/ {"Title": "Task already inserted"}
	   res: 406 Task unknown. Use addtask request to add the task

	   req: POST /oklist/show/: {"Title": "New Task", "Done": true}
	   res: 200
*/	   
func UpdateTask(w http.ResponseWriter, r *http.Request, param httprouter.Params)  {
	key := param.ByName("list")
	req := struct{ 
		Title string
		Done  bool }{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || key == "" || req.Title == "" {
		logutils.Error.Println("UpdateTask:: Bad request received", err)
		http.Error(w, "Missing ToDo list name or task title", http.StatusBadRequest)
		return
	}
	list, err := structutils.GetToDoList(key)
	if err != nil {
		logutils.Error.Println(fmt.Sprintf(
			"UpdateTask:: Error while retrieving ToDo list %s. error={%v}", key, err))
		http.Error(w, err.Error(), http.StatusNotFound)
		return	
	}

	task, err := list.UpdateTask(req.Title, req.Done)
	if err != nil {
		logutils.Error.Println(fmt.Sprintf(
			"UpdateTask:: Error while updating task = {%s} in ToDo list = {%s}. error={%v}",
			req.Title, key, err))
		http.Error(w, "Task unknown. Use addtask request to add the task", http.StatusNotAcceptable)
		return			
	}

	logutils.Info.Println(fmt.Sprintf(
		"UpdateTask:: task updated in  ToDoList '%s': task={title: %s, done=%t}",key, task.Title, task.Done ))
	json.NewEncoder(w).Encode(list.GetAll())
}


/* 
	request type: DELETE
	url: /deletetask/:list/ {"Title": "Title Task"}
	The request body must contain a JSON object with a Title field

	Examples:

	   req: DELETE /deletetask/:list/ {"Title": ""}
	   res: 400 empty title
	   
	   req: POST /deletetask/badlist/ {"Title": ""}
	   res: 404 ToDo list not found

	   req: DELETE /deletetask/:list/ {"Title": "Title Task"}
	   res: 200
*/	   
func RemoveTask(w http.ResponseWriter, r *http.Request, param httprouter.Params)  {
	key := param.ByName("list")
	req := struct{ 
		Title string
		Done  bool }{}
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil || key == "" || req.Title == ""  {
		logutils.Error.Println("RemoveTask:: Bad request received", err)
		http.Error(w, "Missing ToDo list name or task title", http.StatusBadRequest)
		return
	}
	list, err := structutils.GetToDoList(key)
	if err != nil {
		logutils.Error.Println(fmt.Sprintf(
			"RemoveTask:: Error while retrieving ToDo list %s. error={%v}", key, err))
			http.Error(w, err.Error(), http.StatusNotFound)
			return
	}

	task, err := list.RemoveTask(req.Title)
	if err != nil {
		logutils.Error.Println(fmt.Sprintf(
			"RemoveTask:: Error while removing task = {%s} in ToDo list = {%s}. error={%v}",
			req.Title, key, err))
			http.Error(w, err.Error(), http.StatusNotFound)
			return
	}

	logutils.Info.Println(fmt.Sprintf(
		"RemoveTask:: task removed from  ToDoList '%s': task={title: %s, done=%t}",key, task.Title, task.Done ))
	json.NewEncoder(w).Encode(list.GetAll())
}


/* 
	request type: GET
	url: /showtask/:list/

	Examples:

	   req: GET /showtask/oklist//
	   res: 400 empty title
	   
	   req: POST /showtask/badlist/oktitle
	   res: 404 ToDo list not found

	   req:  GET /showtask/oklist/oktitle
	   res: 200
*/	   
func ShowTask(w http.ResponseWriter, r *http.Request, param httprouter.Params)  {
	key := param.ByName("list")
	title := param.ByName("title")
	if  key == "" || title == "" {
		logutils.Error.Println(fmt.Sprintf("ShowTask:: Bad request received, list = '%s', task='%s'", key, title))
		http.Error(w, "Missing ToDo list name or task title", http.StatusBadRequest)
		return
	}
	list, err := structutils.GetToDoList(key)
	if err != nil {
		logutils.Error.Println(fmt.Sprintf(
			"ShowTask:: Error while retrieving ToDo list %s. error={%v}", key, err))
		http.Error(w, err.Error(), http.StatusNotFound)
		return	
	}

	task, found := list.GetTask(title)
	if !found {
		logutils.Error.Println(fmt.Sprintf(
			"ShowTask:: Task = {%s} not found in ToDo list = {%s}",
			title, key))
		http.Error(w,"Task not found", http.StatusNotFound)
		return	
	}
	
	logutils.Info.Println(fmt.Sprintf(
		"ShowTask:: task retrieved from ToDoList '%s': task={title: %s, done=%t}",key, task.Title, task.Done ))
	json.NewEncoder(w).Encode(task)
}