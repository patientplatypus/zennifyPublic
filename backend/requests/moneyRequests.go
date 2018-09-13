package reqt

import (
	//standard library
	"fmt"
	"net/http"
	// "encoding/json"
	//nonstandard libraries
	"github.com/gorilla/mux"
	//project files
	"github.com/patientplatypus/webserver/utilities"
	"github.com/patientplatypus/webserver/database"
)

func ProfitSubmit(w http.ResponseWriter, req *http.Request){
	fmt.Println("inside ProfitSubmit");
	vars := mux.Vars(req)
	fmt.Println("%v", vars["profitType"]);

	profitWriteChannel := make(chan string)
	go data.Write_profit_table(profitWriteChannel, vars["profitType"], req)
	profitRet := <-profitWriteChannel //blocking return
	fmt.Println("value of profitRet, ", profitRet);

	util.HandleRequestResponse(w, "Success", "inside ProfitSubmit")
}

func ExpenseSubmit(w http.ResponseWriter, req *http.Request){
	fmt.Println("inside ExpenseSubmit");
	vars := mux.Vars(req)
	fmt.Println("%v", vars["expenseType"]);

	expenseWriteChannel := make(chan string)
	go data.Write_expense_table(expenseWriteChannel, vars["expenseType"], req)
	expenseRet := <-expenseWriteChannel //blocking return
	fmt.Println("value of expenseRet, ", expenseRet);

	util.HandleRequestResponse(w, "Success", "inside ProfitSubmit")
}