<template>
  <div>
    <!-- モーダル -->
    <modal name="modal" width="600px" height="500px" :clickToClose="false">
      <div class="flex justify-center mt-10">
        <div>
          <div class="double-circle-3"></div>
          <p class="text-5xl misaki text-center ml-5 mt-3">正解！！</p>

          <div class="flex justify-center mt-10 press-start-2">
            <button
              @click="nextTransition"
              class="nes-btn is-primary mr-5 w-32"
            >
              next
            </button>
          </div>
        </div>
      </div>
    </modal>
    <!--  -->
    <!--resultModal  -->
    <modal
      name="resultModal"
      width="600px"
      height="500px"
      :clickToClose="false"
    >
      <div class="flex justify-center mt-10">
        <div>
          <div class="double-circle-3"></div>
          <p class="text-5xl misaki text-center ml-5 mt-3">正解！！</p>

          <div class="flex justify-center mt-10 press-start-2">
            <nuxt-Link
              tag="button"
              to="/result"
              class="nes-btn is-warning ml-5 w-32"
              >Result</nuxt-Link
            >
          </div>
        </div>
      </div>
    </modal>
    <!--  -->
    <div class="">
      <h1 class="press-start-2 text-center text-5xl mt-10">Question</h1>
      <div class="flex">
        <DrawQuizBoard class="flex-1" />
        <Timer class="" @modal="showModal" @resultModal="showResultModal" />
      </div>
      <InputAnswer />
    </div>
    <Player />
  </div>
</template>

<script>
import { defineComponent } from "@nuxtjs/composition-api";

export default defineComponent({
  middleware: "auth",
  methods: {
    showModal() {
      this.$modal.show("modal");
    },
    showResultModal() {
      this.$modal.show("resultModal");
    },
    nextTransition() {
      if (this.$route.params.id == 1) {
        this.$router.push({ params: { id: 2 } });
      } else if (this.$route.params.id == 2) {
        this.$router.push({ params: { id: 3 } });
      } else if (this.$route.params.id == 3) {
        this.$router.push({ params: { id: 4 } });
      } else if (this.$route.params.id == 4) {
        this.$router.push("draw");
      }
    },
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
