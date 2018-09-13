package auth

import (
	// standard library
	"fmt"
	"net/http"
	"encoding/json"
	// "io/ioutil"
	// "time"
	// project files
	"github.com/patientplatypus/webserver/utilities"
	"github.com/patientplatypus/webserver/database"
)

type userRegister struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func RegisterUser(w http.ResponseWriter, req *http.Request) {
	fmt.Println("inside RegisterUser in Golang")

	var uR userRegister
	err := json.NewDecoder(req.Body).Decode(&uR)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	fmt.Println("here is the users email:")
	fmt.Println("%s", uR.Email)

	fmt.Println("here is the users password:")
	fmt.Println("%s", uR.Password)


	// Generate Random Number
	randFuncChannel := make(chan string)
	go util.RandMain(randFuncChannel)
	randRet := <-randFuncChannel //blocking return
	fmt.Println("value of randSeq, ", randRet)

	//Init User in Database
	initUserChannel := make(chan string, 1)
	data.InitUser(uR.Email, uR.Password, randRet, initUserChannel)
	initUserRet := <-initUserChannel //blocking return
	fmt.Println("value of initUserRet, ", initUserRet)

	if initUserRet == "some error occurred"{
		CookiesMain("", "", "login", w)
		util.HandleLoginResponse(w, "Error", "Register", "Some Error Occurred", "")
	}else if initUserRet == "user already added"{
		CookiesMain("", "", "login", w)
		util.HandleLoginResponse(w, "Error", "Register", "Email Already Registered", "")
	}else{
		CookiesMain(uR.Email, uR.Password, "login", w)
		tokenString, _ := CreateJWT(uR.Email)
		util.HandleLoginResponse(w, "Success", "Register", "Email Registered", tokenString)
	}

}