<template>
  <div>
    <div class="mt-32 press-start-2">
      <div class="flex justify-center">
        <div class="nes-container is-rounded bg-gray-200">
          <p class="text-center misaki text-2xl mt-6">
            ***** ルームIDを入力 *****
          </p>
          <div class="nes-field is-inline mt-10">
            <label for="inline_field ">RoomID</label>
            <input type="text" id="inline_field" class="nes-input is-success" />
            <button type="button" class="nes-btn is-success">Join</button>
          </div>
        </div>
      </div>
    </div>
    <div class="flex justify-center mt-32">
      <i class="gg-log-out"></i>
      <p class="press-start-2 mt-16">EXIT</p>
    </div>
  </div>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({});
</script>

<style scoped>
.gg-log-out {
  --ggs: 5;
}
</style>
