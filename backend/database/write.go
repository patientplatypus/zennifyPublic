package data

import(
	// standard library
	"fmt"
	"log"
	"time"
	"strconv"
	"net/http"
	"encoding/json"
	// "database/sql"
	// other libraries
	"github.com/lib/pq"
	// "github.com/gocql/gocql"
	// "golang.org/x/crypto/bcrypt"
	"github.com/patientplatypus/webserver/utilities"
)

type expenseReq struct{
	Type string `json:"type"`
	Expense int `json:"expense"`
	StartDate string `json: "startDate"`
	EndDate string `json:"endDate"`
	DateArray []string `json:"dateArray"`
	Interest float64 `json:"interest"`
	Compounded string `json:"compounded"`
	Recurring string `json:"recurring"`
	Email string `json:"email"`
	Name string `json:"name"`
	Description string `json:"description"`
	ExpenseOwedTo string `json:"expenseOwedTo"`
	JWT string `json:"localJWT"`
}

type profitReq struct {
	Type string `json:"type"`
	Profit int `json:"profit"`
	StartDate string `json: "startDate"`
	EndDate string `json:"endDate"`
	DateArray []string `json:"dateArray"`
	Interest float64 `json:"interest"`
	Compounded string `json:"compounded"`
	Recurring string `json:"recurring"`
	Email string `json:"email"`
	Name string `json:"name"`
	Description string `json:"description"`
	ProfitFrom string `json:"profitFrom"`
	JWT string `json:"localJWT"`
}

type newChannelReq struct{
	ChannelName string `json:"channelName"`
	ChannelDescription string `json:"channelDescription"`
}

type newSubChannelReq struct{
	ChannelName string `json:"channelName"`
	SubChannelName string `json:"subChannelName"`
	SubChannelDescription string `json:"subChannelDescription"`
}

type newHashReq struct{
	Email string `json:"email"`
	HashName string `json:"hashName"`
	HashWrite string `json:"hashWrite"`
}

type sendMsgReq struct{
	MessageFrom string `json:"messageFrom"`
	MessageTo string `json:"messageTo"`
	MessageContent string `json:"messageContent"`
	MessageSubject string `json:"messageSubject"`
}

func Write_expense_table(expenseWriteChannel chan string, expenseType string, req *http.Request){
	var expenseReq expenseReq
	err := json.NewDecoder(req.Body).Decode(&expenseReq)
	if err!=nil{
		fmt.Println("Error in decoding expenseReq in Write_expense_table");
		log.Panic(err)
	}else{
		fmt.Println("Successfully decode expenseReq in Write_expense_table");
	}

	NotInDB, _ := Search_userinfo_table(expenseReq.Email)

	if NotInDB == false {

		fmt.Println("inside NotInDB === false")
		fmt.Println("value of expenseType: ", expenseType);

		if expenseType == "oneTime"{
			fmt.Println("inside oneTime if in Write_expense_table")
			var lastInsertId int
			err2 := db.QueryRow("INSERT INTO profit(email, type, dateArray, amount, interest, compounded, recurring, name, description, profitFrom) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) returning uid;", expenseReq.Email, expenseReq.Type, pq.Array(expenseReq.DateArray), expenseReq.Expense, expenseReq.Interest, expenseReq.Compounded, expenseReq.Recurring, expenseReq.Name, expenseReq.Description, expenseReq.ExpenseOwedTo).Scan(&lastInsertId);
			if err2!=nil{
				fmt.Println("Error in inserting expense into db")
				log.Panic(err2)
			}
		}else if expenseType == "recurring"{
			// var lastInsertId int
			fmt.Println("inside recurring if in Write_expense_table")
			recurringDateArray := util.Parse_Recurring(expenseReq.StartDate, expenseReq.EndDate, expenseReq.Recurring);
			fmt.Println("back from Parse_Recurring and value of recurringDateArray: ")
			fmt.Println(recurringDateArray)
			// err3 := db.QueryRow("INSERT INTO profit(email, type, dateArray, amount, interest, compounded, recurring, name, description, profitFrom) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) returning uid;", expenseReq.Email, expenseReq.Type, pq.Array(expenseReq.DateArray), expenseReq.Expense, expenseReq.Interest, expenseReq.Compounded, expenseReq.Recurring, expenseReq.Name, expenseReq.Description, expenseReq.ExpenseOwedTo).Scan(&lastInsertId);
			// if err3!=nil{
			// 	fmt.Println("Error in inserting expense into db")
			// 	log.Panic(err2)
			// }
		}
	}else{
		fmt.Println("not ok to add expense to db; email not found!")
	}


	expenseWriteChannel<-"finished writing to profit"

}





func Write_profit_table(profitWriteChannel chan string, profitType string, req *http.Request) {
	var profitReq profitReq
	err := json.NewDecoder(req.Body).Decode(&profitReq)
	if err!=nil{
		fmt.Println("Error in decoding profitReq in Write_profit_table")
		log.Panic(err)
	}else{
		fmt.Println("Successfully decoded profitReq in Write_profit_table")
	}

	NotInDB, _ := Search_userinfo_table(profitReq.Email)

	if NotInDB == false {	
		var lastInsertId int
		err2 := db.QueryRow("INSERT INTO profit(email, type, dateArray, amount, interest, compounded, recurring, name, description, profitFrom) VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10) returning uid;", profitReq.Email, profitReq.Type, pq.Array(profitReq.DateArray), profitReq.Profit, profitReq.Interest, profitReq.Compounded, profitReq.Recurring, profitReq.Name, profitReq.Description, profitReq.ProfitFrom).Scan(&lastInsertId);
		if err2!=nil{
			fmt.Println("Error in inserting profit into db")
			log.Panic(err2)
		}
	}else{
		fmt.Println("not ok to add profit to db; email not found!")
	}

	searchProfitbool := Search_profit_table(profitReq.Email)

	// Search_variable_table("profit");

	if searchProfitbool == false{
		fmt.Println("searchProfitbool is false")
	}else{
		fmt.Println("searchProfitbool is true")
	}

	profitWriteChannel<-"finished writing to profit"
}


func Write_channel_table(newTalkChannelWriteChannel chan string, req *http.Request){
	fmt.Println("inside Write_channel_table")
	var newChannelReq newChannelReq
	_ = json.NewDecoder(req.Body).Decode(&newChannelReq)
	OKtoADD := Search_channel_table(newChannelReq.ChannelName)
	if OKtoADD == true && newChannelReq.ChannelName!="feed"{
		var lastInsertId int
		err2 := db.QueryRow("INSERT INTO channel(channelName, channelDescription) VALUES($1, $2) returning uid;", newChannelReq.ChannelName, newChannelReq.ChannelDescription).Scan(&lastInsertId)
		if err2!=nil{
			fmt.Println("Error in inserting channel into db")
			log.Panic(err2)
		}
	}else{
		fmt.Println("not ok to add channel to db; already exists")
	}
	newTalkChannelWriteChannel<-"return from Write_channel_table"
}

func Write_subchannel_table(newTalkSubChannelWriteChannel chan string, req *http.Request){
	fmt.Println("inside Write_channel_table")
	var newSubChannelReq newSubChannelReq
	_ = json.NewDecoder(req.Body).Decode(&newSubChannelReq)
	OKtoADD := Search_subchannel_table(newSubChannelReq.SubChannelName)
	if OKtoADD == true{
		var lastInsertId int
		err2 := db.QueryRow("INSERT INTO subchannel(channelName, subChannelName, subChannelDescription) VALUES($1, $2, $3) returning uid;", newSubChannelReq.ChannelName, newSubChannelReq.SubChannelName, newSubChannelReq.SubChannelDescription).Scan(&lastInsertId)
		if err2!=nil{
			fmt.Println("Error in inserting subchannel into db")
			log.Panic(err2)
		}
	}else{
		fmt.Println("not ok to add subchannel to db; already exists")
	}
	newTalkSubChannelWriteChannel<-newSubChannelReq.ChannelName
}

func Write_hash_table(hashWriteChannel chan string, req *http.Request){
	fmt.Println("inside Write_hash_table")
	var newHashReq newHashReq
	_ = json.NewDecoder(req.Body).Decode(&newHashReq)
	var newHashArray []string
	newHashArray = append(newHashArray, newHashReq.HashWrite)
	fmt.Println("value of newHashReq.HashWrite: ", newHashReq.HashWrite)
	dbExecString := ""
	if newHashReq.HashName=="profileCanvas"{
		dbExecString = "INSERT INTO ipfshash(email, profileCanvas) VALUES($1, $2) ON CONFLICT (email) DO UPDATE SET profileCanvas = EXCLUDED.profileCanvas"
	}else if newHashReq.HashName=="profileMsgArray"{
		dbExecString = "INSERT INTO ipfshash(email, profileMsgArray) VALUES($1, $2) ON CONFLICT (email) DO UPDATE SET profileMsgArray = coalesce(ipfshash.profileMsgArray, '{}') || EXCLUDED.profileMsgArray"
	}


	fmt.Println("value of dbExecString: ", dbExecString)
	OKtoADD, _ := Search_userinfo_table(newHashReq.Email)
	if OKtoADD == false{
		fmt.Println("value of newHashArray: ")
		fmt.Println(newHashArray)
		fmt.Println("value of newHashReq.HashWrite")
		fmt.Println(newHashReq.HashWrite)
		_, err := db.Exec(dbExecString, newHashReq.Email, pq.Array(newHashArray))
		if err != nil {
			log.Panic(err)
		}	
	}else{
		fmt.Println("not ok to add subchannel to db; already exists")
	}
	hashWriteChannel<-"done writing hash"
}

func Write_msg_table(msgChannel chan string, req *http.Request){
	fmt.Println("inside Write_msg_table")

	var sendMsgReq sendMsgReq
	_ = json.NewDecoder(req.Body).Decode(&sendMsgReq)

	fmt.Println("sendMsgReq.MessageTo")
	fmt.Println(sendMsgReq.MessageTo)
	fmt.Println("sendMsgReq.MessageFrom")
	fmt.Println(sendMsgReq.MessageFrom)
	fmt.Println("sendMsgReq.MessageContent")
	fmt.Println(sendMsgReq.MessageContent)
	fmt.Println("sendMsgReq.MessageSubject")
	fmt.Println(sendMsgReq.MessageSubject)

	util.CassSession, _ = util.CassCluster.CreateSession()
	defer util.CassSession.Close()
	keySpaceMeta, _ := util.CassSession.KeyspaceMetadata("platypus")
	valC, exists := keySpaceMeta.Tables["cassmessage"]

	if exists==true {
		fmt.Println("cassmessage exists!!!")
	}else{
		fmt.Println("cassmessage doesnt exist!")
		errCT :=util.CassSession.Query("CREATE TABLE cassmessage(" +
		"messageto text, messagefrom text, messagecontent text, PRIMARY KEY (messageto))").Exec()
		if errCT != nil{
			fmt.Println("there was an error creating the table: ", errCT)
		}
		fmt.Println("created the table cassmessage")
	}

	if valC!=nil{
		fmt.Println("return from valC cassmessage: ", valC)
	}

	updateString:=`UPDATE cassmessage SET messagefrom='`+sendMsgReq.MessageFrom+`', messagecontent='`+sendMsgReq.MessageContent+`' WHERE messageto='`+sendMsgReq.MessageTo+`'`

	// timeString, _ := strconv.ParseInt(time.Now().UnixNano(), 10, 64)

	timeString := strconv.FormatInt(time.Now().UnixNano(), 10)

	insertString:=`INSERT INTO cassmessage (messagefrom, messageto, messagecontent, messagesubject, uniqueID) VALUES('`+sendMsgReq.MessageFrom+`', '`+sendMsgReq.MessageTo+`', '`+sendMsgReq.MessageContent+`', '` +sendMsgReq.MessageSubject+`', '`+timeString+`')`

	fmt.Println("updateString value: ", updateString)
	fmt.Println("insertString value: ", insertString)

	err := util.CassSession.Query(insertString).Exec()

	if err != nil {
		fmt.Println("there was an error in appending data to cassmessage: ", err)
	} else {
		fmt.Println("inserted data into cassmessage successfully")
	}
	msgChannel<-"done"
}