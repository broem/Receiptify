import Vue from 'vue'
import App from './App.vue'
import vuetify from './plugins/vuetify'
// import axios from 'axios'; 
import VueRouter from 'vue-router'
import Landing from './components/Landing';
import Digitize from './components/Digitize';
import Developer from './components/nav/Developer';
import Explore from './components/nav/Explore';
import Premium from './components/nav/Premium';
import Product from './components/nav/Product';
import SignIn from './components/nav/SignIn';
import CreateAccount from './components/CreateAccount';

Vue.use(VueRouter)

Vue.config.productionTip = false
// Vue.prototype.$http = axios 


// 2. Define some routes
// Each route should map to a component. The "component" can
// either be an actual component constructor created via
// `Vue.extend()`, or just a component options object.
// We'll talk about nested routes later.
const routes = [
  { path: '/', component: Landing },
  { path: '/digitize', component: Digitize },
  { path: '/developer', component: Developer },
  { path: '/explore', component: Explore },
  { path: '/premium', component: Premium },
  { path: '/product', component: Product },
  { path: '/signIn', component: SignIn },
  { path: '/createAccount', component: CreateAccount },


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
