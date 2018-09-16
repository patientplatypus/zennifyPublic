<template>
  <div class='me profile'>
    <div class='meProfilePaper paper1'>  
        <h2>
            Profile
        </h2>
        <div class="profileImg">
            <div :style="{height: '90%', maxWidth: '100%'}">
                <img :src="image" :style="{height: '100%', width: '100%'}"/>
            </div>
            <form enctype="multipart/form-data">
                <input type="file" name="file" v-on:change="fileChange" />
                <button type="button" v-on:click="upload()">Upload</button>
            </form>
        </div>
        <div class="shareBox">
            here is what is shared from your data with other users
        </div>
        <div class="wall">
            <div :style="{position: 'relative', height: '100%', width: '100%', overflow: 'hidden'}">
                <canvas class="crayonCursor" id="canvas" v-on:mousedown="handleMouseDown" v-on:mouseup="handleMouseUp" v-on:mousemove="handleMouseMove"
                :style="{position: 'absolute', zIndex: '1'}"
                >
                </canvas>
                <div :style="{position: 'absolute', zIndex: '6', pointerEvents: 'none', bottom: '5px', left: '5px'}">
                    <div v-if="paintButtonValue==='apply'" @click="paintMenuHandler" :style="{pointerEvents: 'all', cursor: 'pointer'}">
                        <font-awesome-icon icon="save" size="3x" :style="{color: 'var(--secondary2)'}"/>
                    </div>
                    <div v-if="paintButtonValue==='palette'" @click="paintMenuHandler" :style="{pointerEvents: 'all', cursor: 'pointer'}">
                        <font-awesome-icon icon="palette" size="3x" :style="{color: 'var(--secondary2)'}"/>
                    </div>
                </div>
                <div 
                :class="{'paintMenuOpen':this.paintMenu==='open', 'paintMenuClose':this.paintMenu==='close', 'paintMenuInitial':this.paintMenu==='initial'}"
                :style="{position: 'absolute', zIndex: '3', height: '100%', width: '100%'}">
                    <div :style="{position: 'relative', height: '100%', width: '100%', pointerEvents: 'none'}">

                        <div class='RpaletteSlider' :style="{marginLeft: '10%', width: '80%', pointerEvents: 'all'}">
                            <el-slider :step="1" :min="0" :max="256"  v-model="RSlider"></el-slider>
                        </div>

                        <div class='GpaletteSlider' :style="{marginLeft: '10%', width: '80%', pointerEvents: 'all'}">
                            <el-slider :step="1" :min="0" :max="256"  v-model="GSlider"></el-slider>
                        </div>
                        <div class='BpaletteSlider' :style="{marginLeft: '10%', width: '80%', pointerEvents: 'all'}">
                            <el-slider :step="1" :min="0" :max="256"  v-model="BSlider"></el-slider>
                        </div>

                        <div class='ApaletteSlider' :style="{marginLeft: '10%', width: '80%', pointerEvents: 'all'}">
                            <el-slider :step=".01" :min="0" :max="1"  v-model="ASlider"></el-slider>
                        </div>
                        
                        <div :style="{marginLeft: '90%', pointerEvents: 'all'}">
                            <el-slider :step=".01" :min="0" :max="10"  v-model="SizeSlider" vertical height="22.5vh"></el-slider>
                        </div>

                   
                        <div :style="{position: 'absolute', zIndex: '5', top: '70%', color: drawColor.drawReturn, fontSize: `${SizeSlider*2}vh`, 
                        textAlign: 'center', width: '100%', marginRight: `${SizeSlider*2}vh`, marginTop: `-${SizeSlider*1}vh`}">
                            <font-awesome-icon icon="paint-brush"/>
                        </div> 
                    </div>
                </div>
            </div>
        </div>
        <div class="stats">
            here are your zennify stats
            <label class="checkBoxContainer">
                one month
                <input type="checkbox" value="test1" v-model="checkboxChecked">
                <span class="checkmarkCheckBox"></span>
            </label>
            <div class="button1" @click="updateSharing">
                <p>
                    update sharing
                </p>
            </div>
        </div>
        <div class="loca">
            here is your location
        </div>
        <div class='aslInfo'>
            here is your aslInfo
        </div>
        <div class="msgBox">
            here is your message box

            <div :style="{position: 'relative', height: '100%', width: '100%'}">
                <div ref="messageHolder" :style="{position: 'absolute', top: '2.5vh', width: '100%', height: '70%', background: 'rgba(200, 200, 100)', color: 'black', overflow: 'hidden', overflowY: 'auto'}">
                    <div v-for="(data, key) in msgData" :key="key">
                        {{data}}
                    </div>
                </div>
                <input
                v-if="localJWT!=null && email!=null"
                v-model="msgInput"
                v-on:keydown.enter="submitMsg"
                class="input1" :style="{position: 'absolute', top: '80%', marginLeft: '1vw'}"/>
            </div>
        </div>
    </div>
  </div>
</template>

<script>

import axios from "axios";
import IPFS from 'ipfs';
import { mapGetters, mapActions } from 'vuex';

import '../../sharedComponents/svg/crayon.svg';

export default {
  name: 'MeProfile',
  data(){
      return{
        files: new FormData(), 
        imgFile: null,
        image: null, 
        msgInput: null, 
        msgNodeREADY: false, 
        msgNode: null, 
        canvasInput: null,
        canvasNodeREADY: false, 
        canvasNode: null, 
        msgHash: null, 
        msgData: [], 
        mouse: {
            current: {
                x: 0,
                y: 0
            },
            previous: {
                x: 0,
                y: 0
            },
            down: false
        },
        RSlider: 100, 
        GSlider: 100, 
        BSlider: 100, 
        ASlider: 0.5,
        SizeSlider: 5,
        paintMenu: 'initial', 
        paintButtonValue: 'palette', 
        canvasURL: null
      }
  }, 
  mounted: function(){

    function fitToContainer(canvas){
        canvas.style.width='100%';
        canvas.style.height='100%';
        canvas.width  = canvas.offsetWidth;
        canvas.height = canvas.offsetHeight;
    }

    var canvas = document.querySelector('canvas'),
    ctx    = canvas.getContext('2d');
    fitToContainer(canvas);


    
    const grabHashes = () => {
        var urlKEY = "getHashMsg";
        var payload = {
            localJWT: this.localJWT,
            email: this.email,
            hashName: "profileMsgArray", 
        };
        this.Request({urlKEY: urlKEY, requestType:'post', payload: payload})

        var urlKEY = "getHashMsg";
        var payload = {
            localJWT: this.localJWT,
            email: this.email,
            hashName: "profileCanvas", 
        };
        this.Request({urlKEY: urlKEY, requestType:'post', payload: payload})
    }

    grabHashes();


    const checkLoggedInFunc = () => {
        if (this.loggedIn === true){
            var imageurl = 'http://localhost:8000/files/email_test';
            axios.get(imageurl, {
                responseType: 'arraybuffer'
            })
            .then( response => {
                console.log('value of response: ', response);
                console.log('value of response.data: ', response.data);
                const base64 = btoa(
                    new Uint8Array(response.data).reduce(
                        (data, byte) => data + String.fromCharCode(byte),
                        '',
                    ),
                );
                console.log('value of base64: ', base64);
                this.image = "data:;base64," + base64 
            });
        }else{
            setTimeout(()=>{
                checkLoggedInFunc()
            }, 20)
        }
    }
    checkLoggedInFunc()

    const createIPFS = () => {
        // https://github.com/ipfs/js-ipfs/blob/master/examples/browser-webpack/src/components/app.js
        this.msgNode = new IPFS({ repo: String('msg' + Math.random() + Date.now()) })

        this.msgNode.once('ready', () => {
            console.log('IPFS message node is ready')
            this.msgNodeREADY = true;
                // ops()
        })

        this.canvasNode = new IPFS({ repo: String('canvas' + Math.random() + Date.now()) })

        this.canvasNode.once('ready', () => {
            console.log('IPFS canvas node is ready')
            this.canvasNodeREADY = true;
                // ops()
        })
    }

    createIPFS();

  },
  methods: {
    ...mapActions([
        'Request',
    ]),
    paintMenuHandler: function(){
        console.log("inside paintMenu");
        if (this.paintMenu === "open"){
            this.paintMenu = "close";
            this.paintButtonValue = 'palette';
        }else if (this.paintMenu === "close" || this.paintMenu === 'initial'){
            this.paintMenu = "open"
            this.paintButtonValue = 'apply';
        }
    },
    canvasDataSend: function(){
        console.log('inside canvasDataSend')
        var c = document.getElementById('canvas');
        var base_image = null;
        
        var ctx = c.getContext("2d");
        var imgData = ctx.getImageData(0, 0, c.width, c.height);
        var buffer = imgData.data.buffer;

        console.log('value of canvas buffer before adding to ipfs: ', buffer)

        this.canvasNode.files.add(Buffer.from(buffer), {pin: false}, (err, filesAdded)=>{
            if (err) { throw err }  
            console.log('inside canvasNode.files.add');
            this.canvasHash = filesAdded[0].hash;
            console.log('and value of this.canvasHash is .... ', this.canvasHash)
            console.log('value of entirety of filesAdded[0]', filesAdded[0])
            this.submitDataRequest(this.canvasHash, "profileCanvas");
        })
    },
    draw: function (event) {
      // requestAnimationFrame(this.draw);
        if (this.mouse.down) {
            var c = document.getElementById("canvas");
            var ctx = c.getContext("2d");
            ctx.lineTo(this.currentMouse.x, this.currentMouse.y);
            ctx.strokeStyle = this.drawColor.drawReturn;
            ctx.lineWidth = this.SizeSlider;
            ctx.stroke()
        }
    },
    handleMouseDown: function (event) {
        console.log('inside handleMouseDown')
        this.mouse.down = true;
        this.mouse.current = {
            x: event.pageX,
            y: event.pageY
        }
        var c = document.getElementById("canvas");
        var ctx = c.getContext("2d");
        // ctx.clearRect(0, 0, c.height, c.width)
        ctx.beginPath();
        ctx.moveTo(this.currentMouse.x, this.currentMouse.y)
    },
    handleMouseUp: function () {
        console.log('inside handleMouseDown')
        this.mouse.down = false;
        var c = document.getElementById("canvas");
        var ctx = c.getContext("2d");
        ctx.closePath();
        if(this.canvasNodeREADY===true){
            this.canvasDataSend();
        }
    },
    handleMouseMove: function (event) {
        // console.log('inside handleMouseMove')
        this.mouse.current = {
            x: event.pageX,
            y: event.pageY
        }
        this.draw(event)
    },
    fileChange(e) {
        this.files.append("uploadFile",  e.target.files[0]);
        this.files.append("name", "test")
        this.files.append("email", "email")
        this.imgFile = e.target.files || e.dataTransfer.files;
    },
    upload() {
        this.createImage(this.imgFile[0]);
        // console.log('value of this.files._boundary', this.files._boundary);

        axios({
            method: 'post',
            url: 'http://localhost:8000/upload',
            data: this.files,
            config: { headers: {'Content-Type': 'multipart/form-data' }}
        })
        .then(function (response) {
            //handle success
            console.log(response);
        })
        .catch(function (response) {
            //handle error
            console.log(response);
        });
    }, 
    createImage: function(file) {
        var image = new Image();
        var reader = new FileReader();
        var vm = this;  
        reader.onload = (e) => {
            vm.image = e.target.result;
        };
        reader.readAsDataURL(file);
    },
    submitMsg: function(){
        console.log('inside submitMsg');
        if (this.msgNodeREADY===true){
            this.msgNode.files.add([Buffer.from(this.msgInput)], {pin: true}, (err, filesAdded)=>{
                if (err) { throw err }  
                this.msgHash = filesAdded[0].hash;
                console.log('and value of this.msgHash is .... ', this.msgHash)
                this.msgNode.files.cat(this.msgHash, (err, data) => {
                        if (err) { throw err }
                        this.msgData.push(data.toString());
                        this.msgInput = null;
                        this.submitDataRequest(this.msgHash, "profileMsgArray");
                        this.scrollDown();
                })
            })
        }
    },
    checkboxChecked: function(e){
        console.log('inside checkboxChecked')
    },
    submitDataRequest: function(hash, hashName){
        console.log('inside submitDataRequest')
        var urlKEY = "newHashMsg";
        var payload = {
            localJWT: this.localJWT,
            email: this.email,
            hashName: hashName, 
            hashWrite: hash
        };
        this.Request({urlKEY: urlKEY, requestType:'post', payload: payload})
    }, 
    scrollDown: function() {
        this.$nextTick(() => {
            if(typeof this.$refs.messageHolder!=='undefined'){
                this.$refs.messageHolder.scrollTop = this.$refs.messageHolder.scrollHeight
            }
        });
    }, 
    updateSharing(){
        console.log('got inside updateSharing')
    }
  }, 
    computed:{
    ...mapGetters([
     'loggedIn', 
     'localJWT', 
     'email', 
     'profileMsgArray',
     'profileCanvas'
    ]),
    currentMouse: function () {
        var c = document.getElementById("canvas");
        var rect = c.getBoundingClientRect();
      
        return {
            x: this.mouse.current.x - rect.left,
            y: this.mouse.current.y - rect.top
        }
    }, 
    drawColor: function(){
        var drawReturn = `rgba(${this.RSlider}, ${this.GSlider}, ${this.BSlider}, ${this.ASlider})`
        return{
            drawReturn
        }
    }
  }, 
  watch: {
      profileMsgArray: function(newVal, oldVal){
          console.log('inside profileMsgArray')
          if (oldVal!=newVal && newVal!=null){
              console.log('new != old')
              console.log('value of new: ', newVal)
              const waitForIPFS = () => {
                  if (this.msgNodeREADY === true){
                    var tempArr = new Array(newVal.length);
                    newVal.forEach((val, index)=>{
                        console.log('value of newVal val: ', val)
                        this.msgNode.files.cat(val, (err, data) => {
                            if (err) { throw err }
                            tempArr[index] = data.toString();
                            if (tempArr.indexOf(undefined)===-1){
                                this.msgData = tempArr;
                                this.$forceUpdate()
                                this.scrollDown();
                            }
                        })    
                    })

                  }else{
                      setTimeout(()=>{
                          waitForIPFS()
                      }, 50)
                  }
              }
              waitForIPFS()
          }
      }, 
      profileCanvas: function(newVal, oldVal){
          console.log('inside profileCanvas');
          if (oldVal!=newVal && newVal!=null){
            console.log('new != old')
            console.log('value of new: ', newVal)
            console.log('value of newVal[0]: ', newVal[0])
            var hashVal = newVal[0]
            const waitForIPFS = () => {
                console.log('inside WAITFORIPFS for canvas')
                if (this.canvasNodeREADY === true){
                    console.log('inside canvasnodeready watch')
                    console.log('value of hashVal: ', hashVal)
                    console.log('this.canvasNode.files')
                    console.log(this.canvasNode.files)
                    this.canvasNode.files.cat(hashVal, (err, data) => {
                        console.log('inside profile canvasNode file cat')
                        console.log('and value of data')
                        console.log(data)
                        if (err) { 
                            console.log('some error in profileCanvas watch')
                            console.log('error; ', err)
                            throw err 
                        }  
                        var c = document.getElementById('canvas');
                        var ctx = c.getContext("2d");
                        var imageData = ctx.createImageData(c.width, c.height);
                        // imageData.data.set(data);

                        // console.log('typeof data: ', typeof data);

                        // console.log('**** create image data start ****')
                        for (var i=0;i < data.length; i+=4) {
                            // console.log('inside imageData.data loop and...')
                            imageData.data[i]   = data[i];   //red
                            // console.log('value of red: ', data[i])
                            imageData.data[i+1] = data[i+1]; //green
                            // console.log('value of green: ', data[i+1])
                            imageData.data[i+2] = data[i+2]; //blue
                            // console.log('value of blue: ', data[i+2])
                            imageData.data[i+3] = data[i+3]; //alpha
                            // console.log('value of alpha: ', data[i+3])
                        }
                        console.log('putting imageData to ctx:')
                        ctx.putImageData(imageData,0,0);



                        // var imageData = new Image()
                        // imageData.onload = () => {
                        //     console.log('inside imageData onload....')
                        //     imageData = ctx.createImageData(ctx.width, ctx.height);
                        //     imageData.data.set(data);
                        // }
                        // console.log('value of data back to utf8')
                        // console.log(data.toString())
                    })
                }else{
                      setTimeout(()=>{
                          waitForIPFS()
                      }, 50)
                }
            }
            waitForIPFS()
          }
      }
  }
}
</script>


<style>

.RpaletteSlider .el-slider__bar{
    background-color: red;
}
    
.RpaletteSlider .el-slider__button{
    border-color: red;
    background-color: red;
}

.GpaletteSlider .el-slider__bar{
    background-color: green;
}
    
.GpaletteSlider .el-slider__button{
    border-color: green;
    background-color: green;
}

.BpaletteSlider .el-slider__bar{
    background-color: blue;
}
    
.BpaletteSlider .el-slider__button{
    border-color: blue;
    background-color: blue;
}

.ApaletteSlider .el-slider__bar{
    background-color: grey;
}
    
.ApaletteSlider .el-slider__button{
    border-color: grey;
    background-color: grey;
}

.VertSizeSlider .el_slider__runway{
    height: 50px;
}

</style>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
@import '../../../assets/styles/screens/me.css';

.paintMenuOpen{
    transform: translate(-100%);
    animation: openPaintMenu 0.3s ease-in; 
    animation-fill-mode: forwards;
    overflow: hidden;
    background: var(--primary1);
}

.paintMenuClose{
    transform: translate(0%);
    animation: closePaintMenu 0.3s ease-in; 
    animation-fill-mode: forwards;
    overflow: hidden;
    background: var(--primary1);
}

.paintMenuInitial{
    transform: translate(-100%);
    overflow: hidden;
    background: var(--primary1);
}

   /* Safari 4.0 - 8.0 */
    @-webkit-keyframes openPaintMenu {
        0%   {transform: translate(-100%)}
        100% {transform: translate(0%)}
    }

    /* Standard syntax */
    @keyframes openPaintMenu {
        0%   {transform: translate(-100%)}
        100% {transform: translate(0%)}
    }

       /* Safari 4.0 - 8.0 */
    @-webkit-keyframes closePaintMenu {
        0%   {transform: translate(0%)}
        100% {transform: translate(-100%)}
    }

    /* Standard syntax */
    @keyframes closePaintMenu {
        0%   {transform: translate(0%)}
        100% {transform: translate(-100%)}
    }


.profileImg{
    position: absolute;
    height: 30vh;
    /* width: 20vw; */
    top: 7.5vh;
    left: 1vw;
    background: salmon;
}

.widthSlider{
    width: 50%;
}

.wall{
    position: absolute;
    height: 45vh;
    width: calc(27vw - 4px);
    top: 7.5vh;
    left: 20vw;
    border-width: 2px;
    border-style: solid;
    border-color: black;
    color: white;
}

.stats{
    position: absolute;
    height: 52.5vh;
    width: 18vw;
    top: 40vh;
    left: 1vw;
    background: brown;
    color: white;
}

.loca{
    position: absolute;
    height: 35vh;
    width: 20vw;
    top: 7.5vh;
    right: 1vw;
    background: green;
    color: white;
}

.aslInfo{
    position: absolute;
    height: 47.5vh;
    width: 20vw;
    top: 45vh;
    right: 1vw;
    background: orange;
    color: white;
}

.msgBox{
    position: absolute;
    height: 37.5vh;
    width: 27vw;
    top: 55vh;
    left: 20vw;
    background: purple;
    color: white;
}

.shareBox{
    position: absolute;
    height: 52.5vh;
    width: 18vw;
    top: 40vh;
    left: 1vw;
    background: red;
    color: white;
}

.crayonCursor{
    cursor:url(http://www.cursor.cc/cursor/263/42/cursor.png), auto; 
}

</style>
