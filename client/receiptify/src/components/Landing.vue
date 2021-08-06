<template>
  <div id="landing">
    <v-tabs>
      <v-tab>Premium</v-tab>
      <v-tab>Explore</v-tab>
      <v-tab>Product</v-tab>
      <v-tab>Developer</v-tab>
      <v-tab>Sign In</v-tab>
    </v-tabs>

    <div class="digitize">
      <aside><router-link to="/digitize">Digitize</router-link></aside>
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

  methods: {},
};
</script>

<style>
.landing {
  position: relative;
  margin: 40px;
  padding: 20px;
  width: 100%;
  height: 100%;
  border: 1px solid red;
}

.digitize {
  border: 1px solid red;
  position: absolute;
  text-align: center;
  align-items: center;
  right: 50%;
  left: 50%;
  width: 100px;
  height: 100px;
  line-height: 100px;
  bottom: 25px;
}
</style>
