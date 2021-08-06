<template>
  <div id="landing">
    <aside> 
      <li><router-link to="/scanning">Scanning</router-link></li>
    </aside>
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
