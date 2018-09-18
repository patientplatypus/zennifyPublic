<template>
  <div>
    <div v-if="!this.loggedIn" :style="{marginTop: '20vh'}">
      <Auth/>
    </div>
    <div v-if="this.loggedIn" class="people">
      <div class="peopleExplorer">
        <div class='card1' :style="{height: '98.5%', width: '95%'}">
          <h3>
            People
          </h3>
          <div v-for="(user, key) in usersList" :key="key">
            <div class='hoverClass' 
            v-if="user!=email"
            :class="{'clickedClass': userClicked===user}"
            :style="{cursor: 'pointer'}" @click="userClickHandler(user)">
              {{user}}
            </div>
          </div>
        </div>
      </div>
      <div class="peopleProfile">
        <div class='paper1' :style="{marginLeft: '1.25%', height: '98.5%', width: '95%'}">
          <h3>
            Profile
          </h3>
          <div>
            <p>
              Here is your mail:
            </p>
            <p>
              Sent Mail
            </p>
            <div :style="{height: '20vh', overflow:'hidden', overflowY: 'auto'}">
              <div v-for="(sent, key) in sentMail" :key="key">
                <div class='hoverClass' 
                :class="{'clickedClass': messageClicked===sent}"
                :style="{cursor: 'pointer'}"
                @click="messageClickHandler(sent)">
                  <p>
                    to: {{sent.Header}}
                  </p>
                  <p>
                    subject: {{sent.Subject}}
                  </p>
                </div>
              </div>
            </div>
            <p>
              Received Mail
            </p>
            <div :style="{height: '20vh', overflow:'hidden', overflowY: 'auto'}">
              <div v-for="(received, key) in receivedMail" :key="key">
                <div class='hoverClass'
                :class="{'clickedClass': messageClicked===received}"
                :style="{cursor: 'pointer'}"
                @click="messageClickHandler(received)">
                  <p>
                    to: {{received.Header}}
                  </p>
                  <p>
                    subject: {{received.Subject}}
                  </p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="peopleMessage">
        <div class='card1' :style="{height: '98.5%', width: '95%'}">
          <h3>
            Message
          </h3>
          <div v-if="userClicked!=null&&userClicked!=email" :style="{width: '100%'}">
            <p>
              Send {{userClicked}} a message!
            </p>
            <p>
              Subject
            </p>
            <input class="input1" v-model="messageSubject"/>
            <p>
              Message
            </p>
            <textarea v-model="messageContent" :style="{height: '10vh', width: '98%'}">
            </textarea>
            <div class='button1' :style="{float: 'right'}" @click="sendMessageHandler">
              send!
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'

import Auth from '../sharedComponents/Auth'

export default {
  name: 'People',
  components: {
    Auth
  },
  props: {
    msg: String,
  },
  data: function(){
    return {
      ipfs: null,
      dataString: 'dataString', 
      ipfsDBREADY: false, 
      orbitdb: null, 
      db: null,
      orbitPackage: null, 
      outerMessages: null, 
      addFunc: null,
      dispNum: 3, 
      userClicked: null, 
      messageSubject: null,
      messageContent: null, 
      messageClicked: null
    }
  },
  watch: {
    changeOperation: function(newVal, oldVal){
      this.calculate = newVal;
    }
  },
  beforeCreate: function(){
  },
  mounted: function() {

    const waitForJWT = () => {
      if(this.localJWT === null || typeof this.localJWT === 'undefined' || this.localJWT===""){
        setTimeout(() => {
          waitForJWT()
        }, 50);
      }else{
        var urlKEY = "getUsers"
        var payload = {localJWT: this.localJWT}
        console.log('value of payload before sending: ', payload);
        this.Request({urlKEY: urlKEY, requestType:'post', payload: payload})
        // this.getMailHandler();

      }
    }
    waitForJWT();


    const waitForUsers = () => {
      if(this.usersList === null || typeof this.localJWT === 'undefined' || this.localJWT===""){
        setTimeout(() => {
          waitForUsers()
        }, 50);
      }else{
        this.getMailHandler();
      }
    }
    waitForUsers();


    //HOW TO GET WEBASSEMBLY OFF THE GROUND WHEN THAT BECOMES IMPORTANT 
    // console.log('98237237492374923792347932 inside mounted!')
    // var importObject = {
    //   imports: { imported_func: arg => console.log(arg) }
    // };
    
    // const checkLoggedInFunc = () => {
    //   if (this.loggedIn === true){
        // WebAssembly.instantiateStreaming(fetch('http://localhost:8000/files/add'))
    //     .then(obj => {
    //       this.addFunc = obj.instance.exports.add_one;
    //       console.log('value of this.addFunc: ', this.addFunc(44))    
    //     })
    //     .catch(error=>{
    //       console.log('98237237492374923792347932 there was some error; ', error)
    //     });
    //   }else{
    //       setTimeout(()=>{
    //           checkLoggedInFunc()
    //       }, 20)
    //   }
    // }
    // checkLoggedInFunc()

  },
  created: function(){
    console.log('dataString is: ' + this.dataString);
    console.log('value of testString: ', this.$store.state);
  },
  computed: {
   ...mapGetters([
     'loggedIn', 
     'localJWT', 
     'usersList', 
     'email', 
     'sentMail', 
     'receivedMail'
   ])
  },
  watch:{
  },
  methods:{    
    ...mapActions([
        'Request',
    ]),
    addOne:function(){
      console.log(this.addFunc);
      this.dispNum = this.addFunc(43);
    }, 
    userClickHandler: function(user){
      this.userClicked = user;
    }, 
    messageClickHandler: function(msg){
      this.messageClicked = msg;
    },
    sendMessageHandler: function(){
      console.log('inside sendMessageHandler')
      var urlKEY = "sendUserMsg"
      var payload = {
        localJWT: this.localJWT, 
        messageFrom: this.email,
        messageTo: this.userClicked,
        messageContent: this.messageContent.toString(), 
        messageSubject: this.messageSubject.toString()
      }
      console.log('value of payload before sending: ', payload);
      this.Request({urlKEY: urlKEY, requestType:'post', payload: payload})
    }, 
    getMailHandler: function(){
      console.log('inside newVal === this.email for userClicked')
      var urlKEY = "getMail"
      var payload = {localJWT: this.localJWT, email: this.email}
      console.log('value of payload before sending: ', payload);
      this.Request({urlKEY: urlKEY, requestType:'post', payload: payload})
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  @import '../../assets/styles/screens/people.css';
  .clickedClass, .hoverClass:hover{
    color: var(--primary2);
    background: var(--secondary1)
  }
</style>
