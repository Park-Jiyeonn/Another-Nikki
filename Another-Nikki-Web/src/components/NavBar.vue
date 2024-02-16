<template>
    <el-menu mode="horizontal" :ellipsis="false">
        <el-menu-item v-if="isLoggedIn" index="1" @click="router.push('/jelly')">
            Jelly
        </el-menu-item>
        <el-menu-item index="2" @click="router.push('/problems')">
            Problem
        </el-menu-item>
        <el-menu-item index="3" @click="router.push('/articles')">
            Articles
        </el-menu-item>
        <div class="flex-grow" />
        <el-menu-item index="4" @click="router.push('/profile')">
            <el-avatar src="https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png"/>
        </el-menu-item>
    </el-menu>
</template>
  
<script lang="ts" setup>
import { useRouter } from 'vue-router'

import { useIsLoggedIn } from '@/hooks/userIsLogin'

const router = useRouter()
const isLoggedIn = useIsLoggedIn()
</script>

<style>
.flex-grow {
  flex-grow: 1;
}
</style>