<template>
  <v-container>
    <div class="file-upload-form">
      Upload an image:
      <input type="file" @change="previewImage" accept="image/*">
    </div>
    <div class="image-preview" v-if="imageData.length > 0">
      <img class="preview" :src="imageData">
    </div>
    <v-btn large outlined :disabled="imageData.length == 0" color="primary" v-on:click="convertImg">Convert</v-btn>
    <div>
      <p>Text</p>
      <div v-if="translatedText.length > 0">
        <pre style="width:100%;overflow:scroll;background-color:aliceblue;">{{ translatedText }}</pre>
      </div>
    </div>
    <v-btn class="file-gen-csv" small outlined :disabled="translatedText.length == 0" color="primary" v-on:click="genCsv">Generate Csv</v-btn>
  </v-container>
</template>

<script>
import API from '../api.js'
export default {
  name: 'Home',
  data: () => ({
    imageData: '',
    translatedText: ''
  }),
  methods: {
  previewImage: function(event) {
    var input = event.target;
    if (input.files && input.files[0]) {

      var reader = new FileReader();
      reader.onload = (e) => {
        this.imageData = e.target.result;
    }
    reader.readAsDataURL(input.files[0]);
    }
  },
  convertImg: function() {
    API.convertRaw({base64: this.imageData})
    .then((resp) => {
      this.translatedText = resp.data
    })
    .catch(function (err) {
      console.log(err);
    })
  },
  genCsv: function() {
    API.convertCsv({rawText: this.translatedText})
    .then((resp) => {
      console.log(resp)
      // this.translatedText = resp.data
    })
    .catch(function (err) {
      console.log(err);
    })
  }
  }
}
</script>