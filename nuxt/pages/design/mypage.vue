<template>
  <div class="press-start-2 relative">
    <div class="absolute inset-y-3 left-20 w-8">
      <i class="gg-arrow-left-o"></i>
    </div>
    <div class="press-start-2p text-5xl text-center mt-10">MY PAGE</div>

    <Mypage />
  </div>
</template>

<script lang="ts">
import Vue from "vue";
import Mypage from "../../components/Mypage.vue";

export default Vue.extend({
  components: { Mypage },
});
</script>

<style scoped>
.gg-arrow-left-o {
  --ggs: 3;
}
</style>
