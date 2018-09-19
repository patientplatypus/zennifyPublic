package main

import (
	// standard library
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
	"bytes"
	"strings"
	"flag"
	// "path/filepath"
	// "context"
	// "io/ioutil"
	// other libraries
	"github.com/rs/cors"
	"github.com/gorilla/mux"
	// "github.com/dgrijalva/jwt-go"
	// project files
	"github.com/patientplatypus/webserver/authentication"
	"github.com/patientplatypus/webserver/database"
	"github.com/patientplatypus/webserver/requests"
	"github.com/patientplatypus/webserver/utilities"
)

type ExampleResponse struct {
	Status string
}

type clientJWT struct {
	ClientJWT string `json:"localJWT"`
}

func JWTHandler(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		fmt.Println("inside JWTHandler")
		fmt.Println("value of req.URL.Path: ", req.URL.Path)
		if req.URL.Path == "/registerUser" || req.URL.Path == "/loginUser" || req.URL.Path == "/cookieLogin"{
			fmt.Println("inside JWTHandler handle Auth")
			next.ServeHTTP(w,req)
		}else if req.URL.Path == "/upload"{
			fmt.Println("uploading file don't know how to validate multiform!")
			next.ServeHTTP(w,req)
		}else if strings.Contains(req.URL.Path, "files"){
			fmt.Println("hello there files")
			next.ServeHTTP(w,req)
		}else{
			fmt.Println("inside JWTHandler authing routes")
			var cJWT clientJWT
			fmt.Println("test line 1")

			buf, _ := ioutil.ReadAll(req.Body)
			rdr1 := ioutil.NopCloser(bytes.NewBuffer(buf))
			rdr2 := ioutil.NopCloser(bytes.NewBuffer(buf))

			fmt.Println("test line 2")
			
			err := json.NewDecoder(rdr1).Decode(&cJWT)
			if err != nil {
				fmt.Println("error in decodeing localJWT in JWTHandler")
				http.Error(w, err.Error(), 400)
				return
			}

			fmt.Println("test line 3")

			req.Body = rdr2 // OK since rdr2 implements the io.ReadCloser interface

			fmt.Println("test line 4")

			fmt.Println("cJWT.ClientJWT: ")

			fmt.Println(cJWT.ClientJWT)

			validateJWTChan := make(chan bool)
			go auth.ValidateJWT(cJWT.ClientJWT, false, validateJWTChan)
			isAuthenticated := <- validateJWTChan
			if isAuthenticated == true {
				fmt.Println("inside JWTHandler isAuthenticated true")
				fmt.Println("value of req.URL.Path: ", req.URL.Path)
				next.ServeHTTP(w, req)
			}else{
				fmt.Println("inside JWTHandler isAuthenticated false")
				util.HandleRequestResponse(w, "Error", "Not Authenticated")
			}
		}
	})
}

func main() {
	var dir string
	flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	
	fmt.Println("value of dir")
	fmt.Println(dir)

	data.InitDB()

	auth.RotateJWTkey()

	fmt.Println("we get past RotateJWTkey function call !!!")

	r := mux.NewRouter()



	r.HandleFunc("/testJWT", reqt.TestJWT)
	r.HandleFunc("/registerUser", auth.RegisterUser)
	r.HandleFunc("/loginUser", auth.LoginUser)
	r.HandleFunc("/cookieLogin", auth.CookieLogin)

	s := r.PathPrefix("/moneySubmit").Subrouter()
	s.HandleFunc("/profit", reqt.ProfitSubmit)
	s.HandleFunc("/expense/{expenseType}", reqt.ExpenseSubmit)

	t := r.PathPrefix("/talk").Subrouter()
	t.HandleFunc("/newChannel", reqt.NewChannel)
	t.HandleFunc("/getChannels", reqt.GetChannels)
	t.HandleFunc("/newSubChannel", reqt.NewSubChannel)
	t.HandleFunc("/getSubChannels", reqt.GetSubChannels)

	h := r.PathPrefix("/hash").Subrouter()
	h.HandleFunc("/newHashMsg", reqt.NewHashMsg)
	h.HandleFunc("/getHashMsg", reqt.GetHashMsg)

	
	r.HandleFunc("/upload", util.UploadFileHandler())
	r.HandleFunc("/files/{filename}", util.ServeFiles)

	u := r.PathPrefix("/users").Subrouter()
	u.HandleFunc("/getUsers", reqt.GetUsers)
	u.HandleFunc("/sendMsg", reqt.SendMsg)
	u.HandleFunc("/getMail", reqt.GetMail)
	u.HandleFunc("/getMsg", reqt.GetMsg)


	JWTMux := JWTHandler(r)
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:8081"},
		AllowCredentials: true,
	})
	handler := c.Handler(JWTMux)
	log.Fatal(http.ListenAndServe(":8000", handler))

}
