<template>
  <div>
    <div class="mt-16 press-start-2">
      <div class="flex justify-center">
        <div class="nes-container is-rounded bg-gray-200 text-center w-3/5">
          <p class="text-2xl">HOST</p>
          <p class="misaki text-2xl mt-6">USER NAME</p>
          <p class="misaki text-2xl mt-6">300本記事ノック 初級</p>
        </div>
      </div>

      <div class="mt-6 flex justify-center">
        <nuxt-Link to="game/draw" tag="button" class="nes-btn is-success"
          >GAME STRAT</nuxt-Link
        >
      </div>

      <div class="flex justify-center mt-4">
        <div class="nes-container is-rounded bg-white w-3/5">
          <div class="mt-5 text-center mr-48 flex justify-center h-30">
            <span class="text-xl mr-10">Room ID</span>
            <span class="mr-4">1111111111</span>
            <i class="gg-copy mt-2"></i>
          </div>
          <div class="flex justify-center ml-6">
            <div class="mt-3 text-center inline-flex">
              <span class="text-xl inline flex-1 mr-6">url</span>
              <p class="inline flex-1 mr-6">http://localhost:3000/standby</p>
              <i class="gg-copy mt-2 flex-1"></i>
            </div>
          </div>
        </div>
      </div>

      <div class="mt-10 flex justify-around">
        <div class="nes-container is-rounded bg-gray-200 text-center w-80">
          <p class="text-2xl">Player 1</p>
          <p class="misaki text-2xl mt-10">USER NAME</p>
        </div>
        <div class="nes-container is-rounded bg-gray-200 text-center w-80">
          <p class="text-2xl">Player 2</p>
          <p class="misaki text-2xl mt-10">USER NAME</p>
        </div>
        <div class="nes-container is-rounded bg-gray-200 text-center w-80">
          <p class="text-2xl">Player3</p>
          <p class="misaki text-2xl mt-10">USER NAME</p>
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
.gg-copy {
  --ggs: 1.2;
}
</style>
