import Vue from "vue";
import App from "./App.vue";
import vuetify from "./plugins/vuetify";
import VueRouter from "vue-router";
import Landing from "./components/Landing";
import Digitize from "./components/Digitize";
import Developer from "./components/nav/Developer";
import Explore from "./components/nav/Explore";
import Premium from "./components/nav/Premium";
import Product from "./components/nav/Product";
import SignIn from "./components/nav/SignIn";
import CreateAccount from "./components/CreateAccount";

Vue.use(VueRouter);

Vue.config.productionTip = false;

const routes = [
  { path: "/", component: Landing },
  { path: "/digitize", component: Digitize },
  { path: "/developer", component: Developer },
  { path: "/explore", component: Explore },
  { path: "/premium", component: Premium },
  { path: "/product", component: Product },
  { path: "/signIn", component: SignIn },
  { path: "/createAccount", component: CreateAccount },
];

const router = new VueRouter({
  routes,
});

new Vue({
  vuetify,
  router,
  render: (h) => h(App),
}).$mount("#app");
