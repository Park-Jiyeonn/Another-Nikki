<script setup lang="ts">
import NavBar from '@/components/NavBar.vue';

</script>

<template>
  <nav-bar></nav-bar>
  <router-view></router-view>
  <el-divider />
  <el-footer class="fa">
    <div>Copyright © 2024 by @Jiyeon. All rights reserved.</div>
    <div> Github: <a href="https://github.com/Park-Jiyeonn/Another-Nikki"> Jiyeon </a></div>
    <div> Email: jiyeonnn2@icloud.com </div>
  </el-footer>
</template>

<style>
@import url('https://cdnjs.cloudflare.com/ajax/libs/KaTeX/0.5.1/katex.min.css');
@import url('https://cdn.jsdelivr.net/github-markdown-css/2.2.1/github-markdown.css');
.fa {
  display: grid;
	grid-template-columns: repeat(auto-fit, var(--card-size));
	justify-content: center;
	grid-gap: var(--gap);
  font-size: 90%; align-items: center;
  margin-top: 30px;
  margin-bottom: 30px;
}
</style>
