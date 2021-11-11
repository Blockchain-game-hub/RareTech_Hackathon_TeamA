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

    onMounted(() => {
      const countdownTimer = () => {
        if (timer.value > 0) {
          timer.value--;
          setTimeout(countdownTimer, 100);
        } else if (timer.value === 0) {
          context.emit("modal");
        }
      };
      countdownTimer();
    });
    return {
      timer,
    };
  },
});
</script>

<style scoped></style>
