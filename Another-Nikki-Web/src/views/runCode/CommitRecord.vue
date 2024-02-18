<script setup lang="ts">
import ContentBase from '@/components/ContentBase.vue';
import { ref } from 'vue'
import { ElMessage } from 'element-plus';
import { User  } from '@/api';
import { Commits } from '@types/User';
import { getCookies, setCookies } from "@/hooks/useCookies";
import { useRoute } from 'vue-router'
const route = useRoute();
const user_id: number = parseInt(String(route.params.user_id));
const commits = ref<Commits[]>([])
const sum = ref(0)
const currentPage = ref(1)
const user_avatar = ref("")
const username = ref("")
const user_description = ref("")
const get_commits_by_page = async (page : number, user_id : number) => {
    const ret = await User.commit_records({user_id : user_id, page_num: page, page_size: 10 })
    if (ret.data.code == 200) {
        commits.value = ret.data.data.commits
    }
    else {
        return ElMessage.error(ret.data.message)
    }
}
const get_count = async(user_id : number) => {
    const ret = await User.count({user_id:user_id})
    sum.value = ret.data.data.sum
}
const get_user = async(user_id : number) => {
    const ret = await User.get_user_by_id({user_id : user_id})
    if (ret.data.code == 200) {
        username.value = ret.data.data.username
        user_avatar.value = ret.data.data.avatar
        user_description.value = ret.data.data.description
    }
    else {
        return ElMessage.error(ret.data.message)
    } 
}
const handleRowClick = async(row: Commits, column: any, event: Event) => {
    if (column.property === 'problem_name') {
      const problemId: number = row.problem_id;
      window.open(`${import.meta.env.VITE_HTTP_URL}/problem/${problemId}`, '_blank');
    }
    if (column.property === 'judge_id' || column.property === 'judge_status' || column.property === 'compile_status') {
      const judgeId: number = row.judge_id;
      window.open(`${import.meta.env.VITE_HTTP_URL}/user/code/${judgeId}`, '_blank');
    }
}
function formatJudgeStatus({ row, column, rowIndex, columnIndex }: { row: Commits, column: any, rowIndex: number, columnIndex: number }){
    if (column.property === 'compile_status') {
        if (row.compile_status === 'Compile Failed') {
            return { cursor:'pointer',color: 'red',}
        } 
        if (row.compile_status !== 'Compile Success') {
            return { cursor:'pointer',color: '#409eff',}
        }
        return { cursor:'pointer',}
    }
    if (column.property === 'judge_status') {
        if (row.judge_status === 'Wrong Answer') {
            return { cursor:'pointer',color: 'red',}
        }
        if (row.judge_status === 'Accept') {
            return { cursor:'pointer',color: '#25bb9b',}
        }
        return { cursor:'pointer',}
    }
}
get_user(user_id)
get_count(user_id)
get_commits_by_page(1, user_id)
</script>

<template>
    <ContentBase>
        <el-row :gutter="20">
            <el-col :span="2">
                <el-avatar :size="60" :src="user_avatar" />
            </el-col>
            <el-col :span="5">
                <div style="text-align: left; color:#25bb9b;">{{username}}</div>
                <br/>
                <div>{{user_description}}</div>
            </el-col>
            <el-col :span="13"></el-col>
            <el-col :span="3" style="margin-top: 5px; text-align: center;">
                提交次数
                <div style="margin-top: 7px;">{{sum}}</div>
            </el-col>
        </el-row>
    </ContentBase>
    <ContentBase>
        <el-table :data="commits" :cell-style="formatJudgeStatus" height="500" @row-click="handleRowClick">
            <el-table-column prop="judge_id" label="#" width="70" class-name="custom-hover-column"/>
            <el-table-column prop="problem_name" label="题目" width="150" class-name="custom-hover-column"/>
            <el-table-column prop="compile_status" label="编译结果" width="100"/>
            <el-table-column prop="judge_status" label="运行结果" width="100"/>
            <el-table-column prop="cpu_time_used" label="运行时间" width="100"/>
            <el-table-column prop="memory_used" label="使用内存" width="100"/>
            <el-table-column prop="language" label="使用语言" width="100"/>
            <el-table-column prop="created_time" label="提交时间" width="180"/>
        </el-table>
        <el-pagination background layout="prev, pager, next" :total="sum" v-model:current-page="currentPage"  :page-size="10" @current-change="get_commits_by_page(currentPage, user_id)" />
    </ContentBase>
</template>
<style>
.custom-hover-column:hover {
  cursor: pointer;
  color: chocolate;
}
</style>