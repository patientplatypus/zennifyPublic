<template>
  <div>
    <div v-if="!this.loggedIn" :style="{marginTop: '20vh'}">
        <Auth/>
    </div>
    <div v-if="this.loggedIn" class="talk">
      <div class='talkHeader'>
        <div class='card1' :style="{maxHeight: '95%', maxWidth: '95%'}">
          <h3>
            Talk
          </h3>
        </div>
      </div>
      <div class='talkChannels'>

        <div :style="{position: 'absolute', top: '1vh', paddingLeft: '0.5rem', paddingRight: '0.5rem'}">
          <div v-for="(channel, index) in channelData" :key="index">
            <div :class="{channelSelected: chatRoomName === channel.ChannelName}"
            :style="{cursor: 'pointer'}" @click="channelSelector(channel)">
              <span :style="{fontWeight: 900}">{{channel.ChannelName}}</span>:{{channel.ChannelDescription}}
            </div>
          </div>
        </div>

        <div :class="{addChannelOpen: addChannelOperation==='open', addChannelClosed: addChannelOperation==='closed', addChannelInitial: addChannelOperation==='initial'}">

          <div v-if="addChannelInputs==='visible'" :style="{position: 'absolute', right: '10vw', top: '1vh', left: '0.5vw'}">
            <p>
              Channel Name
            </p>
            <input 
            v-model="newChannelName"
            :style="{width: '98%'}"
            class='input1'/>
            <p>
              Channel Description
            </p>
            <input 
            v-model="newChannelDescription"
            :style="{width: '98%'}"
            class='input1'/>
          </div>


          <input readonly v-if="addChannelOperation==='open'"
          @click="addChannelHandler('cancel')"
          :style="{position: 'absolute', right: '10vw', bottom: '1vh'}"
          value='cancel' class='button1'
          />

          <input readonly 
          @click="addChannelHandler()"
          :style="{position: 'absolute', right: '0.5vw', bottom: '1vh'}"
          class='button1' :value="addChannelButtonValue"/>

        </div>
      </div>
      <div class='talkPosts'>

        <div :style="{position: 'absolute', top: '1vh', paddingLeft: '0.5rem', paddingRight: '0.5rem'}">
          <div v-for="(subChannel, index) in subChannelData" :key="index">
            <div :class="{channelSelected: subChatRoomName === subChannel.SubChannelName}"
            :style="{cursor: 'pointer'}" @click="channelSelector(subChannel, true)">
              <span :style="{fontWeight: 900}">{{subChannel.SubChannelName}}</span>:{{subChannel.SubChannelDescription}}
            </div>
          </div>
        </div>

        <div v-if="chatRoomName!='main'" :class="{addChannelOpen: addSubChannelOperation==='open', addChannelClosed: addSubChannelOperation==='closed', addChannelInitial: addSubChannelOperation==='initial'}">

          <div v-if="addSubChannelInputs==='visible'" :style="{position: 'absolute', right: '10vw', top: '1vh', left: '0.5vw'}">
            <p>
              Sub-Channel Name
            </p>
            <input 
            v-model="newSubChannelName"
            :style="{width: '98%'}"
            class='input1'/>
            <p>
              Sub-Channel Description
            </p>
            <input 
            v-model="newSubChannelDescription"
            :style="{width: '98%'}"
            class='input1'/>
          </div>


          <input readonly v-if="addSubChannelOperation==='open'"
          @click="addChannelHandler('cancel', true)"
          :style="{position: 'absolute', right: '10vw', bottom: '1vh'}"
          value='cancel' class='button1'
          />

          <input readonly 
          @click="addChannelHandler('', true)"
          :style="{position: 'absolute', right: '0.5vw', bottom: '1vh'}"
          class='button1' :value="addSubChannelButtonValue"/>

        </div>

      </div>
      <div class='talkFeed'>
      </div>
      <ChatWindow
        :style="{pointerEvents: 'all'}"
        :locationVals="chatLocationVals"
        :chatRoomName="composedChatRoomName"
        :hideInputBar="hideInputBarFalse"
      />
      <ChatWindow
        :style="{pointerEvents: 'all'}"
        :locationVals="chatLocationValsSmall"
        :chatRoomName="composedChatRoomName"
        :hideInputBar="hideInputBarTrue"
      />
    </div>
  </div>
</template>

<script>

import IPFS from 'ipfs';
import Room from 'ipfs-pubsub-room';

import ChatWindow from '../sharedComponents/composedElements/ChatWindow';

import Auth from '../sharedComponents/Auth'

import { mapGetters, mapActions } from 'vuex'

export default {
  name: 'Talk',
  components:{
    ChatWindow, 
    Auth
  },
  data () {
    return {
      chatLocationVals: {
        top: '5vh', 
        right: '0.25vw',
        bottom: '1.25vh',
        left: '60vw'
      }, 
      chatLocationValsSmall:{
        top: '75vh', 
        left: '0vw',
        bottom: '1vh',
        right: '71vw'
      },
      hideInputBarTrue: true,
      hideInputBarFalse: false,
      chatRoomName: 'feed', 
      subChatRoomName: '', 
      composedChatRoomName: 'feed',
      //channel
      addChannelButtonValue: 'add channel',
      addChannelOperation: 'initial', 
      addChannelInputs: 'invisible',
      newChannelName: null, 
      newChannelDescription: null, 
      //subchannel
      addSubChannelButtonValue: 'add subchannel',
      addSubChannelOperation: 'initial', 
      addSubChannelInputs: 'invisible',
      newSubChannelName: null, 
      newSubChannelDescription: null, 
    }
  },
  beforeDestroy: function(){
    this.setSingleVar({variableName: 'channelData', variableValue: []})
    this.setSingleVar({variableName: 'subChannelData', variableValue: []})
  },
  mounted: function(){
    setTimeout(()=>{
      console.log('value of channelData after timeout in mount: ', this.channelData)
      console.log("3930932903209320939029032097 inside localJWT setTimeout and newVal: ", this.localJWT);
    }, 10000)

    const jwtLOOPER = ()=>{
      if(this.localJWT === null || typeof this.localJWT==='undefined'){
        setTimeout(()=>{
          jwtLOOPER()
        }, 100)
      }else{
        if (this.loggedIn===false || this.loggedIn === null || typeof this.loggedIn==='undefined'){
          setTimeout(()=>{
            jwtLOOPER()
          }, 100)
        }else{
          var payload = {localJWT: this.localJWT}
          this.Request({urlKEY: 'getChannels', requestType: 'post', payload: payload})
        }
      }
    }
    jwtLOOPER();
  },
  watch: {
    // localJWT: function(newVal, oldVal){
    //   console.log("3930932903209320939029032097 inside localJWT listener and newVal: ", newVal);
    //   if (typeof newVal !== 'undefined' && newVal != null){
    //     var payload = {localJWT: this.localJWT}
    //     this.Request({urlKEY: 'getChannels', requestType: 'post', payload: payload})
    //   }
    // }, 
    channelData: function(newVal, oldVal){
      if(newVal!=oldVal){
        console.log('inside channelData watcher and newVal: ', newVal);
      }
    },
    subChannelData: function(newVal, oldVal){
      if(newVal!=oldVal){
        console.log('inside subChannelData watcher and newVal: ', newVal);
      }
    }
  },
  methods: {
    ...mapActions([
      'Request', 
      'setSingleVar'
    ]),
    addChannelHandler: function(command, subchannel){
      console.log('inside add channel handler and operation: ', this.addChannelOperation)
      if (subchannel===true){
        if (command==='cancel'){
          this.addSubChannelOperation = 'closed'
          this.addSubChannelButtonValue = 'add subchannel';
        }else{
          if (this.addSubChannelOperation === 'initial'){
            this.addSubChannelOperation = 'open';
            this.addSubChannelButtonValue = 'submit!';
          }else if (this.addSubChannelOperation === 'open'){
            this.addSubChannelOperation = 'closed';
            this.addSubChannelButtonValue = 'add subchannel';
            this.handleSubmitChannel(true)
          }else if (this.addSubChannelOperation === 'closed'){
            this.addSubChannelOperation = 'open';
            this.addSubChannelButtonValue = 'submit!';
          }
        }
        this.toggleAddChannelInputs(true);
      }else{
        if (command==='cancel'){
          this.addChannelOperation = 'closed'
          this.addChannelButtonValue = 'add channel';
        }else{
          if (this.addChannelOperation === 'initial'){
            this.addChannelOperation = 'open';
            this.addChannelButtonValue = 'submit!';
          }else if (this.addChannelOperation === 'open'){
            this.addChannelOperation = 'closed';
            this.addChannelButtonValue = 'add channel';
            this.handleSubmitChannel()
          }else if (this.addChannelOperation === 'closed'){
            this.addChannelOperation = 'open';
            this.addChannelButtonValue = 'submit!';
          }
        }
        this.toggleAddChannelInputs();
      }
    }, 
    toggleAddChannelInputs: function(subchannel){
      if (subchannel===true){
        if (this.addSubChannelInputs==='invisible'){
          setTimeout(()=>{this.addSubChannelInputs='visible'}, 400);
        }else{
          this.addSubChannelInputs = 'invisible';
        }
      }else{
        if (this.addChannelInputs==='invisible'){
          setTimeout(()=>{this.addChannelInputs='visible'}, 400);
        }else{
          this.addChannelInputs = 'invisible';
        }
      }
    }, 
    handleSubmitChannel: function(subchannel){
      console.log('inside handleSubmitChannel function')
      if (subchannel===true){
        if(this.newSubChannelName!=null && this.newSubChannelDescription!=null){
          var payload = {
            channelName: this.chatRoomName, 
            subChannelName: this.newSubChannelName,
            subChannelDescription: this.newSubChannelDescription,
            localJWT: this.localJWT
          }
          console.log('value of payload before submitting newSubChannel: ', payload);
          this.newSubChannelName = null;
          this.newSubChannelDescription = null;
          this.Request({urlKEY: 'newSubChannel', requestType: 'post', payload: payload})
        }
      }else{
        if(this.newChannelName!=null && this.newChannelDescription!=null){
          var payload = {
            channelName: this.newChannelName,
            channelDescription: this.newChannelDescription,
            localJWT: this.localJWT
          }
          this.newChannelName = null;
          this.newChannelDescription = null;
          this.Request({urlKEY: 'newChannel', requestType: 'post', payload: payload})
        }
      }
    }, 
    channelSelector: function(channel, subchannel){
      if(subchannel===true){
        this.subChatRoomName = channel.SubChannelName;
        this.composedChatRoomName = this.chatRoomName+"##"+channel.SubChannelName
      }else{
        this.chatRoomName = channel.ChannelName;
        this.composedChatRoomName = channel.ChannelName;
        var payload = {
          localJWT: this.localJWT,
          channelName: channel.ChannelName
        }
        this.Request({urlKEY: 'getSubChannels', requestType: 'post', payload: payload})
      }
    }
  }, 
  computed: {
    ...mapGetters([
     'localJWT', 
     'loggedIn', 
     'channelData',
     'subChannelData'
    ]),
  }
}
</script>

<style scoped>
  @import '../../assets/styles/screens/talk.css';
  /* :style="{position: 'absolute', bottom: '0vh', width: '100%', height: '10vh', background: 'blue'}" */
  
  .channelSelected{
    background: var(--secondary1);
  }
  .addChannelInitial{
    position: absolute;
    bottom: 0vh;
    width: 100%;
    height: 10vh;
    background: blue;
  }
  .addChannelOpen{
    position: absolute;
    bottom: 0vh;
    width: 100%;
    height: 10vh;
    -webkit-animation: openSelectAnim 0.3s ease-in; /* Safari 4.0 - 8.0 */
    animation: openSelectAnim 0.3s ease-in; 
    animation-fill-mode: forwards;
    background: blue;
  }
  .addChannelClosed{
    position: absolute;
    bottom: 0vh;
    width: 100%;
    height: 50vh;
    -webkit-animation: closeSelectAnim 0.3s ease-in; /* Safari 4.0 - 8.0 */
    animation: closeSelectAnim 0.3s ease-in; 
    animation-fill-mode: forwards;
    background: blue;
  }

     /* Safari 4.0 - 8.0 */
    @-webkit-keyframes openSelectAnim {
        0%   {height: 10vh}
        100% {height: 50vh}
    }

    /* Standard syntax */
    @keyframes openSelectAnim {
        0%   {height: 10vh}
        100% {height: 50vh}
    }

    /* Safari 4.0 - 8.0 */
    @-webkit-keyframes closeSelectAnim {
        0%   {height: 50vh}
        100% {height: 10vh}
    }

    /* Standard syntax */
    @keyframes closeSelectAnim {
        0%   {height: 50vh}
        100% {height: 10vh}
    }
</style>
