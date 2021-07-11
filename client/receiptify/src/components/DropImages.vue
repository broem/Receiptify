<template>
  <div
    class="drop"
    :class="getClasses"
    @dragover.prevent="dragOver"
    @dragleave.prevent="dragLeave"
    @drop.prevent="drop($event)"
  >
    <div class="img" v-for="(img, index) in imageSources" v-bind:key="index">
      <img
        :src="img"
        @click="clickImage(img)"
        v-on:click="toggle(index)"
        :class="{ active: index == activeIndex }"
      />
    </div>

    <h1 v-if="imageSources.length == 0 && !isDragging">Drop some images</h1>

    <div class="manual">
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
    clickImage(img) {
      fetch(img)
        .then((res) => res.blob())
        .then((blob) => {
          const file = new File([blob], "dot.png", blob);
          let files = [];
          files.push(file);
          console.log(files);
          this.emitMethod(files);
        });
    },
    toggle: function(index) {
      this.activeIndex = index;
    },
    dragOver() {
      this.isDragging = true;
    },
    dragLeave() {
      this.isDragging = false;
    },
    async drop(e) {
      let files = [...e.dataTransfer.files];
      this.emitMethod(files);
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
.active {
  border: 2px solid green;
}
.drop {
  width: 100%;
  height: 100%;
  background-color: #eee;
  border: 10px solid #eee;
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 1rem;
  transition: background-color 0.2s ease-in-out;
  font-family: sans-serif;
  overflow: hidden;
  position: relative;
}
.isDragging {
  background-color: #999;
  border-color: #fff;
}
.img {
  padding: 1rem;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: 100%;
}
img {
  width: 100%;
  height: 100%;
  object-fit: contain;
}
.manual {
  position: absolute;
  bottom: 0;
  width: 100%;
  text-align: center;
  font-size: 0.8rem;
  text-decoration: underline;
}
#uploadmyfile {
  display: none;
}
</style>
