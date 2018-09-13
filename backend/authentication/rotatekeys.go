package auth 

import (
	//standard library
	"time"
	// "reflect"
	"fmt"
	//nonstandard libraries
	//project files
	// "github.com/jasonlvhit/gocron"
	"github.com/patientplatypus/webserver/utilities"
)

func rotateKeyFunc(){
	fmt.Println("inside rotateKeyFunc")
	keyChan := make(chan string)
	go util.RandMain(keyChan)
	
	Loop:
		for {
			select{
			case randRet := <-keyChan:
				JWTkeyPrev = JWTkey
				JWTkey = []byte(randRet)
				fmt.Println("new value of JWTkey")
				fmt.Println(JWTkey)
				break Loop
			default:
				fmt.Println("still making key")
				time.Sleep(1 * time.Second)
			}
		}
}

func RotateJWTkey(){
	fmt.Println("inside RotateJWTkey in rotatekeys")
	ticker := time.NewTicker(60 * time.Minute)
	go func() {
        for t := range ticker.C {
			fmt.Println("Tick at", t)
			
			keyChan := make(chan string)
			go util.RandMain(keyChan)
			randRet := <-keyChan

			JWTkeyPrev = JWTkey
			JWTkey = []byte(randRet)
			fmt.Println("new value of JWTkey")
			fmt.Println(JWTkey)
		}
    }()

	// s := gocron.NewScheduler()
	// s.Every(5).Minutes().Do(rotateKeyFunc)
	// <- s.Start()

}