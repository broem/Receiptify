<template>
  <v-container>
    <div class="file-upload-form">
      Upload an image:
      <input type="file" @change="previewImage" accept="image/*">
    </div>
    <div class="image-preview" v-if="imageData.length > 0">
      <img class="preview" :src="imageData">
    </div>
  </v-container>
</template>

<script>
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
  }
  }
}
</script>