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
    rgba(31, 173, 192, 0.87),
    rgba(22, 145, 38, 0.61)
  );
}
</style>
