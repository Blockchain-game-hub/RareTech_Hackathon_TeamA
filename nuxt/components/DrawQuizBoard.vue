<template>
  <div>
    <div class="relative">
      <main>
        <div class="flex justify-center ml-32">
          <canvas
            class="bg-white nes-container is-rounded"
            id="gameCanvas"
            ref="gameCanvas"
            @mousemove="draw"
            @mousedown="dragStart"
            @mouseup="dragEnd"
            @mouseout="dragEnd"
          ></canvas>
        </div>
      </main>
    </div>
  </div>
</template>

<script lang="ts">
import {
  defineComponent,
  onMounted,
  reactive,
  ref,
} from "@nuxtjs/composition-api";

export default defineComponent({
  setup() {
    const canvasDefultStatus = {
      width: 900,
      height: 400,
    };

    onMounted(() => {
      const canvas = <HTMLCanvasElement>document.getElementById("gameCanvas");
      canvas.width = canvasDefultStatus.width;
      canvas.height = canvasDefultStatus.height;
    });
  },
});
</script>
