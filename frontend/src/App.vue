<template>
  <div>

    <div class='button1' v-if="loggedIn===true" 
    @click="logoutHandler"
    :style="{position: 'fixed', top: '0vh', left: '50vw', zIndex: '99', height: '2.5vh'}">
      logout
    </div>

    <div class='navMenuTitleBanner'>
      <div class='splashSign'>
        ZENNIFY.ME
      </div>
    </div>

    <div class='extraWhiteSpace'/>

    <table class='navMenuTable'>
      <tr class='tableRowSmall'>
        <td/>
        <td/>
        <td class='hamburger'>
        </td>
      </tr>
      <tr class='tableRowBig' valign='bottom'> 
        <td @click="clickNavItem(1)">
          <router-link class='routerLink' to="/home" :style="{position: 'relative'}">
            home
            <div 
            :style="{position: 'absolute', bottom: '0', left: '0', height: '10vh', width: '100%'}"
            :class="[this.navClicked===1?'fadeInNav highlightNav':'', this.navClickedPrior===1?'fadeOutNav highlightNav':'']">
            </div>
          </router-link>
        </td>
        <td @click="clickNavItem(2)">
          <router-link class='routerLink' to="/me" :style="{position: 'relative'}">
            me
            <div 
            :style="{position: 'absolute', bottom: '0', left: '0', height: '10vh', width: '100%'}"
            :class="[this.navClicked===2?'fadeInNav highlightNav':'', this.navClickedPrior===2?'fadeOutNav highlightNav':'']">
            </div>
          </router-link>
        </td>
        <td @click="clickNavItem(3)">
          <router-link class='routerLink' to="/talk" :style="{position: 'relative'}">
            talk
            <div 
            :style="{position: 'absolute', bottom: '0', left: '0', height: '10vh', width: '100%'}"
            :class="[this.navClicked===3?'fadeInNav highlightNav':'', this.navClickedPrior===3?'fadeOutNav highlightNav':'']">
            </div>
          </router-link>
        </td>
        <td @click="clickNavItem(4)">
          <router-link class='routerLink' to="/people" :style="{position: 'relative'}">
            people
            <div 
            :style="{position: 'absolute', bottom: '0', left: '0', height: '10vh', width: '100%'}"
            :class="[this.navClicked===4?'fadeInNav highlightNav':'', this.navClickedPrior===4?'fadeOutNav highlightNav':'']">
            </div>
          </router-link>
        </td>
        <td @click="clickNavItem(5)">
          <router-link class='routerLink' to="/finance" :style="{position: 'relative'}">
            MONEY
            <div 
            :style="{position: 'absolute', bottom: '0', left: '0', height: '10vh', width: '100%'}"
            :class="[this.navClicked===5?'fadeInNav highlightNav':'', this.navClickedPrior===5?'fadeOutNav highlightNav':'']">
            </div>
          </router-link>
        </td>
        <td @click="clickNavItem(6)">
          <router-link class='routerLink' to="/about" :style="{position: 'relative'}">
            ABOUT
            <div 
            :style="{position: 'absolute', bottom: '0', left: '0', height: '10vh', width: '100%'}"
            :class="[this.navClicked===6?'fadeInNav highlightNav':'', this.navClickedPrior===6?'fadeOutNav highlightNav':'']">
            </div>
          </router-link>
        </td>
        <td @click="clickNavItem(7)">
          <router-link class='routerLink' to="/auth" :style="{position: 'relative'}">
            AUTH
            <div 
            :style="{position: 'absolute', bottom: '0', left: '0', height: '10vh', width: '100%'}"
            :class="[this.navClicked===7?'fadeInNav highlightNav':'', this.navClickedPrior===7?'fadeOutNav highlightNav':'']">
            </div>
          </router-link>
        </td>
      </tr>
    </table>
    <router-view></router-view>
    <!-- <div
    :style="{height: '10vh', color: 'white'}"
    >{{testString}}</div> -->

  </div>
</template>

<script>

  import { mapGetters, mapActions } from 'vuex'

  import About from './components/screens/About.vue';
  import Auth from './components/screens/Auth.vue';
  import People from './components/screens/People.vue';
  import Finance from './components/screens/Finance.vue';
  import Home from './components/screens/Home.vue';
  import Me from './components/screens/Me.vue';
  import Talk from './components/screens/Talk.vue';

  export default {
    name: 'app',
    components: {
      About,
      Auth,
      People,
      Finance,
      Home,
      Me,
      Talk
    }, 
    data: function(){
      return {
        navClicked: -1,
        navClickedPrior: -1,
        testString: 'boop',
        urlAddress: null
      }
    }, 
    mounted(){
      this.urlAddress = window.location.hostname
      if (this.urlAddress.indexOf('localhost') != -1 || this.urlAddress.indexOf('zennify.me') != -1){
        this.setSingleVar({variableName: 'siteName', variableValue: 'zennify.me'})
      }
      this.Request({urlKEY: "cookieLogin", requestType: 'post', payload: {}})
    },
    methods:{
      ...mapActions([
        'setSingleVar', 
        'Request'
      ]),
      clickNavItem: function(itemNum){
        if(itemNum!=this.navClicked){
          this.navClickedPrior = this.navClicked;
          this.navClicked = itemNum;
        }
      }, 
      logoutHandler: function(){
        console.log('inside logoutHandler')
        const delete_cookie = (name) => {
            document.cookie = name + '=;expires=Thu, 01 Jan 1970 00:00:01 GMT;';
        };
        delete_cookie('zennifyEmail')
        delete_cookie('zennifyPassword')
        window.location.href = 'http://localhost:8081/home';
        this.setSingleVar({variableName: 'loggedIn', variableValue: false})
      }
    }, 
    computed: {
      ...mapGetters([
        'loggedIn', 
      ])
    }
  }
</script>

<style>
  /* general */
  /* @import './assets/styles/root.css';
  @import './assets/styles/global.css'; */
  /* elementUI
  @import './assets/styles/elementUI/element-variables.scss'; */
  /* elements */
  @import './assets/styles/elements/screenSize.css';
  @import './assets/styles/elements/buttons.css';
  @import './assets/styles/elements/inputs.css';
  @import './assets/styles/elements/cards.css';
  @import './assets/styles/elements/papers.css';
  @import './assets/styles/elements/arrows.css';
  @import './assets/styles/elements/selects.css';
  @import './assets/styles/elements/tables.css';
  @import './assets/styles/elements/slider.css';
  @import './assets/styles/elements/checkbox.css';
  @import './assets/styles/elements/radio.css';
  /* sheet */
  @import './assets/styles/app.css';



</style>
