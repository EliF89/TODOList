package controller
import (
	"fmt"
	"encoding/json"
	"net/http"

	"github.com/efreddo/todolist/logutils"
)


type ListError struct{
	Errors []CustomError
}

type CustomError struct {
	Code  int  
	ErrorMessage string 
	TechnicalReason  string   
}


func HandleError(w http.ResponseWriter, httpCode, internalCode int, caller,  message, techReason string)  {

	var errorString string
	listErrors := new(ListError)
	errors := make([]CustomError, 1)	
	errors[0] = CustomError{ 
		Code: internalCode,
		ErrorMessage: message,
		TechnicalReason: techReason}
	listErrors.Errors = errors

	errorJson, err := json.Marshal(listErrors)
	if err != nil {
		fmt.Printf("Error: %s", err)
		errorString = message;
	}else{
		errorString = string(errorJson)
	}
	logutils.Error.Println(fmt.Sprintf("%s:: %s. Reason={%s}",caller, message, techReason))
	http.Error(w, errorString, httpCode)

}
