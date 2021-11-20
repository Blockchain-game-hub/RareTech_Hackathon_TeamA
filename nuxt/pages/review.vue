<template>
  <div>
    <h1 class="press-start-2 text-center text-5xl mt-10 misaki">
      間違えた問題
    </h1>

    <div>
      <Review />
    </div>
    <div class="flex justify-center mt-16">
      <div class="inline-flex space-x-32">
        <nuxt-link to="/">
          <div class="flex-1">
            <div class="flex justify-center">
              <i class="gg-home-alt mt-4 items-center"></i>
            </div>
            <br />
            <p class="w-32 text-center mt-4 press-start-2 text-white text-xl">
              Home
            </p>
          </div>
        </nuxt-link>

        <nuxt-Link to="/standby" class="flex-1">
          <div class="flex justify-center">
            <i class="gg-redo"></i>
          </div>
          <br />
          <p class="w-28 text-center mt-7 press-start-2 text-white text-xl">
            Rematch
          </p>
        </nuxt-Link>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({
  middleware: "auth",
});
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
