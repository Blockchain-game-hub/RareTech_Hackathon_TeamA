<template>
  <div>
    <BackButton />
    <div
      v-if="User != null"
      class="press-start-2 text-5xl text-center mt-10"
    >
      {{ User.displayName }} PAGE
      </div>

    <Mypage />
  </div>
</template>

<script lang="ts">
import Vue from "vue";

export default Vue.extend({
  middleware: "auth",
  data() {
    return {
      User: this.$store.state.authUser
    }
  }
});
</script>
