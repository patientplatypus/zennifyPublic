
// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue/dist/vue.js'
import App from './App'

import store from './store'
import router from './routes'
import VueWorker from 'vue-worker';

Vue.use(VueWorker);


Vue.config.productionTip = false;


/* eslint-disable no-new */
var vm = new Vue({
  el: '#app',
  router: router,
  store,
  template: '<App/>',
  components: { App }
})

global.vm = vm; //Define you app variable globally