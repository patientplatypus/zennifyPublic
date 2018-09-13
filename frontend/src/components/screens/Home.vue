<template>
  <div class="home">
    <div class='card1 homeHeader'>
      <div class='flexbox-container' :style="{flexDirection: 'row'}">
        <div :style="{flex: 1}">
          <div class='tao_sketch' id="tao_sketch"></div>
        </div>
        <div class='' :style="{flex: 5, fontSize:'calc( 2.5vw )', lineHeight:'4vh', textAlign: 'center'}">
            <div>
              welcome to {{this.siteName}}
            </div>
            <div :style="{fontSize: 'calc( 1.5vw )', lineHeight: '2vh', marginTop: '2vh', fontFamily: 'var(--alertFont)'}">
              👁 all who come shall pass 🤘
            </div>
        </div>
      </div>
    </div>
    
    <div class='card1 homeWhatis' :style="{position: 'relative'}">
      <h2>
        what is this
      </h2>
      <div :style="{position: 'absolute', top: '7vh', left: '2vw', width: '26vw', height: '60%', overflow: 'hidden', overflowY: 'auto'}">
        <p>
          a web application for the rest of us. <br/>this website has one goal in mind: <br/> to make life easier through <br/> simplicity & organization.
        </p>
      </div>
    </div>

    <div class='homeBlogpaper' :style="{marginTop: '-2vh', width: '100%', height: '100%'}">
      <div class='paper1' :style="{margin:'2.5%', width: '92.5%', height: '100%', overflowY: 'scroll'}">
        <div v-if="writingToggled.blog === true" :style="{}">
          <BlogFile/>
        </div>
        <div v-if="writingToggled.philosophy === true">
          <PhilosophyFile/>
        </div>
        <div v-if="writingToggled.userInfo === true">
          <UserInfoFile/>
        </div>
      </div>
    </div>

    <div class='homeDirections'>
      <div class='paper1' :style="{width: '95%', height: '100%'}">
        <br/>
        <h2>
          user interface
        </h2>
        <br/>

        <div :style="{overflow: 'hidden', overflowY: 'auto', height: '85%'}">

          <div class='card1' :style="{width: '85%', marginLeft: '5%', marginBottom: '2.5vh', textAlign: 'center'}">
            <p :style="{marginTop: '1vh', marginBottom: '1vh'}">
              I'm a card! Here's an
            </p>
            <input class='input1' type='text' placeholder='input field!' :style="{width: '80%', marginBottom: '1vh'}">
            <input readonly class='button1' value='and a button'/>
          </div>

          <div :style="{marginLeft: '50%'}">
            <p>
              I'm paper! Here's a graph and an informative turtle
            </p>
          </div>
          <div :style="{height: '10vh', width: '50%', marginTop: '-8vh'}">
            <canvas id="chartExample" width="100%" height="100%"></canvas>
          </div>
          <div @click="toggleTurtle()" :style="{height: '50px', width: '50px', marginTop: '2.5vh', marginLeft: '50%'}">
            <HelpfulTurtle/>
          </div>

          <div v-if="turtleToggled" :style="{marginLeft: '70%', marginTop: '-60px', fontFamily: 'var(--alertFont)'}">
            Lettuce is my favorite!
          </div>
          <div v-if="!turtleToggled" :style="{marginLeft: '70%', marginTop: '-60px', fontFamily: 'var(--alertFont)'}">
            Click me and I'll tell you things.
          </div>

        </div>


      </div>
    </div>

    <div class='homeExtra card1'>
      <h2>
        writing
      </h2>
      <div class='flexbox-container' :style="{flexDirection: 'row'}">
        <div :style="{flex: 1}">
          <div readonly class='button1' 
          @click="changeWriting('blog')"
          :class="{ toggled: writingToggled.blog}">
            <p>
              blog
            </p>
          </div>
        </div>
        <div :style="{flex: 5, textAlign: 'left', marginLeft: '2vw'}">
          daily writing
        </div>
      </div>
      <div class='flexbox-container' :style="{flexDirection: 'row'}">
        <div :style="{flex: 1}">
          <div readonly class='button1' 
          @click="changeWriting('userInfo')"
          :class="{ toggled: writingToggled.userInfo}">
            <p>
              user info
            </p>
          </div>
        </div>
        <div :style="{flex: 5, textAlign: 'left', marginLeft: '2vw'}">
          user statistics
        </div>
      </div>
      <div class='flexbox-container' :style="{flexDirection: 'row'}">
        <div :style="{flex: 1}">
          <div readonly class='button1' 
          @click="changeWriting('philosophy')"
          :class="{ toggled: writingToggled.philosophy}">
            <p>
              philosophy
            </p>
          </div>
        </div>
        <div :style="{flex: 5, textAlign: 'left', marginLeft: '2vw'}">
          philosophy
        </div>
      </div>
    </div>

    <div class='card1 homeSitemap'>
      <h2>
        sitemap
      </h2>
      <ul :style="{overflow: 'hidden', overflowY: 'scroll', height: '80%'}">
        <li>
            <span class='boldListItem'>home</span> - you are here
        </li>
        <li>
          <span class='boldListItem'>me</span> - your profile page
          <ul>
            <li>
              <span class='boldListItem'>money</span> - cost & income
            </li>
            <li>
              <span class='boldListItem'>health</span> - diet, exercise, sleep
            </li>
            <li>
              <span class='boldListItem'>social</span> - calendar, black book
            </li>
            <li>
              <span class='boldListItem'>stuff</span> - tag, categorize, offer
            </li>
            <li>
              <span class='boldListItem'>chores</span> - calendar, explanations
            </li>
          </ul>
        </li>
        <li>
          <span class='boldListItem'>others</span> - social page
          <ul>
            <li>
              <span class='boldListItem'>chat</span> - chat by category
            </li>
            <li>
              <span class='boldListItem'>message</span> - direct message
            </li>
            <li>
              <span class='boldListItem'>shop</span> - buy & sell stuff
            </li>
          </ul>
        </li>
        <li>
          <span class='boldListItem'>blog</span> - website status, notices, news, legal
        </li>
        <li>
          <span class='boldListItem'>ads</span> - advertiser links
        </li>
      </ul>
    </div>


  </div>
</template>

<script>
import { mapGetters, mapActions } from 'vuex'

import BlogFile from '../subComponents/Home/BlogFile.vue';
import PhilosophyFile from '../subComponents/Home/PhilosophyFile';
import UserInfoFile from '../subComponents/Home/UserInfoFile';

import HelpfulTurtle from '../sharedComponents/composedElements/HelpfulTurtle.vue';

import Chart from 'chart.js';
import exampleChartData from '../sharedComponents/chartJS/exampleChartData.js';

import tao_sketch from '../sharedComponents/p5sketches/tao.js'
import p5 from 'p5'


export default {
  name: 'Home',
  components: {
    BlogFile,
    PhilosophyFile,
    UserInfoFile,
    HelpfulTurtle
  }, 
  data (){
    return {
      taop5: null,
      exampleChartData: exampleChartData,
      turtleToggled: false,
      writingToggled: {
        blog: true, 
        philosophy: false, 
        userInfo: false
      }
    }
  },
  mounted() {
    var elem = document.getElementById('tao_sketch');
    this.taop5 = new p5(tao_sketch, document.getElementById('tao_sketch'))
    this.createChart('chartExample', this.exampleChartData);
  },
  computed: {
   ...mapGetters([
     'siteName'
   ])
  },
  methods:{
    createChart(chartId, chartData) {
      const ctx = document.getElementById(chartId);
      const myChart = new Chart(ctx, {
        type: chartData.type,
        data: chartData.data,
        options: chartData.options,
      });
    },
    toggleTurtle(){
      console.log('inside toggleArrowBox func and the value of this.arrowBoxToggeled: ', this.turtleToggled);
      this.turtleToggled = true;
    }, 
    changeWriting(writingType){
      console.log('inside changeWriting and writingType: ', writingType);
      for (var key in this.writingToggled) {
        if (key === writingType){
          this.writingToggled[key] = true
        }else{
          this.writingToggled[key] = false
        }
      }
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
@import '../../assets/styles/screens/home.css';
#tao_sketch{
  height: calc( 18vh );
  width: calc( 18vh );
}
</style>
