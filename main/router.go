package main

import (
		"net/http"
		"io/ioutil"
		"os"
		
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

	// ToDo Lists 
	r.POST("/create/", controller.CreateToDoList)	
	r.DELETE("/delete/", controller.DeleteToDoList)
	r.GET("/showall/", controller.GetAllToDoList)
	r.GET("/show/:list/", controller.GetToDoList)

	// Tasks
	r.POST("/addtask/:list/",  controller.AddTask)	
	r.DELETE("/deletetask/:list/",  controller.RemoveTask)	
	r.PUT("/updatetask/:list/",  controller.UpdateTask)	
	r.GET("/showtask/:list/:title",  controller.GetTask)

	http.ListenAndServe(":8080" , r)	
	
}

