<script setup lang="ts">
import { ref } from 'vue'
import { ElMessage } from 'element-plus';

import { User } from '@/api';
import { setCookies } from '@/hooks/useCookies';
import router from '@/router';

const username = ref('')
const password = ref('')

const login = async () => {
    const ret = await User.login({username:username.value, password:password.value})
    if (ret.data.code == 200) {
        setCookies("token", ret.data.data.token)
        setCookies("user_id", ret.data.data.user_id)
        setCookies("username", ret.data.data.username)
        setCookies("avatar", ret.data.data.avatar)
        setCookies("description", ret.data.data.description)
        return ElMessage.success("登录成功")
    }
    else {
        return ElMessage.warning(ret.data.message ?? 'Something went wrong, please try again.')
    }
    // setCookies("token", ret.data.data.token)
}

</script>

<template>
    <el-input v-model="username" placeholder="Please input" />
    <el-input v-model="password" type="password" placeholder="Please input password" show-password />
    <el-button type="primary" @click="login">登录</el-button>
</template>