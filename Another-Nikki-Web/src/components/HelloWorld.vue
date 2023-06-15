<script setup lang="ts">
defineProps<{ msg: string }>()
</script>

<template>
  <h1 style="text-align: center;">
    {{ msg }}
  </h1>
</template>

<style scoped>
</style>
