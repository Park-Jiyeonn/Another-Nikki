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

const isLoggedIn = useIsLoggedIn()
const route = useRoute();
const router = useRouter();

const problem = ref<Problem>({
    ID: 1,
    name: "name",
    content: "content",
    CreatedAt: "CreatedAt",
})

const get_problem = async (id: number) => {
    const ret = await ProblemApi.get_problem({ ID: id })
    problem.value = ret.data.data
}

const id: number = parseInt(String(route.params.id));
get_problem(id)

const update_problem = () => {
    router.push(`/problem/update/${id}`)
}


import ContentBase from '@/components/ContentBase.vue';

</script>
<template>
    <ContentBase>
        <h2>
            {{ problem.name }}
        </h2>

        <el-divider />

        <div class="content markdown-body" v-html="renderMarkdown(problem.content)" />

        <el-divider />

        <RunCode messageDefault=""
        :problemName=problem.name
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

        <div v-if="isLoggedIn" style="text-align: center; margin-top: 10px;">
            <el-button class="button" type="primary" @click="update_problem()">
                编辑
            </el-button>
        </div>
    </ContentBase>
</template>

<style scoped>
@import url('https://cdnjs.cloudflare.com/ajax/libs/KaTeX/0.5.1/katex.min.css');
@import url('https://cdn.jsdelivr.net/github-markdown-css/2.2.1/github-markdown.css');
</style>