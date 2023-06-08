<script setup lang="ts">
import HelloWorld from './components/HelloWorld.vue'
import PostBlog from './components/PostBlog.vue';
import ContentBase from './components/ContentBase.vue'

console.log(import.meta.env.VITE_API_URL);
</script>

<template>
  <ContentBase>
    <el-row>
      <el-col :span="6">
        <!-- <img src="/vite.svg" class="logo" alt="Vite logo" /> -->
      </el-col>
      <el-col :span="8">
        <img src="/vite.svg" class="logo" alt="Vite logo" />
      </el-col>
      <el-col :span="8">
        <img src="./assets/vue.svg" class="logo vue" alt="Vue logo" />
      </el-col>
    </el-row>

    <HelloWorld msg="Another Nikki"/>

    <PostBlog />
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
