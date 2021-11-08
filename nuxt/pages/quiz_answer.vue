<template>
  <div>
    <div class="">
      <h1 class="press-start-2 text-center text-5xl mt-10">Question</h1>
      <div class="flex">
        <DrawQuizBoard class="flex-1" />
        <Timer class="" />
      </div>
      <InputAnswer />
    </div>
    <Player />
  </div>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({});
</script>

<style scoped></style>
