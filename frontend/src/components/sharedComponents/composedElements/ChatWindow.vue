<template>
  <div class="ChatWindow" :style="{position: 'fixed', top: locationVals.top, right: locationVals.right, bottom: locationVals.bottom, left: locationVals.left}">
    <div :style="{height: '100%', width: '100%', position: 'relative'}">
        <div :style="{position: 'absolute', top: '0vh', left: '1.25%', width: '100%', height: '100%'}">
            <h3 v-if="hideInputBar!=true" :style="{lineHeight: '5vh', height: '5vh', padding: 0, margin: 0}">
                #{{chatRoomName}}
            </h3>
            <h3 v-if="hideInputBar===true" :style="{lineHeight: '5vh', height: '5vh', padding: 0, margin: 0}">
                #feed
            </h3>
            <div  id='messageHolder' ref="messageHolder" :style="{height: 'calc(100% - 10vh)', background: 'white', width: 'calc(97.5% - 1rem)', color: 'black', overflow: 'hidden', 
            paddingLeft: '0.5rem', paddingRight: '0.5rem', overflowY: 'auto'}">
                <div v-for="(message, index) in chatArray" :key="index">
                    <span :style="{fontWeight: 900}">{{email}}</span>: {{message}}
                </div>
            </div>
        </div>
        <input
        v-if="hideInputBar===false"
        class='input1' 
        v-model="chatInput"
        v-on:keydown.enter="submitChatInput"
        :style="{width: '98%', marginLeft: '0.5%', position: 'absolute', bottom: '1vh'}">
    </div>
  </div>
</template>

<script>

import IPFS from 'ipfs';
import Room from 'ipfs-pubsub-room';

import { mapGetters, mapActions } from 'vuex';

export default {
  name: 'ChatWindow',
  props: ['locationVals', 'chatRoomName', 'hideInputBar'],
  data () {
    return {
        chatRoom: null,
        feedRoom: null,
        ipfsREADY: false,
        chatInput: null,
        chatArray: []
    }
  },
  watch:{
      chatRoomName: function(newVal, oldVal){
          if(oldVal != newVal && this.hideInputBar===false){
              console.log('83298328932898923892374239847 inside watcher and value of chatRoomName is : ', this.chatRoomName)
              this.chatArray = [];
          }
      }
  },
  mounted: function(){

    const ipfs = new IPFS({
      repo: repo(),
      EXPERIMENTAL: {
        pubsub: true
      },
      config: {
        Addresses: {
          Swarm: [
            '/ip4/127.0.0.1/tcp/9090/ws/p2p-websocket-star'
            // '/dns4/ws-star.discovery.libp2p.io/tcp/443/wss/p2p-websocket-star'
          ]
        }
      }
    })

    ipfs.once('ready', () => ipfs.id((err, info) => {
        if (err) { throw err }
        this.ipfsREADY = true;
        var tag = Date.now().toString()
        console.log('IPFS node ready with address ' + info.id)

      //   this.chatRoom.on('peer joined', (peer) => console.log('peer ' + peer + ' joined'))
      //   this.chatRoom.on('peer left', (peer) => console.log('peer ' + peer + ' left'))

      // send and receive messages

        this.chatRoom = Room(ipfs, this.chatRoomName)
        this.chatRoom.setMaxListeners(0)
        this.chatRoom.on('message', (message)=>this.handleMessageReceived(message))

        console.log('2383289328923892389893289238932 value of this.chatRoomName: ', this.chatRoomName);
        if(this.chatRoomName!='feed'){
            console.log('2383289328923892389893289238932')
            this.feedRoom = Room(ipfs, 'feed')
            this.feedRoom.setMaxListeners(0)
            this.feedRoom.on('message', (message)=>{
                this.handleMessageReceived(message)
            })
        }

    //   if(this.hideInputBar===false){
    //     this.chatRoom = Room(ipfs, this.chatRoomName)
    //     this.chatRoom.on('message', (message)=>this.handleMessageReceived(message))
    //     if(this.chatRoom!='feed'){
    //         this.feedRoom = Room(ipfs, 'feed')
    //         this.feedRoom.on('message', (message)=>this.handleMessageReceived(message))
    //     }
    //   }

    }))

    function repo () {
      return 'ipfs/pubsub-demo/' + Math.random()
    }

  },
  methods: {
    submitChatInput: function(){
        console.log('inside submitChatInput')
        if (this.ipfsREADY && this.chatInput!=null){
            this.chatRoom.broadcast(this.chatInput);
            this.chatInput = null;
        }
    }, 
    handleMessageReceived: function(message){
        console.log("2383289328923892389893289238932, inside handleMessageReceived and value of message: ", message.data.toString());
        if(this.chatRoomName!='feed' && this.hideInputBar === true){
            this.chatArray.push("#" + this.chatRoomName + "::" + message.data.toString());
        }else{
            this.chatArray.push(message.data.toString());
        }
        if(this.chatArray.length>99){
            this.chatArray = this.chatArray.slice(Math.max(this.chatArray.length - 99, 0))
        }
        this.scrollDown()
    }, 
    scrollDown: function() {
        this.$nextTick(() => {
            if(typeof this.$refs.messageHolder!=='undefined'){
                this.$refs.messageHolder.scrollTop = this.$refs.messageHolder.scrollHeight
            }
        });
    }
  }, 
  computed: {
    ...mapGetters([
     'email'
    ]),
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
    .ChatWindow{
        background: var(--secondary1);
    }
</style>
