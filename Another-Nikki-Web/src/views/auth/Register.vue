<script setup lang="ts">
import { ref } from 'vue'
import { ElMessage } from 'element-plus';

import { User } from '@/api';
import { setCookies } from '@/hooks/useCookies';
import ContentBase from '@/components/ContentBase.vue';

const username = ref('')
const password1 = ref('')
const password2 = ref('')

const register = async () => {
    const ret = await User.register({username:username.value, password:password1.value, confirm_password:password2.value})
    if (ret.data.code != 200) {
        return ElMessage.warning(ret.data.message ?? 'Something went wrong, please try again.')
    }
    else {
        setCookies("token", ret.data.data.token)
        setCookies("user_id", ret.data.data.user_id)
        setCookies("username", ret.data.data.username)
        return ElMessage.success("注册成功，为您自动跳转")
    }
}

</script>

<template>
    <ContentBase>
        <el-input v-model="username" placeholder="Please input" />
        <el-input v-model="password1" type="password" placeholder="Please input password" show-password />
        <el-input v-model="password2" type="password" placeholder="Please confirm password" show-password />
        <el-button type="primary" @click="register">注册</el-button>
    </ContentBase>
</template>