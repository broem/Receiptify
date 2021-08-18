<template>
<div id="header"> 
      <img class="logo" @click="pushLandingPath()" src="@/assets/receiptifyLogo.png" /> 
     <div class="nav">
      <v-row no-gutters>
      <v-col v-for="(k,i) in navBar.title" :key="i" cols="12" sm="2">
        <v-card class="pa-2 nav_card" @click="pushNavPath(i)">
            <div class="nav_card_text">{{k}}</div> 
        </v-card>
      </v-col>
      <v-row>
    </div> 
   </div>
      
</template>

<script>
export default {
  name:"header",
  data: () => ({
      //Just add title[x] and route[x] to object below, and it will add to nav-bar. :)
    navBar:{
      title: ["Premium", "Explore", "Product", "Developer", "Sign In"],
      route: ["/premium", "/explore", "/product", "/developer", "/signIn"]
    } 
  }),
  methods: {
   pushNavPath(i) {
     this.$router.push(this.navBar.route[i]);
   },
   pushLandingPath() {
     this.$router.push("/");
   }
  },
}
</script>

<style>
 #header {
  position: relative; 
  top: 0%;
  width: 100%;
  height: 4vh;
  background-image: linear-gradient(
    to right bottom,
    rgba(9, 46, 28, 0.788),
    rgba(6, 19, 5, 0.432)
  );
}
 
 @media(max-width:599px)
{
    .logo {
    width: 180px;
    position: absolute;
    left: 0%;  
    z-index: 2;
    box-shadow: 2px 2px 5px rgba(0, 0, 0, 0.5), 10px 10px 15px rgba(0, 0, 0, 0.5),
    -1px -1px 30px rgba(0, 0, 0, 0.2);
    }
    .nav { 
    position: absolute; 
    left: 75%;
    width: 100px; 
    object-fit: contain; 
    z-index: 2;
}
  
}

@media(min-width:600px)
{
    .logo {
    width: 240px;
    position: absolute;
    left: 0%;  
    z-index: 2;
    box-shadow: 2px 2px 5px rgba(0, 0, 0, 0.5), 10px 10px 15px rgba(0, 0, 0, 0.5),
    -1px -1px 30px rgba(0, 0, 0, 0.2);
    }  
    .nav { 
    position: absolute; 
    width: 500px; 
    left: 40%;
    object-fit: contain;
    z-index: 2;
    }
}

@media(min-width:1100px)
{
    .logo {
    width: 400px;
    position: absolute;
    left: 0%;  
    z-index: 2;
    box-shadow: 2px 2px 5px rgba(0, 0, 0, 0.5), 10px 10px 15px rgba(0, 0, 0, 0.5),
    -1px -1px 30px rgba(0, 0, 0, 0.2);
    }  
    .nav { 
    position: absolute; 
    width: 800px; 
    left: 50%;
    object-fit: contain;
    z-index: 2;
    }
}
 
 .nav_card {
  background-image: linear-gradient(
    to right bottom,
    rgba(5, 22, 22, 0.884),
    rgba(6, 5, 24, 0.678)
  );
  border-radius: 15%;

}

.nav_card_text {
  color:rgb(253, 245, 232);
  text-align: center;
  font-size: 12px;
  font-family:'Lucida Sans', 'Lucida Sans Regular', 'Lucida Grande', 'Lucida Sans Unicode', Geneva, Verdana, sans-serif
}
</style>