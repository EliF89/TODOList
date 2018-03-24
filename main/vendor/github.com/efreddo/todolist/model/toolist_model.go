package model


type ToDoList struct {
	Name string			
	TaskNumber int		
	Tasks  []*Task			
}

func CreateToDoList(name string) (*ToDoList, error) {
	db := getConnection()
	defer db.Close()

	list := new(ToDoList)	
	query :=  ` INSERT INTO todolist (name)
				VALUES ($1)
				RETURNING name`
	err := db.QueryRow(query, name).Scan(&list.Name)
		
	if err != nil {
		return &ToDoList{}, err
	} else {
		return list, nil
	}
}

func GetAllToDoList() ([]*ToDoList, error) {
	db := getConnection()
	defer db.Close()

	query :=  ` SELECT l.name, count(t.Title) as numTask 
				FROM todolist as l LEFT OUTER JOIN task as t ON l.name = t.todolist_name
				GROUP BY l.name, t.Title`
	rows, err := db.Query(query)
	defer rows.Close()

	if err != nil {
		return nil, err
	}
	todoLists := make([]*ToDoList, 0)
	for rows.Next() {
		list := new(ToDoList)
		err := rows.Scan(&list.Name, &list.TaskNumber)
		if err != nil {
			return nil, err
		}
		todoLists = append(todoLists, list)
	}
	
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return todoLists, nil
}

func GetToDoList(name string) (*ToDoList, error) {
	db := getConnection()
	defer db.Close()

	list := new(ToDoList)
	query :=  ` SELECT l.name, count(t.Title) as numTask 
				FROM todolist as l LEFT OUTER JOIN task as t ON l.name = t.todolist_name
				WHERE l.name = $1
				GROUP BY l.name, t.Title`
	err := db.QueryRow(query, name).Scan(&list.Name, &list.TaskNumber)
	
	if err != nil {
		return &ToDoList{}, err
	} else {
		return list, nil
	}
}


func DeleteToDoList(name string) (*ToDoList, error) {
	db := getConnection()
	defer db.Close()

	list := new(ToDoList)
	query :=  ` DELETE FROM todolist
				WHERE  name = $1
				RETURNING name` 
	err := db.QueryRow(query, name).Scan(&list.Name)
		
	if err != nil {
		return &ToDoList{}, err
	} else {
		return list, nil
	}
}

func UpdateToDoList(name string, newName string) (*ToDoList, error) {
	db := getConnection()
	defer db.Close()

	list := new(ToDoList)	
	query :=  ` UPDATE todolist 
				SET name = $1
				WHERE name = $2
				RETURNING name`
	err := db.QueryRow(query, newName, name).Scan(&list.Name)
		
	if err != nil {
		return &ToDoList{}, err
	} else {
		return list, nil
	}
}
