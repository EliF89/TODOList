package model

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/efreddo/todolist/logutils"
	_ "github.com/lib/pq"
  )

var (
	host     = os.Getenv("POSTGRES_HOST")
	port     = os.Getenv("POSTGRES_PORT")
	user     = os.Getenv("POSTGRES_USER")
	password = os.Getenv("POSTGRES_PASSWORD")
	dbname   = os.Getenv("POSTGRES_DB")
 )

 func getConnection() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	logutils.Error.Println("DB Connect: ", psqlInfo)

	db, err := sql.Open("postgres", psqlInfo);
	 if  err != nil {		
		logutils.Error.Println("database InitConnection:: error while opening DB connection",  err.Error())
		logutils.Error.Println("DB Connect: ", psqlInfo)
		return nil
	}
		
	if err := db.Ping(); err != nil {
		logutils.Error.Println("database InitConnection:: error while opening DB connection",  err.Error())
		return nil
	}
	
	logutils.Info.Println("database InitConnection:: connected")
	return db
}

