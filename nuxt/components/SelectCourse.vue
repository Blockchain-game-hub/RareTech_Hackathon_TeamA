<template>
  <div>
    <div class="flex flex-col mt-20">
      <div class="flex justify-center">
        <button class="nes-btn is-primary w-2/6 h-20 misaki text-2xl">
          記事300本ノック 初級
        </button>
      </div>
      <div class="flex justify-center mt-8">
        <button class="nes-btn is-primary w-2/6 h-20 misaki text-2xl">
          記事300本ノック 中級
        </button>
      </div>
      <div class="flex justify-center mt-8">
        <button class="nes-btn is-primary w-2/6 h-20 misaki text-2xl">
          記事300本ノック 上級
        </button>
      </div>
      <div class="flex justify-center mt-8">
        <button class="nes-btn is-primary w-2/6 h-20 misaki text-2xl">
          過去動画
        </button>
      </div>
      <div class="flex justify-center mt-20">
        <i class="gg-log-out"></i>
        <p class="press-start-2 mt-14">EXIT</p>
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
