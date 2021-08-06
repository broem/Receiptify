import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
// import axios from 'axios'; 
import VueRouter from 'vue-router'
import Landing from './components/Landing';
import Scanning from './components/Scanning';

Vue.use(VueRouter)

Vue.config.productionTip = false
// Vue.prototype.$http = axios 


// 2. Define some routes
// Each route should map to a component. The "component" can
// either be an actual component constructor created via
// `Vue.extend()`, or just a component options object.
// We'll talk about nested routes later.
const routes = [
  { path: '/landing', component: Landing },
  { path: '/scanning', component: Scanning }

]

// 3. Create the router instance and pass the `routes` option
// You can pass in additional options here, but let's
// keep it simple for now.
const router = new VueRouter({
  routes // short for `routes: routes`
})

new Vue({
  vuetify,
  router,
  render: h => h(App)
}).$mount('#app')
