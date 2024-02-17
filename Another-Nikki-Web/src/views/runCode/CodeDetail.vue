<script setup lang="ts">
import ContentBase from '@/components/ContentBase.vue';
import { useRoute } from 'vue-router';
import { ElMessage } from 'element-plus';
import { ref } from 'vue'
import { User  } from '@/api';
import { Commit } from '@types/User';
import { renderMarkdown } from '@/utils/markdown'
import router from '@/router';
const route = useRoute();
const judge_id: number = parseInt(String(route.params.judge_id));
const commit = ref<Commit>({
})
const get_commit_byId = async (page : number) => {
    const ret = await User.get_commit_by_id({ judge_id: judge_id })
    if (ret.data.code == 200) {
        commit.value = ret.data.data
    }
    else {
        commit.value.code = "请求错误捏, 请不要填奇奇怪怪的参数哦"
        commit.value.problem_name = "空"
        return ElMessage.error(ret.data.message)
    }
}
const toProblemDetail = (id: number) => {
    router.push({ path: `/problem/${id}` })
}
get_commit_byId(judge_id)
</script>

<template>
    <ContentBase>
        <h2 class="custom-hover-column" @click="toProblemDetail(commit.problem_id)">
            {{commit.problem_name}}
        </h2>
        <el-divider />
        <div class="content markdown-body" v-html="renderMarkdown(`${commit.code}`)" />
        <el-divider />
    </ContentBase>
</template>

<style scoped>
@import url('https://cdnjs.cloudflare.com/ajax/libs/KaTeX/0.5.1/katex.min.css');
@import url('https://cdn.jsdelivr.net/github-markdown-css/2.2.1/github-markdown.css');
.custom-hover-column:hover {
  cursor: pointer;
  color: chocolate;
}
</style>