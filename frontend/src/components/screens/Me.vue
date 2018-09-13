<template>
  <div>
    <div v-if="!this.loggedIn" :style="{marginTop: '20vh'}">
      <Auth/>
    </div>
    <div v-if="this.loggedIn">
      <div :class="meClass">
        <div class='me' :style="{position: 'fixed', top: 0, left: 0}">
          <div class='meNavi card1' style="{width: '100%', height: '100%'}">
            <h2>
              navigation
            </h2>
            <div class='flexbox-container' :style="{flexDirection: 'row'}">
              <div :style="{flex: 5, textAlign: 'right', marginRight: '2vw'}">
                my profile
              </div>
              <div :style="{flex: 1}">
                <input readonly class='button1' 
                @click="changeNav('profile')"
                :class="{toggled: navToggled.profile}"
                value='profile'/>
              </div>
            </div>
            <div class='flexbox-container' :style="{flexDirection: 'row'}">
              <div :style="{flex: 5, textAlign: 'right', marginRight: '2vw'}">
                health tracker
              </div>
              <div :style="{flex: 1}">
                <input readonly class='button1' 
                @click="changeNav('health')"
                :class="{toggled: navToggled.health}"
                value='health'/>
              </div>
            </div>
            <div class='flexbox-container' :style="{flexDirection: 'row'}">
              <div :style="{flex: 5, textAlign: 'right', marginRight: '2vw'}">
                money tracker
              </div>
              <div :style="{flex: 1}">
                <input readonly class='button1' 
                @click="changeNav('money')"
                :class="{toggled: navToggled.money}"
                value='money'/>
              </div>
            </div>
              <!-- <input readonly class='button1' value='health'/>
              <input readonly class='button1' value='money'/>
              <input readonly class='button1' value='stuff'/>
              <input readonly class='button1' value='social'/>
              <input readonly class='button1' value='chores'/> -->
          </div>
        </div>

        <div v-if="meClass==='health'">
          <MeHealth/>
        </div>

        <div v-if="meClass==='money'">
          <MeMoney/>
        </div>

        <div v-if="meClass==='profile'">
          <MeProfile/>
        </div>
        
      </div>
    </div>
  </div>
</template>

<script>

import MeHealth from '../subComponents/Me/MeHealth'
import MeMoney from '../subComponents/Me/MeMoney'
import MeProfile from '../subComponents/Me/MeProfile'


import Auth from '../sharedComponents/Auth'
import { mapGetters, mapActions } from 'vuex'

export default {
  name: 'Me',
  components: {
    Auth, 
    MeHealth,
    MeMoney, 
    MeProfile
  },
  props: {
    msg: String
  },
  data(){
    return {
      navToggled: {
        profile: true,
        health: false, 
        money: false, 
        stuff: false,
        social: false,
        chores: false
      }, 
      meClass: 'profile'
    }
  }, 
  computed: {
   ...mapGetters([
     'loggedIn'
   ])
  }, 
  methods: {
    changeNav(navType){
      // console.log('inside changeWriting and writingType: ', writingType);
      for (var key in this.navToggled) {
        if (key === navType){
          this.navToggled[key] = true
          this.meClass = key.toString()
        }else{
          this.navToggled[key] = false
        }
      }
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
  @import '../../assets/styles/screens/me.css';
</style>
