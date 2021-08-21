<template>
  <div id="landing">
    <div class="create_account"><createAccount></createAccount></div>
    <div class="digitize_label">
      <img class="d_label" src="@/assets/digitizeLabel.png" />
    </div>
    <div class="digitize">
      <router-link to="/digitize"
        ><img src="@/assets/digitize.png" @click="digitizeNavigationAlert()" />
      </router-link>
    </div>

    <div class="content">
      <event-hub></event-hub>
      <router-view></router-view>
    </div>
  </div>
</template>

<script>
import EventBus from "../services/event-bus";
import CreateAccount from "./CreateAccount.vue";

export default {
  name: "landing",
  components: {
    CreateAccount,
  },
  data: () => ({
    alertNavBar: true,
  }),

  methods: {
    digitizeNavigationAlert() {
      EventBus.$emit("alertDigitizeNav", this.alertNavBar);
    },
  },
};
</script>

<style>
#landing {
  position: relative;
  width: 100%;
  height: 100vh;
}

@media (max-width: 599px) {
  .create_account {
    position: absolute;
    left: 60%;
    top: 50%;
    bottom: 50%;
    width: 170px;
    height: 50px;
  }
  .digitize {
    position: absolute;
    object-fit: contain;
    width: 100%;
    left: 55%;
    bottom: 0%;
  }
  .digitize_label {
    position: absolute;
    object-fit: contain;
    width: auto;
    left: 15%;
    bottom: 7%;
  }

  .d_label {
    width: 180px !important;
  }
}
@media (min-width: 600px) {
  .create_account {
    position: absolute;
    left: 60%;
    top: 50%;
    bottom: 50%;
    width: 170px;
    height: 50px;
  }

  .digitize {
    position: absolute;
    object-fit: contain;
    width: 100%;
    left: 60%;
    bottom: 2%;
  }
  .digitize_label {
    position: absolute;
    object-fit: contain;
    width: auto;
    left: 15%;
    bottom: 7%;
  }

  .d_label {
    width: 200px !important;
  }
}

@media (min-width: 1100px) {
  .create_account {
    position: absolute;
    left: 80%;
    top: 50%;
    bottom: 50%;
    width: 170px;
    height: 50px;
  }
  .digitize {
    position: absolute;
    object-fit: contain;
    width: 100%;
    left: 50%;
    bottom: 2%;
  }
  .digitize_label {
    position: absolute;
    object-fit: contain;
    width: auto;
    left: 20%;
    bottom: 7%;
  }

  .d_label {
    width: 280px !important;
  }
}
</style>
