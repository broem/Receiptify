<template>
  <div class="foot">
    Developed with ❤️ by
    <a href="https://www.github.com/broem" target="_blank">Broem</a>
    &copy; {{ new Date().getFullYear() }}
  </div>
</template>

<style>
.foot {
  position: absolute;
  text-align: right;
  bottom: 0%;
  width: 100%;
  background-image: linear-gradient(
    to right bottom,
    rgba(6, 7, 20, 0.808),
    rgba(3, 14, 8, 0.637)
  );
}
</style>
