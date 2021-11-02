<template>
  <div class="flex flex-col mt-32">
    <link
      rel="stylesheet"
      href="https://unpkg.com/nes.css@2.3.0/css/nes.min.css"
    />
    <div class="flex justify-center">
      <div class="nes-container is-dark with-title w-1/2">
        <p class="title">Match Results</p>
        <p>Battles</p>
        <p>1st</p>
        <p>2nd</p>
      </div>
    </div>

    <div class="flex justify-center mt-16">
      <div class="nes-container is-dark with-title w-1/2">
        <p class="title">Total Point</p>
        <p>Point</p>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({});
</script>

<style scoped>
@import url("https://fonts.googleapis.com/css2?family=Press+Start+2P&display=swap");
.press-start-2p {
  font-family: "Press Start 2P", cursive;
}
</style>
