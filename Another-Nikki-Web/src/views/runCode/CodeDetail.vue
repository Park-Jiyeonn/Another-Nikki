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
get_commit_byId(judge_id)
</script>

<template>
    <ContentBase>
        <h2 class="custom-hover-column" @click="router.push({ path: `/problem/${commit.problem_id}` })">
            {{commit.problem_name}}
        </h2>
        <el-divider />
        <span class="custom-hover-column" @click="router.push({ path: `/profile/${commit.user_id}` })" style="color: #25bb9b"> {{commit.username}}</span> 提交的代码 
        <br/><br/>
        <div style="font-size: 90%;"> 提交时间: {{commit.created_time}} &emsp;  语言: {{commit.language}} &emsp;  运行时间: {{commit.cpu_time_used}} &emsp;  占用内存: {{commit.memory_used}}</div> 
        <div style="font-size: 90%;"> 运行状态: {{commit.judge_status}}</div> 
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