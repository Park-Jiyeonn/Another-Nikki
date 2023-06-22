<script setup lang="ts">
import { ref } from 'vue'

import { User } from '@/api';
import { setCookies } from '@/hooks/useCookies';

const username = ref('')
const password = ref('')

const login = async () => {
    const ret = await User.login({username:username.value, password:password.value})
    console.log(ret.data.data.token)

    setCookies("token", ret.data.data.token)
}

</script>

<template>
    <el-input v-model="username" placeholder="Please input" />
    <el-input v-model="password" type="password" placeholder="Please input password" show-password />
    <el-button type="primary" @click="login">登录</el-button>
</template>