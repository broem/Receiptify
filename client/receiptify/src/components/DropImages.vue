<template>
  <div class="main-container">
    <div
      v-if="imageSources.length == 0 && !isDragging"
      class="drop-some-images"
    >
      <h1>
        Drop some images
      </h1>
    </div>
    <div
      class="drop"
      :class="getClasses"
      @dragover.prevent="dragOver"
      @dragleave.prevent="dragLeave"
      @drop.prevent="drop($event)"
    >
      <div class="img" v-for="(img, i) in imageSources" v-bind:key="i">
        <div
          class="image-button-container"
          @mouseenter="activateButton(i)"
          @mouseleave="deactivateButton(i)"
        >
          <img
            :src="img"
            @click="clickImage(img, i)"
            :class="{ active: i == activeIndex }"
          />
          <v-btn
            class="delete-button"
            small
            outlined
            v-if="activateDeleteButton && i == activeIndex"
            color="warning"
            v-on:click="removeImage(i)"
          >
            Delete</v-btn
          >
        </div>
      </div>
    </div>
    <div class="manual-image-drop">
      <label for="uploadmyfile">
        <p>or pick from device</p>
      </label>
      <input
        type="file"
        id="uploadmyfile"
        :accept="'image/*'"
        multiple
        @change="requestUploadFile"
      />
    </div>
  </div>
</template>

<script>
import EventBus from "../services/event-bus";
export default {
  name: "dropImages",

  data() {
    return {
      imageData: "",
      isDragging: false,
      imageSources: [],
      activeIndex: 0,
      activateDeleteButton: false,
      onlyOnDrop: false,
    };
  },
  computed: {
    getClasses() {
      return { isDragging: this.isDragging };
    },
  },
  methods: {
    emitMethod(files) {
      if (files.length > 0 && files != null) {
        EventBus.$emit("transferImage", files);
      }
    },
    activateButton(i) {
      if (i == this.activeIndex) {
        this.activateDeleteButton = true;
      }
    },
    deactivateButton(i) {
      if (i == this.activeIndex) {
        this.activateDeleteButton = false;
      }
    },
    removeImage(i) {
      this.imageSources.splice(i, 1);
    },
    clickImage(img, i) {
      this.activateDeleteButton = true;
      fetch(img)
        .then((res) => res.blob())
        .then((blob) => {
          const file = new File([blob], "dot.png", blob);
          let files = [];
          files.push(file);
          this.emitMethod(files);
        });

      this.activeIndex = i;
    },
    dragOver() {
      this.isDragging = true;
    },
    dragLeave() {
      this.isDragging = false;
    },
    async drop(e) {
      let files = [...e.dataTransfer.files];
      let filesforTransfer = [];
      files.forEach((file) => {
        filesforTransfer.push(file);
      });
      this.emitMethod(filesforTransfer.splice(filesforTransfer.length - 1));
      let images = files.filter((file) => file.type.indexOf("image/") >= 0);
      let promises = [];
      images.forEach((file) => {
        promises.push(this.getBase64(file));
      });
      let sources = await Promise.all(promises);
      this.imageSources = this.imageSources.concat(sources);
      this.isDragging = false;
      this.activeIndex = this.imageSources.length - 1;
    },
    requestUploadFile() {
      var src = this.$el.querySelector("#uploadmyfile");
      this.drop({ dataTransfer: src });
    },
    getBase64(file) {
      const reader = new FileReader();
      return new Promise((resolve) => {
        reader.onload = (ev) => {
          resolve(ev.target.result);
        };
        reader.readAsDataURL(file);
      });
    },
  },
};
</script>

<style scoped>
.main-container {
  position: relative;
  width: 100%;
  align-items: center;
  height: auto;
}
.image-button-container {
  position: relative;
  width: 100%;
  height: auto;
}
.delete-button {
  position: absolute;
  top: 6%;
  left: 90%;
  transform: translate(-50%, -50%);
  -ms-transform: translate(-50%, -50%);
  padding: 12px 24px;
  cursor: pointer;
  z-index: 2;
}
.active {
  border: 2px solid green;
}
.drop {
  overflow: auto;
  position: relative;
  width: 100%;
  height: 875px;
  background-color: #eee;
  border: 10px solid #eee;
  display: flex;
  align-items: center;
  justify-content: right;
  padding: 1rem;
  transition: background-color 0.2s ease-in-out;
  font-family: sans-serif;
}
.isDragging {
  background-color: #999;
  border-color: #fff;
}
.drop-some-images {
  position: absolute;
  left: 50;
  padding-top: 400px;
  width: 100%;
  align-items: center;
  justify-content: center;
  text-align: center;
  font-size: 2 rem;
  text-decoration: underline;
  z-index: 3;
}
.manual-image-drop {
  position: absolute;
  width: 100%;
  padding-top: 20px;
  align-items: center;
  justify-content: center;
  text-align: center;
  font-size: 1 rem;
  text-decoration: underline;
}
#uploadmyfile {
  display: none;
}
@media (max-width: 600px) {
  .img {
    padding: 1rem;
    display: flex;
    align-items: center;
    justify-content: center;
    max-width: 300px;
    height: 200px;
    z-index: 1;
    object-fit: contain;
  }
}
@media (min-width: 601px) {
  .img {
    padding: 1rem;
    display: flex;
    align-items: center;
    justify-content: center;
    height: 300px;
    z-index: 1;
    object-fit: contain;
  }
}
@media (min-width: 1000px) {
  .img {
    padding: 1rem;
    display: flex;
    align-items: center;
    justify-content: center;
    height: 400px;
    z-index: 1;
    object-fit: contain;
  }
}
@media (min-width: 1300px) {
  .img {
    padding: 1rem;
    display: flex;
    align-items: center;
    justify-content: center;
    height: 500px;
    z-index: 1;
    object-fit: contain;
  }
}
</style>
