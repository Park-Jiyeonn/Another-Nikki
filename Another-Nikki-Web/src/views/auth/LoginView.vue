<script setup lang="ts">
import { ref } from 'vue'
import { ElMessage } from 'element-plus';

import { User } from '@/api';
import { setCookies } from '@/hooks/useCookies';
import router from '@/router';
import ContentBase from '@/components/ContentBase.vue';

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
        ElMessage.success("登录成功")
        router.push(`/problems/`).then(() => {
            // 跳转后刷新页面
            window.location.href = '/problems/';
        });
    }
    else {
        return ElMessage.warning(ret.data.message ?? 'Something went wrong, please try again.')
    }
}

const create_tourist_account = async () => {
    const ret = await User.create_tourist_account({})
    if (ret.data.code == 200) {
        setCookies("token", ret.data.data.token)
        setCookies("user_id", ret.data.data.user_id)
        setCookies("username", ret.data.data.username)
        setCookies("avatar", ret.data.data.avatar)
        setCookies("description", ret.data.data.description) 
        ElMessage.success("一键登录成功")
        router.push(`/problems/`).then(() => {
            // 跳转后刷新页面
            window.location.href = '/problems/';
        });
    }
    else {
        return ElMessage.warning(ret.data.message ?? 'Something went wrong, please try again.')
    }
}

</script>

<template>
    <ContentBase>
        <el-input v-model="username" placeholder="Please input" />
        <el-input v-model="password" type="password" placeholder="Please input password" show-password />
        <el-button type="primary" @click="login">登录</el-button>
        <el-button type="primary" @click="create_tourist_account">一键登录</el-button>
    </ContentBase>
</template>