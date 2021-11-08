<template>
  <div class="flex justify-center mt-20">
    <div class="flex flex-col mt-20 nes-container is-rounded is-dark w-5/6">
      <div class="flex justify-center mt-8">
        <nuxt-Link
          to="/standby"
          tag="button"
          class="nes-btn is-primary w-2/6 h-20 misaki text-2xl"
        >
          記事300本ノック 初級
        </nuxt-Link>
      </div>
      <div class="flex justify-center mt-8">
        <nuxt-Link
          to="/standby"
          tag="button"
          class="nes-btn is-primary w-2/6 h-20 misaki text-2xl"
        >
          記事300本ノック 中級
        </nuxt-Link>
      </div>
      <div class="flex justify-center mt-8">
        <nuxt-Link
          to="/standby"
          tag="button"
          class="nes-btn is-primary w-2/6 h-20 misaki text-2xl"
        >
          記事300本ノック 上級
        </nuxt-Link>
      </div>
      <div class="flex justify-center mt-8">
        <nuxt-Link
          to="/standby"
          tag="button"
          class="nes-btn is-primary w-2/6 h-20 misaki text-2xl mb-8"
        >
          過去動画
        </nuxt-Link>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({});
</script>

<style scoped>
.gg-log-out {
  --ggs: 5;
}
</style>
