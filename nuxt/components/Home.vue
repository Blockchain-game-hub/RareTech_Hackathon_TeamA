<template>
  <div>
    <link
      rel="stylesheet"
      href="https://unpkg.com/nes.css@2.3.0/css/nes.min.css"
    />
    <div class="flex flex-col mt-32">
      <div class="flex justify-center">
        <button class="nes-btn is-primary w-2/6 h-20 press-start-2p">
          CREATE ROOM
        </button>
      </div>
      <div class="flex justify-center mt-16">
        <button class="nes-btn is-primary w-2/6 h-20 press-start-2p">
          JOIN ROOM
        </button>
      </div>
      <div class="mt-16 flex justify-center">
        <div class="flex-row">
          <button class="nes-btn is-primary w-60 h-14 m-10">MY PAGE</button>
          <button class="nes-btn is-primary w-60 h-14 m-10">LOG OUT</button>
        </div>
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
