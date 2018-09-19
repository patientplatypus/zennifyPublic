package reqt

import (
	//standard library
	"fmt"
	"net/http"
	"encoding/json"
	//nonstandard libraries
	// "github.com/gorilla/mux"
	//project files
	"github.com/patientplatypus/webserver/utilities"
	"github.com/patientplatypus/webserver/database"
)

func GetUsers(w http.ResponseWriter, req *http.Request){
	fmt.Println("inside NewHashMsg")

	usersChannel := make(chan []string)
	go data.Retrieve_users(usersChannel)
	userRet := <-usersChannel //blocking return
	fmt.Println("value of userRet, ", userRet);

	util.HandleRequestResponseArray(w, "UsersList", userRet)
}

func SendMsg(w http.ResponseWriter, req *http.Request){
	fmt.Println("inside SendMsg")

	msgChannel := make(chan string)
	go data.Write_msg_table(msgChannel, req)
	msgRet := <-msgChannel //blocking return
	fmt.Println("value of msgRet, ", msgRet);

	util.HandleRequestResponse(w, "MsgSent", "Success")
}

func GetMail(w http.ResponseWriter, req *http.Request){
	fmt.Println("inside GetMail")

	mailChannel := make(chan util.ReturnMessages)
	go data.Retrieve_mail(mailChannel, req)
	mailRet := <-mailChannel //blocking return
	fmt.Println("value of mailRet, ", mailRet);

	json.NewEncoder(w).Encode(mailRet)	

	// util.HandleRequestResponse(w, "Success", "Success")
}

func GetMsg(w http.ResponseWriter, req *http.Request){
	fmt.Println("inside GetMsg")

	mailMsgChannel := make(chan util.ReturnMsg)
	go data.Retrieve_msg(mailMsgChannel, req)
	msgRet := <- mailMsgChannel //blocking return
	fmt.Println("value of msgRet, ", msgRet);

	json.NewEncoder(w).Encode(msgRet)	

	// util.HandleRequestResponse(w, "Success", "Success")
}

