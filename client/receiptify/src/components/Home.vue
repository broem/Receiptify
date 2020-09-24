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
        <pre style="width:100%;overflow:scroll;">{{ translatedText }}</pre>
      </div>
    </div>
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
    API.convertRaw(this.imageData)
    .then((resp) => {
      this.translatedText = resp.data
    })
    .catch(function (err) {
      console.log(err);
    })
  }
  }
}
</script>