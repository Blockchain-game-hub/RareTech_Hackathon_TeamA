<template>
  <div class="flex justify-center mt-10">
    <div class="flex flex-col mt-32 nes-container is-rounded is-dark w-5/6">
      <nuxt-Link tag="div" to="/select" class="flex justify-center">
        <button
          class="
            nes-btn
            is-primary
            w-2/6
            h-20
            press-start-2p press-start-2
            mt-8
          "
        >
          CREATE ROOM
        </button>
      </nuxt-Link>
      <nuxt-Link tag="div" to="/join" class="flex justify-center mt-16 -mb-8">
        <button
          class="nes-btn is-primary w-2/6 h-20 press-start-2p press-start-2"
        >
          JOIN ROOM
        </button>
      </nuxt-Link>
      <div class="mt-16 flex justify-center">
        <div class="flex-row">
          <nuxt-Link
            tag="button"
            to="/mypage"
            class="nes-btn is-primary w-60 h-14 m-10 press-start-2"
          >
            MY PAGE
          </nuxt-Link>

          <nuxt-Link
            tag="button"
            to="/login"
            class="nes-btn is-primary w-60 h-14 m-10 press-start-2"
          >
            LOG OUT
          </nuxt-Link>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({});
</script>

<style scoped></style>
