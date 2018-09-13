package util

import (
	// standard library
	"fmt"
	"net/http"
	"encoding/json"
	// "io/ioutil"
	// "time"
	// project files
	// "github.com/patientplatypus/webserver/utilities"
	// "github.com/patientplatypus/webserver/database"
)


var UpdateJWT = "" //if user is caught across JWT token update boundary, send updated jwt token.

type LoginResponse struct {
	Status string
	Logintype string
	Message string
	JWT string
}

type RequestResponse struct {
	Status string
	Message string
	JWT string
}

type RequestResponseArray struct {
	Name string
	Value []string
	JWT string
}

func HandleLoginResponse(w http.ResponseWriter, status string, logintype string, message string, jwt string){
	fmt.Println("inside HandleLoginResponse")
	fmt.Println(status)
	fmt.Println(logintype)
	fmt.Println(message)
	fmt.Println(jwt)
	response := LoginResponse{Status: status, Logintype: logintype, Message: message, JWT: jwt}
	json.NewEncoder(w).Encode(response)	
}

func HandleRequestResponse(w http.ResponseWriter, status string, message string){
	if UpdateJWT != ""{
		response := RequestResponse{Status: status, Message: message, JWT: UpdateJWT}
		json.NewEncoder(w).Encode(response)	
	}else{
		response := RequestResponse{Status: status, Message: message}
		json.NewEncoder(w).Encode(response)	
	}
}

func HandleRequestResponseArray(w http.ResponseWriter, name string, value []string){
	if UpdateJWT != ""{
		response := RequestResponseArray{Name: name, Value: value, JWT: UpdateJWT}
		json.NewEncoder(w).Encode(response)	
	}else{
		response := RequestResponseArray{Name: name, Value: value}
		json.NewEncoder(w).Encode(response)	
	}
}