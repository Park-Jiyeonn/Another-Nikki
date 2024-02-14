<script setup lang="ts">
import { ref } from 'vue'
import { Problem } from '@/types/Problem';
import { ProblemApi } from '@/api';
import router from '@/router';
const problems = ref<Problem[]>([])

const problem_page_que = async (page_num: number, page_size: number) => {
    const ret = await ProblemApi.page_que({ 
        page_num,
        page_size,
     })
    problems.value = ret.data.data.problems
}
problem_page_que(10, 1)

const toProblemDetail = (id: number) => {
    console.log(id)
    router.push({ path: `/problem/${id}` })
}

</script>

<template>
    <div class="container">
        <el-card class="problem-card" v-for="problem in problems" @click="toProblemDetail(problem.problem_id)">
            <span> {{ problem.problem_title }} </span>
            <span class="date">{{ problem.created_time }}</span>
        </el-card>
    </div>
</template>

<style scoped>
.container {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-wrap: wrap;
}

.problem-card {
    cursor: pointer;
    width: 960px;
}

.problem-card:hover {
    color: chocolate;
}

.date {
    list-style: none;
    padding: 0px;
    color: var(--el-text-color-secondary);
    font-size: var(--el-font-size-small);
    flex-direction: row;
    float: right;
}</style>