<script setup lang="ts">
import { useRoute } from 'vue-router';
import { useRouter } from 'vue-router'
import { ref } from 'vue';

import { Problem } from '@/types/Problem';
import { ProblemApi } from '@/api';
// import 'highlight.js/styles/github-dark-dimmed.css'
import { renderMarkdown } from '@/utils/markdown'
import { useIsLoggedIn } from '@/hooks/userIsLogin'
import RunCode from '@/components/RunCode.vue';
import ContentBase from '@/components/ContentBase.vue';
import { getCookies } from "@/hooks/useCookies";

const isLoggedIn = useIsLoggedIn()
const route = useRoute();
const router = useRouter();
const user_id : number = getCookies("user_id")

const problem = ref<Problem>({
    problem_id: 1,
    problem_title: "",
    problem_content: "",
})

const get_problem = async (id: number) => {
    const ret = await ProblemApi.get_problem({ problem_id: id })
    if (ret.data.code < 200 || ret.data.code > 200) {
        problem.value.problem_id = -1
        problem.value.problem_title = "请求错误捏, 请不要填奇奇怪怪的参数哦"
        problem.value.problem_content = "空"
        return
    }
    problem.value = ret.data.data
}

const id: number = parseInt(String(route.params.id));
get_problem(id)

const update_problem = () => {
    router.push(`/problem/update/${id}`)
}

</script>
<template>
    <ContentBase>
        <h2>
            {{ problem.problem_title }}
        </h2>
        <el-divider />
        <div class="content markdown-body" v-html="renderMarkdown(`${problem.problem_content}`)" />
        <el-divider />
        <RunCode v-if="problem.problem_id != -1" messageDefault=""
        :problemName=problem.problem_title
    :id=id
    :inputAreaMinRow=5
    :inputAreaMaxRow=5
    :run=false
    :defaultCode=false
    :clearCode=false
    :submitCode=true
    :testCode=true
            />
        <el-divider />
        <div v-if="isLoggedIn && user_id == 1" style="text-align: center; margin-top: 10px;">
            <el-button class="button" type="primary" @click="update_problem()">
                编辑
            </el-button>
        </div>
    </ContentBase>
</template>