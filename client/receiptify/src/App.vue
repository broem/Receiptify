<template>
  <v-app>
    <v-main>
      <v-parallax class="parallax" src="@/assets/background2.png"> </v-parallax>
      <Header />
      <div id="app">
        <div class="content">
          <!-- Main content -->
          <event-hub></event-hub>
          <router-view></router-view>
        </div>
      </div>
      <Footer />
    </v-main>
  </v-app>
</template>

<script>
import Footer from "./components/Footer";
import Header from "./components/Header";

export default {
  name: "App",

  components: {
    Footer,
    Header,
  },

  data: () => ({
    //
  }),
};
</script>

<style>
/* changes all parallax elements in app.:/ */
.v-parallax__image {
  overflow: hidden;
  transform: none !important;
  height: 1487px !important;
  width: 1924px !important;
  left: 0%;  
  bottom: 0%; 
}
.parallax {
  overflow: hidden;
  position: fixed;
  height: 1487px !important;
  width: 1924px !important;
  left: 0%;
  bottom: 0%;
}
</style>
