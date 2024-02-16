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
            <el-dropdown trigger="click">
                <el-avatar :src="user_avatar"/>
                <template #dropdown> <el-dropdown-menu>
                    <el-dropdown-item> 我的主页 </el-dropdown-item>
                    <el-dropdown-item @click="logout()"> 退出登录 </el-dropdown-item>
                  </el-dropdown-menu> </template>
            </el-dropdown>
        </el-menu-item>
    </el-menu>
</template>
  
<script lang="ts" setup>
import { ref, watchEffect } from 'vue'
import { useRouter } from 'vue-router'
import { useIsLoggedIn } from '@/hooks/userIsLogin'
import { getCookies, removeCookies } from "@/hooks/useCookies";
const router = useRouter()
const isLoggedIn = useIsLoggedIn()
const user_avatar = ref(getCookies("avatar"))
const logout = async () => {
    removeCookies("username")
    removeCookies("avatar")
    removeCookies("user_id")
    removeCookies("token")
    location.reload();
}
watchEffect(() => {
  user_avatar.value = getCookies("avatar");
});
</script>

<style>
.flex-grow {
  flex-grow: 1;
}
</style>