<script setup lang="ts">
import { ref } from 'vue'

import { User } from '@/api';

const username = ref('')
const password1 = ref('')
const password2 = ref('')

const register = () => {
    User.register({username:username.value, password1:password1.value, password2:password2.value})
}

</script>

<template>
    <el-input v-model="username" placeholder="Please input" />
    <el-input v-model="password1" type="password" placeholder="Please input password" show-password />
    <el-input v-model="password2" type="password" placeholder="Please confirm password" show-password />
    <el-button type="primary" @click="register">注册</el-button>
</template>