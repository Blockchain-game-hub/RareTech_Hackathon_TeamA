<template>
  <div class="flex justify-center mt-2 mr-16 -ml-20">
    <div div class="nes-container is-rounded is-dark h-20 w-36">
      <p class="text-center press-start-2 text-5xl">{{ timer }}</p>
    </div>
  </div>
</template>

<script lang="ts">
import { ref, defineComponent, onMounted } from "@nuxtjs/composition-api";

export default defineComponent({
  setup(pops, context) {
    const timer = ref<number>(60);
    const round = ref<number>(0);
    const getSessionStorage = sessionStorage.getItem("round");
    console.log(getSessionStorage);

    onMounted(() => {
      const countdownTimer = () => {
        if (timer.value > 0) {
          timer.value--;
          setTimeout(countdownTimer, 1000);
        } else if (timer.value === 0 && getSessionStorage == "1") {
          context.emit("modal");
          sessionStorage.setItem("round", "2");
        } else if (timer.value === 0 && getSessionStorage == "2") {
          context.emit("modal");
          sessionStorage.setItem("round", "3");
        } else if (timer.value === 0 && getSessionStorage == "3") {
          sessionStorage.setItem("round", "3");
          sessionStorage.removeItem("round");
          context.emit("resultModal");
        } else {
          context.emit("modal");
          sessionStorage.setItem("round", "1");
        }
      };
      countdownTimer();
    });
    return {
      timer,
      round,
    };
  },
});
</script>

<style scoped>
.double-circle-3 {
  width: 300px;
  height: 300px;
  position: relative;
  border: solid #fa8072 15px;
  border-radius: 50%;
  box-sizing: border-box;
}
</style>
