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
      <Landing/>
      <Footer/>
    </v-main>
  </v-app>
</template>

<script> 
import Footer from './components/Footer'; 
import Landing from './components/Landing';

export default {
  name: 'App',

  components: { 
    Footer, 
    Landing
  },

  data: () => ({
    //
  }),
};
</script>
