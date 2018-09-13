package data

import(
	// standard library
	"fmt"
	"log"
	"time"
	// "database/sql"
	// other libraries
	_ "github.com/lib/pq"
	"github.com/gocql/gocql"
	// "golang.org/x/crypto/bcrypt"
	//project files
	"github.com/patientplatypus/webserver/utilities"

)

func create_cassandra(){
	util.CassCluster = gocql.NewCluster("cassandra00")
	util.CassCluster.ProtoVersion = 4
	util.CassCluster.Keyspace = "platypus"
	util.CassSession, _ = util.CassCluster.CreateSession()
	defer util.CassSession.Close()
	util.CassCluster.Timeout = 5 * time.Second
}

func cassandra_message(){
	fmt.Println("cassandra_message table started")
	util.CassSession, _ = util.CassCluster.CreateSession()
	defer util.CassSession.Close()
	keySpaceMeta, _ := util.CassSession.KeyspaceMetadata("system")
	errC, exists := keySpaceMeta.Tables["cassmessage"];

    if exists != true {
        // Create a table		
        util.CassSession.Query("CREATE TABLE cassmessage(" +
            "messageto text, messagefrom text, messagecontent text, messagesubject text, uniqueID text, PRIMARY KEY (uniqueID))").Exec()
		fmt.Println("created the table cassmessage")
    }else{
		fmt.Println("skipping creation of cassmessage table - already exists")
	}
	if errC != nil{
		fmt.Println("there was an error in Cassandra create userinfo table!")
		log.Panic(errC)
	}
}

func create_userinfo_table() {
	fmt.Println("create_userinfo_table started")
	_, err2 := db.Exec(`CREATE TABLE IF NOT EXISTS userinfo(email varchar(40) PRIMARY KEY NOT NULL, password varchar(240) NOT NULL, regString  varchar(32) NOT NULL, regBool bool NOT NULL, uid serial NOT NULL);`)
	if err2 != nil {
		fmt.Println("inside create table error for create_userinfo_table")
		log.Panic(err2)
	}

	fmt.Println("replicating table in cassandra")
	util.CassSession, _ = util.CassCluster.CreateSession()
	defer util.CassSession.Close()
	keySpaceMeta, _ := util.CassSession.KeyspaceMetadata("system")
	errC, exists := keySpaceMeta.Tables["userinfo"];

    if exists != true {
        // Create a table		
        util.CassSession.Query("CREATE TABLE userinfo(" +
            "email text, password text, regString text, regBool text, uid int" +
            "PRIMARY KEY (email))").Exec()
    }
	if errC != nil{
		fmt.Println("there was an error in Cassandra create userinfo table!")
		log.Panic(errC)
	}

	fmt.Println("create_userinfo_table finished")
}

func create_profit_table(){
	fmt.Println("create_profit_table started")
	_, err2 := db.Exec(`CREATE TABLE IF NOT EXISTS profit(email varchar(40) REFERENCES userinfo, type varchar(240) NOT NULL, dateArray TEXT[] NOT NULL, amount INTEGER, interest DECIMAL, compounded TEXT, recurring TEXT, name TEXT NOT NULL, description TEXT, profitFrom TEXT, uid serial NOT NULL);`)
	if err2 != nil {
		fmt.Println("inside create table error for create_profit_table")
		log.Panic(err2)
	}
	fmt.Println("create_profit_table finished")
}

func create_expense_table(){
	fmt.Println("create_expense_table started")
	_, err2 := db.Exec(`CREATE TABLE IF NOT EXISTS expense(email varchar(40) REFERENCES userinfo, type varchar(240) NOT NULL, dateArray TEXT[] NOT NULL, amount INTEGER, interest DECIMAL, compounded TEXT, recurring TEXT, name TEXT NOT NULL, description TEXT, owedTO TEXT, uid serial NOT NULL);`)
	if err2 != nil {
		fmt.Println("inside create table error for create_expense_table")
		log.Panic(err2)
	}
	fmt.Println("create_expense_table finished")
}

func create_channel_table(){
	fmt.Println("create_channel_table started")
	_, err2 := db.Exec("CREATE TABLE IF NOT EXISTS channel(channelName varchar (40) PRIMARY KEY NOT NULL, channelDescription varchar(1000) NOT NULL, uid serial NOT NULL)")
	if err2 != nil {
		fmt.Println("inside create table error for create_channel_table")
		log.Panic(err2)
	}
}

func create_subchannel_table(){
	fmt.Println("create_subchannel_table started")
	_, err2 := db.Exec("CREATE TABLE IF NOT EXISTS subchannel(channelName varchar(40) REFERENCES channel, subChannelName varchar(40) NOT NULL, subchannelDescription varchar(1000) NOT NULL, uid serial NOT NULL)")
	if err2 != nil {
		fmt.Println("inside create table error for create_subchannel_table")
		log.Panic(err2)
	}
}

func create_ipfshash_table(){
	fmt.Println("create_IPFSHASH_table started")
	_, err2 := db.Exec("CREATE TABLE IF NOT EXISTS ipfshash(email varchar(40) UNIQUE REFERENCES userinfo, profileMsgArray TEXT[], profileCanvas TEXT[])")
	if err2 != nil {
		fmt.Println("inside create table error for create_ipfshash_table")
		log.Panic(err2)
	}
}