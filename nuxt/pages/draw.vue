<template>
  <div>
    <div class="">
      <h1 class="press-start-2 text-center text-5xl mt-10">Drawing</h1>
      <div class="flex">
        <Draw class="flex-1" />
        <Timer class="" />
      </div>

      <Theme />
    </div>
    <Player />
  </div>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({
  middleware: "auth",
});
</script>

<style scoped></style>
