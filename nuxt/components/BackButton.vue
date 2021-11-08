<template>
  <div @click="backToPage" class="press-start-2 relative back-button">
    <div class="absolute inset-y-3 left-20 w-8">
      <i class="gg-arrow-left-o"></i>
    </div>
  </div>
</template>

<script>
export default {
  methods: {
    backToPage() {
      this.$router.back();
    },
  },
};
</script>

<style scoped>
.gg-arrow-left-o {
  --ggs: 3;
}

.back-button:hover {
  opacity: 0.8;
}
.back-button:active {
  -webkit-transform: translate(0, 2px);
  -moz-transform: translate(0, 2px);
  transform: translate(0, 2px);
}
</style>
