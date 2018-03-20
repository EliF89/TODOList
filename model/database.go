package model

import (
	"database/sql"
	"fmt"
  
	"github.com/efreddo/todolist/logutils"
	_ "github.com/lib/pq"
  )

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "pwd12345"
	dbname   = "dbtodolist"
  )

 func getConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
		
	db, err := sql.Open("postgres", psqlInfo);
	 if  err != nil {		
		logutils.Error.Println("database InitConnection:: error while opening DB connection",  err.Error())
		return nil
	}
		
	if err := db.Ping(); err != nil {
		logutils.Error.Println("database InitConnection:: error while opening DB connection",  err.Error())
		return nil
	}
	
	logutils.Info.Println("database InitConnection:: connected")
	return db
}

