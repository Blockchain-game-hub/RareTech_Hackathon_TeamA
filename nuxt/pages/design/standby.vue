<template>
  <div>
    <div class="press-start-2 relative">
      <div class="absolute inset-y-3 left-20 w-8">
        <i class="gg-arrow-left-o"></i>
      </div>
      <h1 class="press-start-2 text-center text-5xl mt-10">StandBy ...</h1>
    </div>
    <StandBy />
  </div>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({});
</script>

<style scoped>
.gg-arrow-left-o {
  --ggs: 3;
}
</style>
