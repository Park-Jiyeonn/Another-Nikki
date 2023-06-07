<script setup lang="ts">
import 'bootstrap/dist/css/bootstrap.min.css';
// import 'bootstrap/dist/ts/bootstrap';

import HelloWorld from './components/HelloWorld.vue'
import PostBlog from './components/PostBlog.vue';
import ContentBase from './components/ContentBase.vue'
</script>

<template>
  <div>
    <img src="/vite.svg" class="logo" alt="Vite logo" />
    <img src="./assets/vue.svg" class="logo vue" alt="Vue logo" />
  </div>
  <HelloWorld msg="网站雏形" />
  <ContentBase>
    <PostBlog/>
  </ContentBase>
</template>

<style scoped>
.logo {
  height: 6em;
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
