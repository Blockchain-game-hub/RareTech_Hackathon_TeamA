<template>
  <div>
    <h1 class="press-start-2 text-center text-5xl mt-10">Result</h1>

    <div>
      <Result />
    </div>
    <div class="flex justify-center mt-16">
      <div class="inline-flex space-x-32">
        <div class="flex-1">
          <nuxt-Link to="/home" class="flex justify-center">
            <i class="gg-home-alt mt-4 items-center"></i>
          </nuxt-Link>
          <br />
          <p class="w-32 text-center mt-4 press-start-2 text-white text-xl">
            Home
          </p>
        </div>
        <nuxt-Link to="/standby" class="flex-1">
          <div class="flex justify-center">
            <i class="gg-redo"></i>
          </div>
          <br />
          <p class="w-28 text-center mt-7 press-start-2 text-white text-xl">
            Rematch
          </p>
        </nuxt-Link>
        <nuxt-Link to="/review" class="flex-1">
          <div class="flex justify-center">
            <i class="gg-notes"></i>
          </div>
          <br />
          <p class="w-36 text-center mt-6 misaki text-white text-xl">
            間違えた問題を復習
          </p>
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
.gg-home-alt {
  --ggs: 5;
}
.gg-redo {
  --ggs: 5;
}
.gg-notes {
  --ggs: 5;
}
</style>
