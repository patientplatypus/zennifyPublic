package reqt

import (
	//standard library
	"fmt"
	"net/http"
	// "encoding/json"
	//nonstandard libraries
	//project files
	"github.com/patientplatypus/webserver/utilities"
)

func TestJWT(w http.ResponseWriter, req *http.Request){
	fmt.Println("inside TestJWT")
	util.HandleRequestResponse(w, "Success", "inside TestJWT")
}