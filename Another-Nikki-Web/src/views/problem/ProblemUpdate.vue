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
    problem_id: 1,
    problem_title: "name",
    problem_content: "content",
    created_time: "CreatedAt",
})
const problem_name = ref('')
const problem_content = ref('')
const loading = ref(false);

const get_problem = async (id: number) => {
    const ret = await ProblemApi.get_problem({ problem_id: id })
    if (ret.data.code == 200) {
        problem.value = ret.data.data
        problem_name.value = problem.value.problem_title
        problem_content.value = problem.value.problem_content
    }
    else {
        return ElMessage.error("获取内容失败！")
    }
}

const id: number = parseInt(String(route.params.id));
get_problem(id)

const update_problem = async () => {
    const ret = await ProblemApi.update_problem({ 
        problem_id: id, 
        problem_title: problem_name.value, 
        problem_content:problem_content.value
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