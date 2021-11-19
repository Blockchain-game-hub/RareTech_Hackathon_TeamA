<template>
  <div>
    <div class="flex justify-center mt-10">
      <div>
        <div class="double-circle-3"></div>
        <p class="text-5xl misaki text-center ml-5 mt-3">正解！！</p>

        <div class="flex justify-center mt-10 press-start-2">
          <nuxt-Link
            tag="button"
            to="/quiz_answer"
            class="nes-btn is-primary mr-5 w-32"
            >next</nuxt-Link
          >
        </div>
      </div>
    </div>
  </div>
</template>
