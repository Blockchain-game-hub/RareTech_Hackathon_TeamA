<template>
  <div>
    <!-- モーダル実験 -->
    <modal name="modal" class="nes-container is-rounded press-start-2">
      <p>modal</p>
      <p>modal</p>
      <p>modal</p>
      <p>modal</p>
    </modal>
    <!--  -->
    <div class="">
      <h1 class="press-start-2 text-center text-5xl mt-10">Drawing</h1>
      <button @click="show">aaaa</button>
      <div class="flex">
        <Draw class="flex-1" />
        <Timer class="" @modal="showModal" />
      </div>

      <Theme />
    </div>
    <Player />
  </div>
</template>

<script lang="ts">
import { defineComponent } from "@nuxtjs/composition-api";

export default defineComponent({
  methods: {
    showModal() {
      this.$modal.show("modal");
    },
  },

  setup(props, context) {
    return {};
  },
});
</script>

<style scoped></style>
