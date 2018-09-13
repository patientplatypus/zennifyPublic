package data

import(
	// standard library
	"fmt"
	"log"
	"database/sql"
	// other libraries
	_ "github.com/lib/pq"
	// "github.com/gocql/gocql"
	// "golang.org/x/crypto/bcrypt"
)

var db *sql.DB


func InitDB() {
	var err error

	db, err = sql.Open("postgres", "host=db port=5432 user=docker password=docker dbname=zennify sslmode=disable")	
	fmt.Println("value of db: ", db);

    if err != nil {
        log.Panic(err)
    }else{
		fmt.Println("db init successful")
		create_cassandra()
		cassandra_message()
		create_userinfo_table()	
		create_profit_table()
		create_expense_table()
		create_channel_table()
		create_subchannel_table()
		create_ipfshash_table()
	}
}

func InitUser(email string, password string, regString string, initUserChannel chan string){
	var lastInsertId int
	OKtoAdd, _ := Search_userinfo_table(email)
	fmt.Println("value of OKtoAdd, %t", OKtoAdd)
	if OKtoAdd == true{
		hash, _ := HashPassword(password) // ignore error for the sake of simplicity
		err := db.QueryRow("INSERT INTO userinfo(email, password, regString, regBool) VALUES($1, $2, $3, $4) returning uid;", email, hash, regString, true).Scan(&lastInsertId)
		if err != nil {
			log.Panic(err)
			initUserChannel<-"some error occurred"
		}else{
			fmt.Println("last inserted id =", lastInsertId)
			initUserChannel<-hash
		}
	}else{
		initUserChannel<-"user already added"
	}
}