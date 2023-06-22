<script setup lang="ts">
import NavBar from '@/components/NavBar.vue';

console.log(import.meta.env.VITE_API_URL);
</script>

<template>
  <nav-bar></nav-bar>
  <router-view></router-view>
</template>

<style scoped>
.logo {
  height: 4em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}

.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}

.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
