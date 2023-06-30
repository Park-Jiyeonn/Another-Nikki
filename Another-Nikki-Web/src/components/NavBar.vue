<template>
    <el-menu mode="horizontal">
        <el-menu-item v-if="isLoggedIn" index="1" @click="router.push('/jelly')">
            Jelly
        </el-menu-item>
        <el-menu-item  v-if="isLoggedIn" index="3" @click="router.push('/logs')">
            Logs
        </el-menu-item>
        
        <el-menu-item index="2" @click="router.push('/runcode')">
            Run Code
        </el-menu-item>
        <el-menu-item index="4" @click="router.push('/articles')">
            Articles
        </el-menu-item>
    </el-menu>
</template>
  
<script lang="ts" setup>
import { useRouter } from 'vue-router'

import { useIsLoggedIn } from '@/hooks/userIsLogin'

const router = useRouter()
const isLoggedIn = useIsLoggedIn()
</script>