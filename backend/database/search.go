package data

import(
	// standard library
	"fmt"
	"log"
	// "database/sql"
	// other libraries
	"github.com/lib/pq"
	// "golang.org/x/crypto/bcrypt"
)


func join(strs ...string) string {
	var ret string
	for _, str := range strs {
		ret += str
	}
	return ret
}

func Search_variable_table(tableName string){
	fmt.Println("search_variable_table started")
	rows, _ := db.Query("SELECT * FROM " + tableName)
	columns, _ := rows.Columns()
    count := len(columns)
    values := make([]interface{}, count)
	valuePtrs := make([]interface{}, count)
	
	for rows.Next() {
        for i, _ := range columns {
            valuePtrs[i] = &values[i]
        }
        rows.Scan(valuePtrs...)
		
        for i, col := range columns {
            var v interface{}
            val := values[i]
            b, ok := val.([]byte)
            if (ok) {
                v = string(b)
            } else {
                v = val
			}
			fmt.Println("printing from search_variable_table: ", tableName, ": ",col, v)
			if col == "uid"{
				fmt.Println("---------------------------")
			}
        }
    }
}

func Search_profit_table(searchEmail string)(bool){
	fmt.Println("search_profit_table started")
	fmt.Println("value of searchEmail: ", searchEmail)
	rows, err1 := db.Query("SELECT * FROM profit")
	if err1 != nil {
		log.Panic(err1)
	}

	defer rows.Close()
	for rows.Next() {	
		var email string
		var profitType string
		var dateArray []string
		var profit int
		var interest float64
		var compounded string
		var recurring string
		var name string
		var description string
		var profitFrom string
		var uid int
		err2 := rows.Scan(&email, &profitType, pq.Array(&dateArray), &profit, &interest, &compounded, &recurring, &name, &description, &profitFrom, &uid)
		if err2 != nil {
			log.Panic(err2)
		}
		fmt.Println("email | profit")
		fmt.Printf("%s | %s", email, profit)
		// if email == searchEmail{
		// 	return false
		// }
	}
	return true
}


func Search_userinfo_table(searchEmail string)(bool, string) {
	fmt.Println("search_userinfo_table started")
	rows, err1 := db.Query("SELECT * FROM userinfo")
	if err1 != nil {
		log.Panic(err1)
	}
	defer rows.Close()
	for rows.Next() {
		var email string
		var password string
		var regString string
		var regBool bool
		var uid int
		err2 := rows.Scan(&email, &password, &regString, &regBool, &uid)
		if err2 != nil {
			log.Panic(err2)
		}
		fmt.Println("email | password")
		fmt.Printf("%s | %s", email, password)
		if email == searchEmail{
			return false, password
		}
	}
	return true, ""
}


func Search_channel_table(searchName string)(bool){
	fmt.Println("search_channel_table started")
	rows, err1 := db.Query("SELECT * FROM channel")
	if err1 != nil {
		log.Panic(err1)
	}
	defer rows.Close()
	for rows.Next() {
		var channelName string
		var channelDescription string
		var uid int
		err2 := rows.Scan(&channelName, &channelDescription, &uid)
		if err2 != nil {
			log.Panic(err2)
		}
		if (searchName == channelName){
			return false
		}
	}
	return true;
}

func Search_subchannel_table(searchName string)(bool){
	fmt.Println("search_subchannel_table started")
	rows, err1 := db.Query("SELECT * FROM subchannel")
	if err1 != nil {
		log.Panic(err1)
	}
	defer rows.Close()
	for rows.Next() {
		var channelName string
		var subChannelName string
		var subChannelDescription string
		var uid int
		err2 := rows.Scan(&channelName, &subChannelName, &subChannelDescription, &uid)
		if err2 != nil {
			log.Panic(err2)
		}
		if (searchName == subChannelName){
			return false
		}
	}
	return true;
}