<template>
  <v-app>
    <v-app-bar app color="primary" dark>
      <div class="d-flex align-center">
        <h1>Receiptify</h1>
      </div>

      <v-spacer></v-spacer>
    </v-app-bar>

    <v-main>
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

export default {
  name: "App",

  components: {
    Footer, 
  },

  data: () => ({
    //
  }),
};
</script>
