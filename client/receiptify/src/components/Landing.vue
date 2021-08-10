<template>
  <div id="landing">
    <div class="create_account">
      <v-chip
        style="width:100%; height:100%; text-align:center;"
        color="blue"
        link
        outlined
        >Create Account ></v-chip
      >
    </div>
    <div class="digitize_text">Start Digitizing >>></div>
    <div class="digitize">
      <router-link to="/digitize"
        ><img src="@/assets/RButton.png" />
      </router-link>
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
};
</script>

<style>
.landing {
  position: relative;
  width: 100vh;
  height: 100vh;
}

.create_account {
  position: absolute;
  left: 60%;
  top: 50%;
  bottom: 50%;
  width: 140px;
  height: 50px;
}

@media(max-width:599px)
{
  .digitize {
    position: absolute;
    object-fit: contain;
    width: 100%;
    left: 70%;
    bottom: 2%; 
  }
  .digitize_text {
    position: absolute;
    color: rgb(255, 255, 255); 
    background-image: linear-gradient(
    to right bottom,
    rgba(10, 15, 15, 0.774),
    rgba(7, 10, 8, 0.151)
  );
    border-radius: 5%;
    font-size: 20px;
    object-fit: contain;
    width: 175px;
    left: 10%;
    bottom: 7%; 
  }
}
@media(min-width:600px)
{
  .digitize {
    position: absolute;
    object-fit: contain;
    width: 100%;
    left: 60%;
    bottom: 2%; 
  }
  .digitize_text {
    position: absolute;
    color: rgb(255, 255, 255); 
    background-image: linear-gradient(
    to right bottom,
    rgba(10, 15, 15, 0.774),
    rgba(7, 10, 8, 0.151)
  );
    border-radius: 5%;
    font-size: 20px;
    object-fit: contain;
    width: 175px;
    left: 15%;
    bottom: 7%;
  }
}

@media(min-width:1100px)
{ 
  .digitize {
    position: absolute;
    object-fit: contain;
    width: 100%;
    left: 50%;
    bottom: 2%;
  } 
  .digitize_text {
    position: absolute;
    color: rgb(255, 255, 255); 
    background-image: linear-gradient(
    to right bottom,
    rgba(10, 15, 15, 0.774),
    rgba(7, 10, 8, 0.151)
  );
    border-radius: 5%;
    font-size: 20px;
    object-fit: contain;
    width: 175px;
    left: 30%;
    bottom: 7%;
  }
}
</style>
