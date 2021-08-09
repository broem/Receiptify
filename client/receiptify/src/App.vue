<template>
  <v-app>
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
