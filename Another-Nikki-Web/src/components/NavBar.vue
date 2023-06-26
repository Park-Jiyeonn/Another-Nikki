<template>
    <el-menu class="el-menu-demo" mode="horizontal">
        <el-menu-item index="1" @click="router.push('/jelly')">
            Jelly
        </el-menu-item>
        <el-menu-item index="2" @click="router.push('/runcode')">
            Run Code 
        </el-menu-item>
        <el-menu-item index="4" @click="router.push('/articles')">
            Articles
        </el-menu-item>
        <el-menu-item index="3" @click="router.push('/logs')">
            Logs
        </el-menu-item>
    </el-menu>
</template>
  
<script lang="ts" setup>
import { useRouter } from 'vue-router'
const router = useRouter()
</script>