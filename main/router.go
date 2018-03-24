package main

import (
		"net/http"
		"io/ioutil"
		"os"
		"fmt"
		"github.com/efreddo/todolist/controller"
		"github.com/efreddo/todolist/logutils"
		"github.com/julienschmidt/httprouter"
)

func main(){
	logutils.InitLogs(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)
	RegisterHandlers()
}

func RegisterHandlers() {

	r := httprouter.New()

	// test
	r.GET("/test/", testWorking)

	// ToDo Lists 
	r.POST("/lists/", controller.CreateToDoList)	
	r.DELETE("/lists/:list", controller.DeleteToDoList)
	r.PUT("/lists/:list",  controller.UpdateToDoList)	
	r.GET("/lists/", controller.GetAllToDoList)
	r.GET("/lists/:list/", controller.GetToDoList)

	// Tasks
	r.POST("/lists/:list/tasks",  controller.CreateTask)	
	r.DELETE("/lists/:list/tasks/:task",  controller.DeleteTask)	
	r.PUT("/lists/:list/tasks/:task",  controller.UpdateTask)	
	r.GET("/lists/:list/tasks/:task",  controller.GetTask)

	http.ListenAndServe(":8080" , r)	
	
}

func testWorking(w http.ResponseWriter, r *http.Request, param httprouter.Params){
	fmt.Fprintf(w, "WORKING!!!")
}