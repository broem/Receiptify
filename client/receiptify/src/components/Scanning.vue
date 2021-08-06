<template>
  <v-container>
    <div class="drop-images"><dropImages></dropImages></div>

    <div class="image-preview" v-if="imageData.length > 0">
      <img class="preview" :src="imageData" />
    </div>
    <v-btn
      large
      outlined
      :disabled="imageData.length == 0"
      color="primary"
      v-on:click="convertImg"
      >Convert</v-btn
    >
    <div>
      <p>Text</p>
      <div v-if="translatedText.length > 0">
        <pre style="width:100%;overflow:scroll;background-color:aliceblue;">{{
          translatedText
        }}</pre>
      </div>
    </div>
    <v-btn
      class="file-gen-csv"
      small
      outlined
      :disabled="translatedText.length == 0"
      color="primary"
      v-on:click="genCsv"
      >Generate Csv</v-btn
    >
  </v-container>
</template>

<script>
import API from "../api.js";
import DropImages from "./DropImages.vue";
import EventBus from "../services/event-bus";
export default {
  name: "scanning",
  components: {
    DropImages,
  },
  data: () => ({
    imageData: "",
    translatedText: "",
    images: [],
  }),
  mounted() {
    EventBus.$on("transferImage", (images) => {
      this.images = images;
      var reader = new FileReader();
      reader.onload = (e) => {
        this.imageData = e.target.result;
      };
      reader.readAsDataURL(this.images[0]);
    });
  },
  methods: {
    convertImg: function() {
      API.convertRaw({ base64: this.imageData })
        .then((resp) => {
          this.translatedText = resp.data;
        })
        .catch(function(err) {
          console.log(err);
        });
    },
    genCsv: function() {
      API.convertCsv({ rawText: this.translatedText })
        .then((resp) => {
          console.log(resp);
          // this.translatedText = resp.data
        })
        .catch(function(err) {
          console.log(err);
        });
    },
  },
};
</script>

<style>
.drop-images {
  margin-bottom: 1rem;
  width: 100%;
  height: 100%; /* was 50% */
}
</style>
