<template>
  <div class="flex justify-center mt-2 mr-16 -ml-20">
    <div div class="nes-container is-rounded is-dark h-20 w-36">
      <p class="text-center press-start-2 text-5xl">{{ timer }}</p>
    </div>
  </div>
</template>

<script>
import Vue from "vue";

export default Vue.extend({
  data() {
    return {
      timer: 60,
    };
  },
  mounted() {
    const countdownTimer = () => {
      if (this.timer > 0) {
        this.timer--;
        setTimeout(countdownTimer, 1000);
      }
    };
    countdownTimer();
  },
});
</script>

<style scoped></style>
