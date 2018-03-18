package main

import (
		"log"
		"net/http"
		"io/ioutil"
		"os"

		"github.com/efreddo/todolist/handler"
		"github.com/efreddo/todolist/logutils"
		"github.com/julienschmidt/httprouter"
)


// Status code in reply = StatusBadRequest
type badRequest struct{ error }

// Status code in reply = StatusNotFound
type notFound struct{ error }

func main(){
	// init logs
	logutils.InitLogs(ioutil.Discard, os.Stdout, os.Stdout, os.Stderr)

	RegisterHandlers()
}


func RegisterHandlers() {
	defer errorHandler()

	r := httprouter.New()

	// ToDo Lists 
	r.POST("/create/", handler.CreateToDoList)	
	r.DELETE("/delete/", handler.DeleteToDoList)
	r.GET("/showall/", handler.ShowAllToDoList)
	r.GET("/show/:list/", handler.ShowToDoList)

	// Tasks
	r.POST("/addtask/:list/", handler.AddTask)	
	r.DELETE("/deletetask/:list/", handler.RemoveTask)	
	r.PUT("/updatetask/:list/", handler.UpdateTask)	
	r.GET("/showtask/:list/:title", handler.ShowTask)

	http.ListenAndServe(":8080" , r)	
	
}

func errorHandler(){
	err := recover()

	if err == nil {
		return
	}

	log.Fatal("ListenAdnServe: ", err )
/*
	switch err.(type) {
	case badRequest:
		http.Error(w, err.Error(), http.StatusBadRequest)
	case notFound:
		http.Error(w, "task not found", http.StatusNotFound)
	default:
		log.Println(err)
		http.Error(w, "oops", http.StatusInternalServerError)
	}
	*/
}

