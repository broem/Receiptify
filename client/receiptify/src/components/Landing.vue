<template>
  <div id="landing">
    <div class="logo"><img src="@/assets/RLogo.png" /></div>
    <div class="nav">
      <v-row no-gutters>
      <v-col style="width:100%" v-for="(k,i) in navBar.title" :key="i" cols="12" sm="2">
        <v-card class="pa-2" @click="pushPath(i)" outlined tile>
          {{k}} 
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
  margin: 40px;
  padding: 20px;
  width: 100%;
  height: 100%;
}

.nav {
  border: 2px solid red;
  position: absolute;
  width: 100vh;
  height: 100px;
  left: 50%;
  top: 0%;
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
