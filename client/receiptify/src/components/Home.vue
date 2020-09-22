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
  </v-container>
</template>

<script>
import API from '../api.js'
export default {
  name: 'Home',
  data: () => ({
    imageData: ''
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
    .then(function(resp) {
      console.log(resp);
    })
    .catch(function (err) {
      console.log(err);
    })
  }
  }
}
</script>