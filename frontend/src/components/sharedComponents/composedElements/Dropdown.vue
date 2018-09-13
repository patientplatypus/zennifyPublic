<template>
  <div class='Dropdown' :class='dataArrayName'>
    <div class='selectBox1' :style="{position: 'relative', cursor: 'pointer', pointerEvents: 'none'}">

        <div class='flexbox-container' :style="{flexDirection: 'row'}">
            <div :style="{flex: 10}">
                {{selectText}}
            </div>
            <div :style="{flex: 1, marginRight: '0.5vw'}">
                <div class='arrow-down' :class="{arrowRotateUp: this.outsideClick==='openSelect', arrowRotateDown: this.outsideClick==='closeSelect'}"></div>
            </div>
        </div>

        <!-- <div class='selectStick'></div> -->
        <div :style="{position: 'absolute', right: 0, top: 'calc(1.3*var(--select-font-size-1))', margin:0, zIndex: 3, height: this.listHeight, cursor: this.outsideClick==='openSelect'?'pointer':'default', width: '100%'}">
            <div :class="handleOpen" class='optionBox1' @click='openStatus = false' :style="{overflowY: this.setOverflowY===true?'auto':'none'}">
                <div class="listItem" :class='dataArrayName' v-for="datum in this.dataArray" :style="{textAlign: 'right', paddingLeft: '1vw', paddingRight: '1vw', pointerEvents: 'all'}">
                    <div 
                    v-on:click="handleSelect(datum)">
                        {{datum}}
                    </div>
                </div>
            </div>
        </div>
    </div>
  </div>
</template>

<script>

import { Bus } from '../bus/Bus';

export default {
  name: 'Dropdown',
  props: ['dropDataArray'],
  data: function(){
    return {
        openStatus: false,
        neverClicked: true,
        listHeight: '0vh', 
        outsideClick: 'closeSelectInitial',
        selectText: 'please choose',
        dataArray: [],
        dataArrayName: 'placeholder',
        selectName: 'placeholder',
        setOverflowY: false
    }
  }, 
  beforeDestroy: function(){
      window.removeEventListener('click', this.windowFunc)
  },
  watch: {
    $props: {
      handler() {
        this.parseData();
      },
      deep: true,
      immediate: true,
    }
  },
  mounted: function(){

    // console.log('^^^^value of dataArray: ', this.dataArray);

    var vHArray; 
    //can't scroll through a list yet for some dumb reason - ruhroh
    if (this.dataArray.length>3){
        var vHArray = 3;
        this.setOverflowY = true;
    }else{
        var vHArray = this.dataArray.length;
    }
    var vHArrayString = "calc(1.2*var(--select-font-size-1) * "+vHArray.toString()+")";
    this.listHeight = vHArrayString;

    window.addEventListener('click', this.windowFunc);
  
  },
  methods:{
      handleSelect: function(datum){
            this.selectText = datum;
            setTimeout(()=>{
                console.log("3828293839283 value of this.dropDataArray[0]: ", this.dropDataArray[0]);
                // this.$emit('selectReturn', {name: this.dataArrayName, value: datum});
                this.$emit('selectReturn', {name: this.dropDataArray[0], value: datum})
            }, 300)
      },
      windowFunc: function(event){
        // console.log(event.target);
        // console.log(event.target.tagName);
        // console.log(event.target.className);
        // console.log('value of this.dataArrayName: ', this.dataArrayName);
        // console.log("this.dataArrayName indexOf: ", event.target.className.indexOf(this.dataArrayName));
        if (event.target.className.indexOf(this.dataArrayName)===-1){
            if(this.neverClicked===false){
                console.log('inside window click if statement')
                this.outsideClick = 'closeSelect';
                this.openStatus = false;
            }
        }else{
            if(this.openStatus===false){
                this.openStatus = true;
                this.neverClicked = false; 
                this.outsideClick = 'openSelect'
            }else{
                this.openStatus = false;
                this.outsideClick = 'closeSelect'
            }
        }
      }, 
      parseData: function(){
        console.log('3828293839283 value of this.dropDataArray: ', this.dropDataArray);
        this.dataArrayName = this.dropDataArray[0]
        this.dataArray = this.dropDataArray.slice(1);
        this.outsideClick = 'closeSelectInitial';
        this.selectText = 'please choose';
      }
  }, 
  computed:{
      handleOpen: function(){
          return(this.outsideClick)
      }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

    .arrowRotateUp{
        -webkit-animation: rotateUpAnim 0.3s ease-in; /* Safari 4.0 - 8.0 */
        animation: rotateUpAnim 0.3s ease-in; 
        animation-direction: normal;
        animation-fill-mode: forwards;
    }


    .arrowRotateDown{
        transform: rotate(180deg);
        -webkit-animation: rotateDownAnim 0.3s ease-in; /* Safari 4.0 - 8.0 */
        animation: rotateDownAnim 0.3s ease-in; 
        animation-direction: normal;
        animation-fill-mode: forwards;
    }


    .openSelect{
        overflow: hidden;
        height: 0%;
        -webkit-animation: openSelectAnim 0.3s ease-in; /* Safari 4.0 - 8.0 */
        animation: openSelectAnim 0.3s ease-in; 
        animation-fill-mode: forwards;
    }

    .closeSelect{
        height: 100%;
        -webkit-animation: closeSelectAnim 0.3s ease-in; /* Safari 4.0 - 8.0 */
        animation: closeSelectAnim 0.3s ease-in; 
        animation-fill-mode: forwards;
        overflow: hidden;
    }

   /* Safari 4.0 - 8.0 */
    @-webkit-keyframes rotateUpAnim {
        0%   {transform: rotate(0deg)}
        100% {transform: rotate(180deg)}
    }

    /* Standard syntax */
    @keyframes rotateUpAnim {
        0%   {transform: rotate(0deg)}
        100% {transform: rotate(180deg)}
    }

    @-webkit-keyframes rotateDownAnim {
        0%   {transform: rotate(180deg)}
        100% {transform: rotate(0deg)}
    }

    /* Standard syntax */
    @keyframes rotateDownAnim {
        0%   {transform: rotate(180deg)}
        100% {transform: rotate(0deg)}
    }

    /* Safari 4.0 - 8.0 */
    /* @-webkit-keyframes openSelectAnim {
        0%   {height: 0%; border-color: transparent; visibility: visible;}
        1%   {border-color:var(--secondary2)}
        100% {height: 100%;}
    } */

    /* Standard syntax */
    /* @keyframes openSelectAnim {
        0%   {height: 0%; border-color: transparent; visibility: visible;}
        1%   {border-color:var(--secondary2)}
        100% {height: 100%;}
    } */

        /* Safari 4.0 - 8.0 */
    @-webkit-keyframes openSelectAnim {
        0%   {height: 0%; border-width: 2px;}
        /* 1%   {border-color:var(--secondary2)} */
        100% {height: 100%;}
    }

    /* Standard syntax */
    @keyframes openSelectAnim {
        0%   {height: 0%; border-width: 2px;}
        /* 1%   {border-color:var(--secondary2)} */
        100% {height: 100%;}
    }

    /* Safari 4.0 - 8.0 */
    @-webkit-keyframes closeSelectAnim {
        0%   {height: 100%;}
        /* 99%  {border-color: var(--seconday2)} */
        100% {height: 0%; border-width: 0px;}
    }

    /* Standard syntax */
    @keyframes closeSelectAnim {
        0%   {height: 100%;}
        /* 99%  {border-color: var(--seconday2)} */
        100% {height: 0%; border-width: 0px;}
    }

    .closeSelectInitial{
        height: 0;
        overflow: hidden;
        visibility: hidden;
        /* border-color: transparent; */
    }

    .itemHover{
        cursor: pointer;
        /* padding-top: 1vh; */
    }

    .listItem:hover{
        background: var(--secondary1);
        color: var(--primary2);
    }

</style>
