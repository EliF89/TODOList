package model

type Task struct {
	ToDoList  string  
	Title 	string 
	Done  bool   
}

func AddTask(todoListName string, taskTitle string) (*Task, error) {
	db := getConnection()
	defer db.Close()

	task := new(Task)
	query :=  ` INSERT INTO task (todolist_name, title)
				VALUES ($1, $2) RETURNING todolist_name, title, done` 

	err := db.QueryRow(query, todoListName, taskTitle).Scan(&task.ToDoList, &task.Title, &task.Done)
		
	if err != nil {
		return &Task{}, err
	} else {
		return task, nil
	}
}


func GetTask(todoListName string, taskTitle string) (*Task, error) {
	db := getConnection()
	defer db.Close()

	task := new(Task)
	query :=  ` SELECT todolist_name, title, done
				FROM task
				WHERE  todolist_name = $1 AND title = $2` 

	err := db.QueryRow(query, todoListName, taskTitle).Scan(&task.ToDoList, &task.Title, &task.Done)
		
	if err != nil {
		return &Task{}, err
	} else {
		return task, nil
	}
}


func UpdateTask(todoListName string, taskTitle string, newTitle string, done bool) (*Task, error) {
	db := getConnection()
	defer db.Close()

	task := new(Task)
	query :=  ` UPDATE task 
				SET done = $1, title = $2
				WHERE  todolist_name = $3 AND title = $4
				RETURNING todolist_name, title, done` 

	err := db.QueryRow(query, done, newTitle, todoListName, taskTitle).Scan(&task.ToDoList, &task.Title, &task.Done)
		
	if err != nil {
		return &Task{}, err
	} else {
		return task, nil
	}
}


func RemoveTask(todoListName string, taskTitle string) (*Task, error) {
	db := getConnection()
	defer db.Close()

	task := new(Task)
	query :=  ` DELETE FROM task
				WHERE  todolist_name = $1 AND title = $2
				RETURNING todolist_name, title, done` 
	err := db.QueryRow(query, todoListName, taskTitle).Scan(&task.ToDoList, &task.Title, &task.Done)
		
	if err != nil {
		return &Task{}, err
	} else {
		return task, nil
	}
}
