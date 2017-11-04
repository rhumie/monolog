import Vue from 'vue'
import Router from 'vue-router'
import axios from 'axios'

import Login from '@/components/Login'
import LogDashboard from '@/components/LogDashboard'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [{
    path: '/',
    name: 'LogDashboard',
    component: LogDashboard
  }, {
    path: '/login',
    name: 'Login',
    component: Login
  }]
})

axios.interceptors.response.use((response) => {
  return Promise.resolve({
    res: response
  })
}, (error) => {
  return Promise.resolve({
    error: error.response
  })
})
