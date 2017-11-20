// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import 'material-design-icons/iconfont/material-icons.css'
import 'material-components-web/dist/material-components-web.css'
import 'roboto-mono-webfont/roboto-mono.min.css'
import Vue from 'vue'
import App from './App'
import router from './router'
import axios from 'axios'

Vue.config.productionTip = false

// axios global setting.
axios.defaults.baseURL = process.env.BASE_URL
axios.interceptors.request.use((config) => {
  config.headers = {
    'X-ServiceToken': `Logs ${localStorage.getItem('token') || ''}`
  }
  return config
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

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  template: '<App/>',
  components: {
    App
  }
})
