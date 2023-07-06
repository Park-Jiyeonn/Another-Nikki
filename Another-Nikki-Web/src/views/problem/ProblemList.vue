<script setup lang="ts">
import { ref } from 'vue'
import { Problem } from '@/types/Problem';
import { ProblemApi } from '@/api';
import router from '@/router';
const problems = ref<Problem[]>([])

const problem_page_que = async (page: number) => {
    const ret = await ProblemApi.page_que({page:page})
    problems.value = ret.data.data
    problems.value.forEach((problem) => {
        problem.CreatedAt = problem.CreatedAt.substring(0, 10) + " " +
                            problem.CreatedAt.substring(11, 16)
    });
}
problem_page_que(1)

const toProblemDetail = (id:number) => {
    console.log(id)
    router.push({path: `/problem/${id}`})
}

</script>

<template>
    <el-card class="problem-card" v-for="problem in problems" @click="toProblemDetail(problem.ID)">
        <el-row :gutter="24">
            <el-col :span="18">
                <h3>
                    {{ problem.name }}
                </h3>
                <ul class="problem-footer">
                    <li class="problem-footer-li">
                        <span class="li-title">{{problem.CreatedAt}}</span>
                    </li>
                </ul>
            </el-col>
        </el-row>
    </el-card>
</template>

<style scoped>
.problem-card {
  cursor: pointer;
}

.problem-description {
    display: -webkit-box;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 2;
    overflow: hidden;
    margin-top: 5px;
    font-size: 14px;
}

.problem-footer {
    list-style: none;
    padding: 0px;
    color: var(--el-text-color-secondary);
    font-size: var(--el-font-size-small);
    display: inline-block;
    flex-direction: row;

    .li-title {
        vertical-align: middle;
    }
}
</style>