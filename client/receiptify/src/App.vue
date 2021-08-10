<template>
  <v-app>
    <v-main>
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
    Header
  },

  data: () => ({
    //
  }),
};
</script>
