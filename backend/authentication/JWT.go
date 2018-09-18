package auth

import (
	//standard libraries
	// "time"
	"reflect"
	"fmt"
	// nonstandard libraries
	"github.com/dgrijalva/jwt-go"
	// project files
	"github.com/patientplatypus/webserver/database"
	"github.com/patientplatypus/webserver/utilities"
)

var JWTkey = []byte("7kShHNdi3rLhtmxYuforGuBNYT49V337")
var JWTkeyPrev = []byte("7kShHNdi3rLhtmxYuforGuBNYT49V337")

func CreateJWT(email string)(string, error){
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"email": email, "nbf": 1000})

	// Sign and get the complete encoded token as a string using the secret
	// this should probably be rotated on a chron job
	tokenString, err := token.SignedString(JWTkey)
	fmt.Println("format of tokenString in createjwt")
	fmt.Println(reflect.TypeOf(tokenString))

	fmt.Println(tokenString, err)
	return tokenString, err
}

func keyRotator(tokenString string, JWTkeyLocal []byte)(*jwt.Token, error){
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return JWTkeyLocal, nil
	})
	return token, err
}

func ValidateJWT(tokenString string, prevKey bool, validateJWTChan chan bool){
	fmt.Println("inside ValidateJWT function")
	JWTkeyLocal := []byte("")
	if prevKey == false{
		JWTkeyLocal = JWTkey
	}else if prevKey == true{
		JWTkeyLocal = JWTkeyPrev
	}

	token, _ := keyRotator(tokenString, JWTkeyLocal)
	
	fmt.Println("value of token")
	fmt.Println(token)

	claims, ok := token.Claims.(jwt.MapClaims);
	fmt.Println("value of claims: ")
	fmt.Println(claims)
	fmt.Println("value of OK")
	fmt.Println(ok)
	if ok && token.Valid {
		fmt.Println(claims["email"], claims["nbf"])
		UserNotInDatabase, _ := data.Search_userinfo_table(claims["email"].(string))
		if UserNotInDatabase == false{
			if prevKey == true{
				newJWT, _ := CreateJWT(claims["email"].(string))
				util.UpdateJWT = newJWT
			}
			validateJWTChan<-true
		}else{
			validateJWTChan<-false
		}
	} else if prevKey == false{
		ValidateJWT(tokenString, true, validateJWTChan)
	} else if prevKey == true{
		validateJWTChan<-false
	}
}