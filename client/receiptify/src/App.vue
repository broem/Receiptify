<template>
  <v-app>
    <v-app-bar
      app
      color="primary"
      dark
    >
      <div class="d-flex align-center">
        <h1>Receiptify</h1>
      </div>

      <v-spacer></v-spacer>
    </v-app-bar>

    <v-main>
      <Home/>
      <Footer/>
    </v-main>
  </v-app>
</template>

<script>
import Home from './components/Home';
import Footer from './components/Footer';

export default {
  name: 'App',

  components: {
    Home,
    Footer
  },

  data: () => ({
    //
  }),
};
</script>
