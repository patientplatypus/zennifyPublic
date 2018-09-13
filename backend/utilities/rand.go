package util


import(
	"math/rand"
	"time"
)


func genSeed(){
	rand.Seed(time.Now().UnixNano())
}

var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func randSeq(n int) string {
    b := make([]rune, n)
    for i := range b {
        b[i] = letters[rand.Intn(len(letters))]
    }
    return string(b)
}

func RandMain(randFuncChannel chan string){
	genSeed()
	randFuncChannel <- randSeq(32)
}