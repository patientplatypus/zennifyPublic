<template>
  <div class='screenSize'>
    <div v-if="!this.loggedIn" class='card1' :style="{marginLeft: '37.5vw'}">
      <h2>
        authentication
      </h2>
      <p>
        email
      </p>
      <input class='input1' v-on:input='updateEmail' :style="{marginTop: '1vh'}">
      <p>
        password
      </p>
      <input class='input1' type='password' v-on:input='updatePassword' :style="{marginTop: '1vh'}">
      <div :style="{marginTop: '2.5vh', marginLeft: '9vw'}">
        <input class='button1' value="register!" type='button' @click="handleUser('register')"/>
        <input class='button1' value="login!" type='button' @click="handleUser('login')"/>
        <input class='button1' value="testJWT!" type='button' @click="handleUser('testJWT')"/>
      </div>
      </div>
    </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex';
export default {
  name: 'Auth',
  props: {
    msg: String
  },
  data: function(){
    return {
      auth: {
        email: '',
        password: '',
      },
      alertNotice: ''
    }
  }, 
  computed: {
   ...mapGetters([
     'localJWT', 
     'loggedIn'
   ])
  },
  methods: {
    ...mapActions([
     'Request',
    ]),
    updateEmail: function(newEmail){
      this.auth.email = newEmail.target.value
    },
    updatePassword: function(newPassword){
      this.auth.password = newPassword.target.value
    }, 
    handleUser: function(buttonType){
      console.log('inside registerUser');
      if (this.auth.email.includes('.') && this.auth.email.includes('@')){
        if (this.auth.email!='' && this.auth.password!=''){
          this.alertNotice = ''
          if(buttonType==='register'){
            this.Request({urlKEY: 'registerUser', 
            requestType: 'post', payload: {email: this.auth.email, password: this.auth.password, localJWT: this.localJWT}})
          }else if(buttonType==='testJWT'){
            this.Request({urlKEY: 'testJWT', requestType:'post', payload: {localJWT: this.localJWT, garbage: "blahblah"}})
          }
          else{
            this.Request({urlKEY: 'loginUser', 
            requestType: 'post', payload: {email: this.auth.email, password: this.auth.password, localJWT: this.localJWT}})
          }
        }else{
          this.alertNotice = 'username and password cannot be blank!'
        }
      }else{
        this.alertNotice = 'you must use a valid email address'
      }
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
