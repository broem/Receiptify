<template>
  <div id="header">
    <link
      rel="stylesheet"
      href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/4.7.0/css/font-awesome.min.css"
    />

    <div class="topnav" id="myTopnav">
      <img class="logo" src="@/assets/receiptifyLogo.png" />

      <a @click="pushLandingPath()" class="active" v-if="landingActive">Home</a>
      <a @click="pushLandingPath()" v-if="!landingActive">Home</a>

      <a @click="pushPremiumPath()" class="active" v-if="premiumActive"
        >Premium</a
      >
      <a @click="pushPremiumPath()" v-if="!premiumActive">Premium</a>

      <a @click="pushExplorePath()" class="active" v-if="exploreActive"
        >Explore</a
      >
      <a @click="pushExplorePath()" v-if="!exploreActive">Explore</a>

      <a @click="pushProductPath()" class="active" v-if="productActive"
        >Product</a
      >
      <a @click="pushProductPath()" v-if="!productActive">Product</a>

      <a @click="pushDeveloperPath()" class="active" v-if="developerActive"
        >Developer</a
      >
      <a @click="pushDeveloperPath()" v-if="!developerActive">Developer</a>

      <a @click="pushSignInPath()" class="active" v-if="signInActive"
        >Sign In</a
      >
      <a @click="pushSignInPath()" v-if="!signInActive">Sign In</a>

      <a href="javascript:void(0);" class="icon" @click="myFunction()">
        <i class="fa fa-bars"></i
      ></a>
    </div>
  </div>
</template>

<script>
export default {
  name: "header",
  data: () => ({
    landingActive: true,
    premiumActive: false,
    exploreActive: false,
    productActive: false,
    developerActive: false,
    signInActive: false,
  }),
  methods: {
    pushLandingPath() {
      this.$router.push("/");
      this.landingActive = true;
      this.premiumActive = false;
      this.exploreActive = false;
      this.productActive = false;
      this.developerActive = false;
      this.signInActive = false;
    },
    pushPremiumPath() {
      this.$router.push("/premium");
      this.landingActive = false;
      this.premiumActive = true;
      this.exploreActive = false;
      this.productActive = false;
      this.developerActive = false;
      this.signInActive = false;
    },
    pushExplorePath() {
      this.$router.push("/explore");
      this.landingActive = false;
      this.premiumActive = false;
      this.exploreActive = true;
      this.productActive = false;
      this.developerActive = false;
      this.signInActive = false;
    },
    pushProductPath() {
      this.$router.push("/product");
      this.landingActive = false;
      this.premiumActive = false;
      this.exploreActive = false;
      this.productActive = true;
      this.developerActive = false;
      this.signInActive = false;
    },
    pushDeveloperPath() {
      this.$router.push("/developer");
      this.landingActive = false;
      this.premiumActive = false;
      this.exploreActive = false;
      this.productActive = false;
      this.developerActive = true;
      this.signInActive = false;
    },
    pushSignInPath() {
      this.$router.push("/signIn");
      this.landingActive = false;
      this.premiumActive = false;
      this.exploreActive = false;
      this.productActive = false;
      this.developerActive = false;
      this.signInActive = true;
    },
    myFunction() {
      var x = document.getElementById("myTopnav");
      if (x.className === "topnav") {
        x.className += " responsive";
      } else {
        x.className = "topnav";
      }
    },
  },
};
</script>

<style>
#header {
  position: relative;
}

/* Add a black background color to the top navigation */
.topnav {
  background-color: rgb(51, 51, 51);
  overflow: hidden;
}

/* Style the links inside the navigation bar */
.topnav a {
  float: right;
  display: block;
  color: #f2f2f2;
  text-align: center;
  padding: 14px 16px;
  text-decoration: none;
  font-size: 17px;
}

/* Change the color of links on hover */
.topnav a:hover {
  background-color: #ddd;
  color: black;
}

/* Add an active class to highlight the current page */
.topnav a.active {
  background-color: #04aa6d;
  color: white;
}

/* Hide the link that should open and close the topnav on small screens */
.topnav .icon {
  display: none;
}
.logo {
  float: left !important;
  width: 200px;
  z-index: 2;
}
/* The "responsive" class is added to the topnav with JavaScript when the user clicks on the icon. This class makes the topnav look good on small screens (display the links vertically instead of horizontally) */
@media screen and (max-width: 600px) {
  .topnav a:not(:first-child) {
    display: none;
  }
  .topnav a.icon {
    float: right;
    display: block;
  }
  .topnav.responsive {
    position: relative;
  }
  .topnav.responsive a.icon {
    position: absolute;
    right: 0;
    top: 0;
  }
  .topnav.responsive a {
    float: none;
    display: block;
    text-align: left;
  }
}
</style>
