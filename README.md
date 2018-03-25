# ToDolist



## Services 

The todo server exposes json services for getting, putting, and deleting todo lists and tasks.

- ToDo list services

Create a new list with name "ToDo list name":
```
POST /lists/ 
Body: {"name": "<ToDo list name>"}
Reponse: {"Name":"<ToDo list name>","Tasks":null,"TaskNumber":0}
```

Modify the name of ToDo "ToDo list name" to "New ToDo list name":
```
PUT /lists/<ToDo list name>/ 	
Body: {"Name": "<New ToDo list name>"}
Reponse: {"Name":"<New ToDo list name>","Tasks":null,"TaskNumber":0}
```

Get the requested ToDo list "ToDo list name":
```
GET /lists/<ToDo list name>/ 	
Reponse: {"Name":"<New ToDo list name>","Tasks":null,"TaskNumber":0}
```

Get all the ToDo lists inserted
```
GET /lists/
Response: [{"Name":"<ToDo list 1>","Tasks":null,"TaskNumber":0}, {"Name":"<ToDo list 2>","Tasks":null,"TaskNumber":0}]
```

Delete a ToDo list
```
DELETE /lists/<ToDo list name>/ 	
Reponse: {"Name":"<ToDo list name>","Tasks":null,"TaskNumber":0}
```



- Task services

Add a task in a ToDo list
```
POST /lists/<ToDo list name>/tasks 
Body: {"Title": "<Task Title>"}
Reponse: {"ToDoList":"<ToDo list name>","Title":"<Task Title>","Done":false}
```

Get task "Task Title" in list "ToDo list name":
```
GET /lists/<ToDo list name>/tasks/<Task Title>
Reponse: {"ToDoList":"<ToDo list name>","Title":"<Task Title>","Done":false/true}
```

Update task "Task Title" in ToDo list "ToDo list name" to modify name and status (done/not done)
```
POST /lists/<ToDo list name>/tasks/<Task Title>
Body: {"Title": "<Task Title>", "Done": true}
Reponse: {"ToDoList":"<ToDo list name>","Title":"<Task Title>","Done":true}
```

Delete task "Task Title" from ToDo list "ToDo list name"
```
DELETE /lists/<ToDo list name>/tasks/<Task Title>
Reponse: {"ToDoList":"<ToDo list name>","Title":"<Task Title>","Done":false/true}
```

- Errors

json response in case of errors:
```
{"Errors":
  [
   {
	  "Code":<internal code>,
	  "ErrorMessage":"<Error message for the user>",
		"TechnicalReason":"<Technical message>"
		}
	]
}
```

## Tests

Unit test are provided to test list and task functionalities:  

```
cd model
go test 
```

Postman tests are also available. To run postman tests with newman:

```
cd postmantests
newman run todoLists.test.json -g postman_globals.json
```


## Docker Hub

The todo list application is deployed on docker hub. To run it locally:

```
docker pull elisafreddo/todolist
docker images //list the image availables
docker run -p 8080:8080 <docker image>  
```

## Go dependencies

- github.com/julienschmidt/httprouter



