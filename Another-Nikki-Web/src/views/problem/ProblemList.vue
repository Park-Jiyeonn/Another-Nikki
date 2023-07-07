<script setup lang="ts">
import { ref } from 'vue'
import { Problem } from '@/types/Problem';
import { ProblemApi } from '@/api';
import router from '@/router';
const problems = ref<Problem[]>([])

const problem_page_que = async (page: number) => {
    const ret = await ProblemApi.page_que({ page: page })
    problems.value = ret.data.data
    problems.value.forEach((problem) => {
        problem.CreatedAt = problem.CreatedAt.substring(0, 10) + " " +
            problem.CreatedAt.substring(11, 16)
    });
}
problem_page_que(1)

const toProblemDetail = (id: number) => {
    console.log(id)
    router.push({ path: `/problem/${id}` })
}

</script>

<template>
    <div class="container">
        <el-card class="problem-card" v-for="problem in problems" @click="toProblemDetail(problem.ID)">
            <span> {{ problem.name }} </span>
            <span class="date">{{ problem.CreatedAt }}</span>
        </el-card>
    </div>
</template>

<style scoped>
.container {
    display: flex;
    justify-content: center;
    align-items: center;
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