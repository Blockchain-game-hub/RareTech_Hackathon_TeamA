<template>
  <div>
    <link
      rel="stylesheet"
      href="https://unpkg.com/nes.css@2.3.0/css/nes.min.css"
    />

    <div class="absolute inset-y-3 left-10 w-8">
      <button type="button" class="nes-btn w-32 h-24">Return to Home</button>
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

<style>
@import url("https://fonts.googleapis.com/css2?family=Press+Start+2P&display=swap");
.press-start-2p {
  font-family: "Press Start 2P", cursive;
}
.body-class {
  background-color: #409fd5;
}
</style>
