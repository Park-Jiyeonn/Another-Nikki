<script setup lang="ts">
import { ElMessage } from 'element-plus';

import { ref } from 'vue'
import { ProblemApi } from '@/api';

const problem_name = ref('')
const problem_content = ref('')
const loading = ref(false);

const post_problem = async () => {
    const ret = await ProblemApi.post_problem({ 
        name: problem_name.value, 
        content:problem_content.value
    })
    // console.log(ret.data)
    if (ret.data.code == 200) {
        ElMessage.success(ret?.data?.message ?? 'success')
    }
    else {
        return ElMessage.error("发布失败！")
    }
}


</script>
<template>
    <el-input v-model="problem_name" :autosize="{ minRows: 2, maxRows: 2 }" type="textarea" placeholder="Please input"
        style="margin-top: 10px;" />
    <el-input v-model="problem_content" :autosize="{ minRows: 10, maxRows: 10 }" type="textarea" placeholder="Please input"
        style="margin-top: 10px;" />

    <div style="text-align: center; margin-top: 10px;">
        <el-button class="button" type="primary" :loading="loading" @click="post_problem()">
            发布
        </el-button>
    </div>
</template>