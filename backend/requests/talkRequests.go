package reqt

import (
	//standard library
	"fmt"
	"net/http"
	"encoding/json"
	//nonstandard libraries
	// "github.com/gorilla/mux"
	//project files
	// "github.com/patientplatypus/webserver/utilities"
	"github.com/patientplatypus/webserver/database"
)

// type channelInfo struct{
// 	ChannelName string
// 	ChannelDescription string
// }

type getChannelName struct{
	ChannelName string `json:"channelName"`
}

type channelReturn struct{
	ChannelData []data.ChannelInfo
	Status string
	Message string
}	

type subChannelReturn struct{
	SubChannelData []data.SubChannelInfo
	Status string
	Message string
}

//channel

func NewChannel(w http.ResponseWriter, req *http.Request){
	fmt.Println("inside NewChannel");

	newTalkChannelWriteChannel := make(chan string)
	go data.Write_channel_table(newTalkChannelWriteChannel, req)
	talkRet := <-newTalkChannelWriteChannel //blocking return
	fmt.Println("value of talkRet, ", talkRet);

	// util.HandleRequestResponse(w, "Success", "inside NewChannel")

	channelData := data.Retrieve_channels()
	response := channelReturn{Status: "Success", Message: "ChannelData", ChannelData: channelData}
	json.NewEncoder(w).Encode(response)	
}

func GetChannels(w http.ResponseWriter, req *http.Request){
	fmt.Println("inside GetChannels")
	
	// channelData := []channelInfo{}

	channelData := data.Retrieve_channels()
	response := channelReturn{Status: "Success", Message: "ChannelData", ChannelData: channelData}
	json.NewEncoder(w).Encode(response)	
}


//subchannel

func NewSubChannel(w http.ResponseWriter, req *http.Request){
	fmt.Println("inside NewSubChannel");

	newTalkSubChannelWriteChannel := make(chan string)
	go data.Write_subchannel_table(newTalkSubChannelWriteChannel, req)
	talkRet := <-newTalkSubChannelWriteChannel //blocking return
	fmt.Println("value of talkRet, ", talkRet);

	// util.HandleRequestResponse(w, "Success", "inside NewChannel")

	subChannelData := data.Retrieve_subchannels(talkRet)
	response := subChannelReturn{Status: "Success", Message: "SubChannelData", SubChannelData: subChannelData}
	json.NewEncoder(w).Encode(response)	
}

func GetSubChannels(w http.ResponseWriter, req *http.Request){
	fmt.Println("inside GetSubChannels")

	var getChannelName getChannelName;
	_ = json.NewDecoder(req.Body).Decode(&getChannelName)

	fmt.Println("value of getChannelName.ChannelName")
	fmt.Println(getChannelName.ChannelName)

	subChannelData := data.Retrieve_subchannels(getChannelName.ChannelName)
	response := subChannelReturn{Status: "Success", Message: "SubChannelData", SubChannelData: subChannelData}
	json.NewEncoder(w).Encode(response)	
}