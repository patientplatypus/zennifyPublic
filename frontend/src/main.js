
// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue/dist/vue.js'
import App from './App'

import store from './store'
import router from './routes'
import VueWorker from 'vue-worker'
import Element from 'element-ui'

import { library } from '@fortawesome/fontawesome-svg-core'
import { faSave, faPencilAlt, faPalette, faSplotch, faPaintBrush } from '@fortawesome/free-solid-svg-icons'
import { FontAwesomeIcon } from '@fortawesome/vue-fontawesome'

library.add(faSave, faPencilAlt, faPalette, faSplotch, faPaintBrush)

Vue.component('font-awesome-icon', FontAwesomeIcon)

import './assets/styles/root.css';
import './assets/styles/global.css';
import '../src/assets/styles/elementUI/element-variables.scss'

Vue.use(VueWorker);
Vue.use(Element);


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