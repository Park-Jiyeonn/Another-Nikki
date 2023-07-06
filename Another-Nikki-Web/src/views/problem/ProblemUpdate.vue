<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useRoute } from 'vue-router';
import { ElMessage } from 'element-plus';

import { ref } from 'vue'
import { Problem } from '@/types/Problem';
import { ProblemApi } from '@/api';


const route = useRoute();
const router = useRouter();
const problem = ref<Problem>({
    ID: 1,
    name: "name",
    content: "content",
    CreatedAt: "CreatedAt",
})
const problem_name = ref('')
const problem_content = ref('')
const loading = ref(false);

const get_problem = async (id: number) => {
    const ret = await ProblemApi.get_problem({ ID: id })
    if (ret.data.code == 200) {
        problem.value = ret.data.data
        problem_name.value = problem.value.name
        problem_content.value = problem.value.content
    }
    else {
        return ElMessage.error("获取内容失败！")
    }
    
}

const id: number = parseInt(String(route.params.id));
get_problem(id)

const update_problem = async () => {
    const ret = await ProblemApi.update_problem({ 
        ID: id, 
        name: problem_name.value, 
        content:problem_content.value
    })
    // console.log(ret.data)
    if (ret.data.code == 200) {
        ElMessage.success(ret?.data?.message ?? 'success')
        return router.push(`/problem/${id}`)
    }
    else {
        return ElMessage.error("更新失败！")
    }
}


</script>
<template>
    <el-input v-model="problem_name" :autosize="{ minRows: 2, maxRows: 2 }" type="textarea" placeholder="Please input"
        style="margin-top: 10px;" />
    <el-input v-model="problem_content" :autosize="{ minRows: 10, maxRows: 10 }" type="textarea" placeholder="Please input"
        style="margin-top: 10px;" />

    <div style="text-align: center; margin-top: 10px;">
        <el-button class="button" type="primary" :loading="loading" @click="update_problem()">
            更新
        </el-button>
    </div>
</template>