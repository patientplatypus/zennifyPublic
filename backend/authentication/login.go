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

type userLogin struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

func cookieStringConvert(cookieRaw *http.Cookie,) string {
	return cookieRaw.Value
}


func CookieLogin(w http.ResponseWriter, req *http.Request) {
	fmt.Println("inside CookieLogin")
	cookieEmail, errEmail := req.Cookie("zennifyEmail")
	cookiePassword, errPassword:= req.Cookie("zennifyPassword")
	if errEmail != nil || errPassword != nil{
		fmt.Println("cookieEmail/cookieEmail error not nil handler")
		util.HandleLoginResponse(w, "Error", "Login", "Cookies Not Present", "")	
	}else{
		cookieEmailString := cookieStringConvert(cookieEmail)
		cookiePasswordString := cookieStringConvert(cookiePassword)
		loginHandler(cookieEmailString, cookiePasswordString, w)
	}
}

func LoginUser(w http.ResponseWriter, req *http.Request) {
	fmt.Printf("inside loginUser in Golang")
	
	var uL userLogin
	err := json.NewDecoder(req.Body).Decode(&uL)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	loginHandler(uL.Email, uL.Password, w)
}

func loginHandler(email string, password string, w http.ResponseWriter){

	EmailNotInDatabase, savedHashedPassword := data.Search_userinfo_table(email)

	if EmailNotInDatabase == true{
		util.HandleLoginResponse(w, "Error", "Login", "Email Not In Database", "")		
	}else{
		hashEqual := data.CheckPasswordHash(password, savedHashedPassword)
		if hashEqual{
			tokenString, _ := CreateJWT(email)
			CookiesMain(email, password, "login", w)
			util.HandleLoginResponse(w, "Success", "Login", "Email and Password Match", tokenString)
		}else{
			util.HandleLoginResponse(w, "Error", "Login", "Password Not Correct", "")	
		}
	}
}