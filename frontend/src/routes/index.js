

import About from '../components/screens/About.vue'
import Auth from '../components/screens/Auth.vue'
import People from '../components/screens/People.vue'
import Finance from '../components/screens/Finance.vue'
import Home from '../components/screens/Home.vue'
import Me from '../components/screens/Me.vue'
import Talk from '../components/screens/Talk.vue'
import FourOhFour from '../components/screens/FourOhFour.vue'

import Vue from 'vue/dist/vue.js'
import VueRouter from 'vue-router'


Vue.use(VueRouter)

const routes = [
  {
    path: '/about',
    name: 'About',
    component: About 
  }, 
  {
    path: '/auth',
    name: 'Auth',
    component: Auth 
  },
  {
    path: '/people',
    name: 'People',
    component: People 
  },
  {
    path: '/finance',
    name: 'Finance',
    component: Finance 
  },
  {
    path: '/home',
    name: 'Home',
    component: Home 
  },
  {
    path: '/me',
    name: 'Me',
    component: Me 
  },
  {
    path: '/talk',
    name: 'Talk',
    component: Talk 
  },
  {
    path: "*",
    name: "FourOhFour",
    component: FourOhFour 
  }
]

var router = new VueRouter({
  routes: routes,
  mode: 'history'
})


export default router;