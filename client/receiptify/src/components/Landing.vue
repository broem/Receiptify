<template>
  <div id="landing">
    <div class="logo"><img src="@/assets/RLogo.png" /></div>
    <div class="nav">
      <v-row no-gutters>
      <v-col v-for="(k,i) in navBar.title" :key="i" cols="12" sm="2">
        <v-card class="pa-2 nav_card" @click="pushPath(i)">
            <div class="nav_card_text">{{k}}</div> 
        </v-card>
      </v-col>
      <v-row>
    </div>

    <div class="create_account">
      <v-chip
        style="width:100%; height:100%; text-align:center;"
        color="blue"
        link
        outlined
        >Create Account  ></v-chip
      >
    </div>
    <div class="digitize">
      <aside>Start Digitizing
        <router-link to="/digitize"
          ><img src="@/assets/RButton.png"
        /></router-link>
      </aside>
    </div>
    <div class="content">
      <event-hub></event-hub>
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
export default {
  name: "landing",
  data: () => ({
    navBar:{
      title: ["Premium", "Explore", "Product", "Developer", "Sign In"],
      route: ["/premium", "/explore", "/product", "/developer", "/signIn"]
    } 
  }),
  methods: {
   pushPath(i) {
     this.$router.push(this.navBar.route[i]);
   }
  },
};
</script>

<style>
.landing {
  position: relative; 
  width: 100vh;
  height: 100vh;
}

.nav { 
  position: absolute; 
  width: 100vh;
  height: 100px; 
  left: 57.5%;
}

.nav_card {
  background-image: linear-gradient(
    to right bottom,
    rgba(10, 15, 15, 0.774),
    rgba(7, 10, 8, 0.151)
  );
  border-radius: 15%;

}

.nav_card_text {
  color:rgb(253, 245, 232);
  text-align: center;
  font-size: 20px;
  font-family:'Lucida Sans', 'Lucida Sans Regular', 'Lucida Grande', 'Lucida Sans Unicode', Geneva, Verdana, sans-serif
}

.logo {
  width: 100%;
  position: absolute;
  left: 0%; 
}

.create_account {
  position: absolute;
  left: 80%;
  top: 50%;
  bottom: 50%;
  width: 160px;
  height: 50px;
}

.digitize {
  position: absolute;
  object-fit: contain;
  width: 100%;
  left: 50%;
  bottom: 2%;
}
</style>
