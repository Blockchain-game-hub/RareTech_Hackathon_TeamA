<template>
  <div>
    <BackButton />
    <h1 class="press-start-2 text-center text-5xl mt-10">StandBy ...</h1>

    <StandBy />
  </div>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({
  middleware: "auth",
});
</script>

<style scoped>
.gg-arrow-left-o {
  --ggs: 3;
}
</style>
