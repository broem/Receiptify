<template>
  <div id="digitize">
    <div class="margin">
      <div class="drop-images"><dropImages></dropImages></div>

      <div class="button_row">
        <v-btn
          class="convert_button"
          large
          outlined
          :disabled="imageData.length == 0"
          color="primary"
          v-on:click="convertImg"
          >Convert</v-btn
        >
        <v-btn
          class="file-gen-csv csv_button"
          small
          outlined
          :disabled="translatedText.length == 0"
          color="primary"
          v-on:click="genCsv"
          >Generate Csv</v-btn
        >
      </div>
      <div class="flex_wrapper">
        <div class="image-preview" v-if="imageData.length > 0">
          <img class="preview" :src="imageData" />
        </div>
        <div class="translated_text">
          {{ translatedText }}
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import API from "../api.js";
import DropImages from "./DropImages.vue";
import EventBus from "../services/event-bus";
export default {
  name: "digitize",
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
#digitize {
  position: relative;
  width: 100%;
  height: auto;
}

.drop-images {
  position: relative;
  width: 100%;
  height: 100%;
  padding-top: 100px;
}

.convert_button {
  padding: 1rem;
  position: relative;
  align-items: center;
  justify-content: center;
  text-align: center;
  width: 200px;
  height: 120px;
}

.csv_button {
  padding: 1rem;
  position: relative;
  align-items: center;
  justify-content: center;
  text-align: center;
  width: 200px;
  height: 120px;
}

@media (max-width: 599px) {
  .flex_wrapper {
    padding: 1rem;
    position: relative;
    align-items: center;
    justify-content: center;
    text-align: center;
    height: 1000px;
    width: auto;
    margin-bottom: 100px;
  }

  .button_row {
    padding: 1rem;
    position: relative;
    align-items: center;
    justify-content: center;
    text-align: center;
    height: 120px;
    width: auto;
    border: 2px solid red;
  }
  .image_preview {
    padding: 1rem;
    position: relative;
    align-items: center;
    justify-content: center;
    text-align: center;
    max-width: 300px;
    height: 400px;
  }

  .preview {
    padding: 1rem;
    position: relative;
    align-items: center;
    justify-content: center;
    text-align: center;
    overflow-x: scroll;
    height: 400px;
    max-width: 300px;
  }

  .translated_text {
    color: rgb(40, 185, 65);
    padding: 1rem;
    position: relative;
    align-items: center;
    justify-content: center;
    text-align: center;
    width: 300px;
    max-height: 400px;
    word-wrap: break-word;
    overflow-y: auto;
    left: 9%;
  }

  .margin {
    margin-left: 1vh;
    margin-right: 1vh;
  }
}
@media (min-width: 600px) {
  .flex_wrapper {
    padding: 1rem;
    position: relative;
    display: flex;
    align-items: center;
    justify-content: center;
    text-align: center;
    height: 500px;
    width: auto;
    margin-bottom: 100px;
  }

  .button_row {
    padding: 1rem;
    position: relative;
    display: flex;
    align-items: center;
    justify-content: center;
    text-align: center;
    height: 100px;
    width: auto;
  }
  .image_preview {
    padding: 1rem;
    position: relative;
    align-items: center;
    justify-content: center;
    text-align: center;
    height: 400px;
  }

  .preview {
    padding: 1rem;
    position: relative;
    align-items: center;
    justify-content: center;
    text-align: center;
    height: 400px;
  }

  .translated_text {
    color: rgb(40, 185, 65);
    padding: 1rem;
    position: relative;
    align-items: center;
    justify-content: center;
    text-align: center;
    width: 400px;
    max-height: 400px;
    word-wrap: break-word;
    overflow-y: auto;
  }

  .margin {
    margin-left: 5vh;
    margin-right: 5vh;
  }
}

@media (min-width: 1100px) {
  .margin {
    margin-left: 10vh;
    margin-right: 10vh;
  }
}
</style>
