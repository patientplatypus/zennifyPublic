
console.log('*** inside the store/index file ***');

import Vue from 'vue/dist/vue.js'
import Vuex from 'vuex'

import * as getters from './getters'
import * as actions from './actions'
import * as mutations from './mutations';

Vue.use(Vuex);

const state = {
    testString: 'initial',
    channelData: [],
    subChannelData: [],
    profileMsgArray: [],
    profileCanvas: [],
    usersList: [],  
    receivedMail: [],
    sentMail: [],
    email: '',
    localJWT: '', 
    siteName: '',
    loggedIn: false
}
  
const store = new Vuex.Store({
    state,
    getters,
    actions,
    mutations
})
  
  if (module.hot) {
    module.hot.accept([
      './getters',
      './actions',
      './mutations'
    ], () => {
      store.hotUpdate({
        getters: require('./getters'),
        actions: require('./actions'),
        mutations: require('./mutations')
      })
    })
  }
  
  export default store