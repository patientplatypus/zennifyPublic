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


type getHashRet struct{
	Status string
	Message string
	Hashes []string
}

type newHashRet struct{
	Status string
	Message string
}


func NewHashMsg(w http.ResponseWriter, req *http.Request){
	fmt.Println("inside NewHashMsg")

	hashWriteChannel := make(chan string)
	go data.Write_hash_table(hashWriteChannel, req)
	hashRet := <-hashWriteChannel //blocking return
	fmt.Println("value of hashRet, ", hashRet);

	// util.HandleRequestResponse(w, "Success", hashRet, JWT: "")
	response := newHashRet{Status: "Success", Message: hashRet}
	json.NewEncoder(w).Encode(response)	
}

func GetHashMsg(w http.ResponseWriter, req *http.Request){
	fmt.Println("inside GetHashMsg")

	hashRetrieveChannel := make(chan util.HashRetChan)
	go data.Retrieve_ipfshash(hashRetrieveChannel, req)
	hashRet := <-hashRetrieveChannel //blocking return
	fmt.Println("value of hashRet, ", hashRet);

	response := getHashRet{Status: "Success", Message: hashRet.Name, Hashes: hashRet.Value}
	json.NewEncoder(w).Encode(response)	
}
