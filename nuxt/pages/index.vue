<template>
  <div>
    <div>
      <h1 class="text-center text-5xl mt-10 press-start-2">GAMETITLE</h1>
      <Home class="anim-box" />
    </div>

    <div class="h-44 relative w-screen">
      <div class="press-start-2 bg-gray-400">
        <div class="nes-balloon from-left absolute left-40 bottom-50 bg-white">
          <p class="">Hello World!!</p>
        </div>
        <i class="nes-octocat animate absolute bottom-3 left-20"></i>
      </div>
      <section class="message -right fadeIn">
        <!-- Balloon -->

        <div
          class="
            nes-balloon
            from-right
            absolute
            right-40
            bottom-50
            press-start-2
          "
        >
          <p class="w-96">Thou hast had a good night's sleep, I hope.</p>
        </div>
        <nuxt-link to="/quiz_answer">
          <i class="nes-kirby absolute bottom-3 right-20"></i>
        </nuxt-link>
      </section>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({});
</script>

<style scoped>
.fadeIn {
  animation-name: fadeInAnime;
  animation-fill-mode: backwards;
  animation-duration: 5s;
  animation-timing-function: ease;
  animation-delay: 0.5s;
  animation-direction: normal;
}

/*1で解説*/
@keyframes fadeInAnime {
  0% {
    opacity: 0;
  }

  100% {
    opacity: 1;
  }
}

.anim-box {
  animation: zoomIn 1.5s cubic-bezier(0.25, 1, 0.5, 1) 1 forwards;
}

@keyframes zoomIn {
  0% {
    transform: scale(0.8);
    opacity: 0;
  }
  100% {
    opacity: 1;
    transform: scale(1);
  }
}
</style>
