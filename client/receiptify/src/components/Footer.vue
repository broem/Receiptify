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
    rgba(14, 16, 82, 0.616),
    rgba(25, 19, 56, 0.349)
  );
}
</style>
