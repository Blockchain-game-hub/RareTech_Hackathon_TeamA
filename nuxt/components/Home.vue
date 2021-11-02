<template>
  <div>
    <div class="flex flex-col mt-32">
      <div class="flex justify-center">
        <button
          class="nes-btn is-primary w-2/6 h-20 press-start-2p press-start-2"
        >
          CREATE ROOM
        </button>
      </div>
      <div class="flex justify-center mt-16">
        <button
          class="nes-btn is-primary w-2/6 h-20 press-start-2p press-start-2"
        >
          JOIN ROOM
        </button>
      </div>
      <div class="mt-16 flex justify-center">
        <div class="flex-row">
          <button class="nes-btn is-primary w-60 h-14 m-10 press-start-2">
            MY PAGE
          </button>
          <button class="nes-btn is-primary w-60 h-14 m-10 press-start-2">
            LOG OUT
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({});
</script>

<style scoped></style>
