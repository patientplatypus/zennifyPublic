package data

import(
	// standard library
	"fmt"
	"log"
	// "time"
	"net/http"
	"sort"
	"encoding/json"
	// "database/sql"
	// other libraries
	"github.com/lib/pq"
	// "github.com/gocql/gocql"
	// "golang.org/x/crypto/bcrypt"
	"github.com/patientplatypus/webserver/utilities"
)

type ChannelInfo struct{
	ChannelName string
	ChannelDescription string
}

type SubChannelInfo struct{
	ChannelName string
	SubChannelName string
	SubChannelDescription string
}

type StringSlice []string

type getHashReq struct{
	Email string `json:"email"`
	HashName string `json:"hashName"`
}

type mailReq struct{
	Email string `json:"email"`
}

type msgReq struct {
	ID string `json:"ID"`
}


func Retrieve_users(usersChannel chan []string, )(){
	fmt.Println("retrieve_users func started")
	rows, err1 := db.Query("SELECT * FROM userinfo")
	var returnEmailList []string
	if err1 != nil {
		log.Panic(err1)
	}
	defer rows.Close()
	for rows.Next() {
		var email string
		var password string
		var regString string
		var regBool string
		var uid int
		err2 := rows.Scan(&email, &password, &regString, &regBool, &uid)
		if err2 != nil {
			log.Panic(err2)
		}
		returnEmailList = append(returnEmailList, email)
	}
	usersChannel<-returnEmailList
}



func Retrieve_channels()([]ChannelInfo){
	fmt.Println("retrieve_channels func started")

	channelData := []ChannelInfo{}

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
		fmt.Println("ChannelName | ChannelDescription")
		fmt.Println(channelName, channelDescription)
		channelData = append(channelData, ChannelInfo{ChannelName: channelName, ChannelDescription: channelDescription})
	}
	return channelData
}

func Retrieve_subchannels(channelSearchName string)([]SubChannelInfo){
	fmt.Println("retrieve_subchannels func started")
	fmt.Println("value of channelSearchName")
	fmt.Println(channelSearchName)

	subChannelData := []SubChannelInfo{}

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
		if channelSearchName == channelName{
			subChannelData = append(subChannelData, SubChannelInfo{ChannelName: channelName, SubChannelName: subChannelName, SubChannelDescription: subChannelDescription})
		}
	}
	return subChannelData
}

func Retrieve_ipfshash(hashRetrieveChannel chan util.HashRetChan, req *http.Request){
	var email string
	var profileMsgArray []string
	var profileCanvas []string
	fmt.Println("inside retrieve_ipfshash")
	var getHashReq getHashReq
	_ = json.NewDecoder(req.Body).Decode(&getHashReq)
	OKtoADD, _ := Search_userinfo_table(getHashReq.Email)
	if OKtoADD == false{
		// dbExecString := "SELECT * FROM ipfshash WHERE email = '" + getHashReq.Email + "'"
		dbExecString := "SELECT * FROM ipfshash"
		fmt.Println("value of dbExecString")
		fmt.Println(dbExecString)
		rows, err1 := db.Query(dbExecString)
		if err1 != nil {
			log.Panic(err1)
		}
		defer rows.Close()
		for rows.Next(){
			fmt.Println("iterating in Retrieve_ipfshash rows.Next")
			err := rows.Scan(&email, pq.Array(&profileMsgArray), pq.Array(&profileCanvas))
			fmt.Println("value of email: ", email)
			fmt.Println("value of profileCanvas: ", profileCanvas)
			if err != nil {
				fmt.Println("in panic for retrieve_ipfshash")
				log.Panic(err)
			}
		}
		if getHashReq.HashName == "profileMsgArray"{
			retReturn:=util.HashRetChan{Name: "profileMsgArray", Value: profileMsgArray}
			hashRetrieveChannel<-retReturn
		}else{
			retReturn:=util.HashRetChan{Name: "profileCanvas", Value: profileCanvas}
			hashRetrieveChannel<-retReturn
		}
		
	}
}

func Retrieve_mail(mailChannel chan util.ReturnMessages, req *http.Request)(){
	fmt.Println("inside Retrieve_mail")

	var mailReq mailReq
	_ = json.NewDecoder(req.Body).Decode(&mailReq)

	util.CassSession, _ = util.CassCluster.CreateSession()
	defer util.CassSession.Close()
	keySpaceMeta, _ := util.CassSession.KeyspaceMetadata("platypus")
	valC, exists := keySpaceMeta.Tables["cassmessage"]
	fmt.Println(valC, exists)

	// queryString := `SELECT messageto, messagecontent, messagefrom FROM cassmessage WHERE messagefrom='`+mailReq.Email+`'`//returns nothing, should return many rows
	// queryString2 := `SELECT messageto, messagecontent, messagefrom FROM cassmessage`//returns only last entry, should return many rows
	// queryString3 := `SELECT * FROM cassmessage WHERE messagefrom='`+mailReq.Email+`'`//returns nothing, should return many rows
	// queryAllString := `SELECT * FROM platypus.cassmessage` //returns only last entry, should return many rows

	var messageto string
	var messagesubject string
	var messagefrom string
	var uniqueID string

	var newSentMail []util.Mail
	var newReceivedMail []util.Mail

	iterC := util.CassSession.Query(`SELECT messageto, messagefrom, messagesubject, uniqueID FROM cassmessage`).Iter()
    for iterC.Scan(&messageto, &messagefrom, &messagesubject, &uniqueID) {
		fmt.Println("value of messageto: ", messageto)
		fmt.Println("value of messagefrom: ", messagefrom)
		fmt.Println("value of uniqueID: ", uniqueID)
		if messageto == mailReq.Email{
			newMail := util.Mail{Header: messageto, Subject:messagesubject, ID: uniqueID}
			fmt.Println("value of newMail: ", newMail)
			newReceivedMail = append(newReceivedMail, newMail)
		}else if messagefrom == mailReq.Email{
			newMail := util.Mail{Header: messageto, Subject:messagesubject, ID: uniqueID}
			fmt.Println("value of newMail: ", newMail)
			newSentMail = append(newSentMail, newMail)
		}
	}

	fmt.Println("after iterC Scan for and value of newSentMail")
	fmt.Println(newSentMail)
	fmt.Println("after iterC Scan for and value of newReceivedMail")
	fmt.Println(newReceivedMail)

	sort.Slice(newSentMail, func(i, j int) bool {return newSentMail[i].ID > newSentMail[j].ID}) 
	sort.Slice(newReceivedMail, func(i, j int) bool {return newReceivedMail[i].ID > newReceivedMail[j].ID}) 

	fmt.Println("after iterC Scan for and value of newSentMail AFTER SORT")
	fmt.Println(newSentMail)
	fmt.Println("after iterC Scan for and value of newReceivedMail AFTER SORT")
	fmt.Println(newReceivedMail)


	messagesReturned := util.ReturnMessages{SentMail: newSentMail, ReceivedMail: newReceivedMail}

	fmt.Println("value of newMessagesToReturn")
	fmt.Println(messagesReturned)

	
	mailChannel <- messagesReturned
}

func Retrieve_msg(mailMsgChannel chan util.ReturnMsg, req *http.Request){
	fmt.Println("inside Retrieve_msg")
	var msgReq msgReq
	_ = json.NewDecoder(req.Body).Decode(&msgReq)

	util.CassSession, _ = util.CassCluster.CreateSession()
	defer util.CassSession.Close()
	keySpaceMeta, _ := util.CassSession.KeyspaceMetadata("platypus")
	valC, exists := keySpaceMeta.Tables["cassmessage"]
	fmt.Println(valC, exists)

	var messagecontent string
	var returnMsg util.ReturnMsg
	fmt.Println("value of msgReq.ID: ", msgReq.ID)
	queryString := `SELECT messagecontent from cassmessage WHERE uniqueID='`+msgReq.ID+`' allow filtering;`
	fmt.Println("value of queryString in Retrieve_msg: ", queryString)
	iterC := util.CassSession.Query(queryString).Iter()
	for iterC.Scan(&messagecontent){
		fmt.Println("value of messagecontent: ", messagecontent)
			returnMsg = util.ReturnMsg{Subject: "returnMsg", Content: messagecontent, ID: msgReq.ID}
	}
	mailMsgChannel <- returnMsg  
}

