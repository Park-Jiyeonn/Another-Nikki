<script setup lang="ts">
import ContentBase from '@/components/ContentBase.vue';
import { ref } from 'vue'
import { ElMessage } from 'element-plus';
import { User  } from '@/api';
import { Commits } from '@types/User';
const commits = ref<Commits[]>([])
const sum = ref(0)
const currentPage = ref(1)
const get_commits_by_page = async (page : number) => {
    const ret = await User.commit_records({ page_num: page, page_size: 20 })
    if (ret.data.code == 200) {
        commits.value = ret.data.data.commits
    }
    else {
        return ElMessage.error(ret.data.message)
    }
}
const get_count = async() => {
    const ret = await User.count()
    sum.value = ret.data.data.sum
}
get_commits_by_page(1)
get_count()
</script>

<template>
    <ContentBase>
        <div style=" margin-top: 10px; ">
        我的提交:
        <el-table
        :data="commits" 
        stripe style="margin-top: 20px;"
        height="500"  
        >
            <el-table-column prop="judge_id" label="#" width="70" />
            <el-table-column prop="problem_name" label="题目" width="150"/>
            <el-table-column prop="compile_status" label="编译结果" width="100"/>
            <el-table-column prop="judge_status" label="运行结果" width="100"/>
            <el-table-column prop="cpu_time_used" label="运行时间" width="100"/>
            <el-table-column prop="memory_used" label="使用内存" width="100"/>
            <el-table-column prop="language" label="使用语言" width="100"/>
            <el-table-column prop="created_time" label="提交时间" width="180"/>
        </el-table>
    </div>
    <el-pagination background layout="prev, pager, next" :total="sum" v-model:current-page="currentPage"  :page-size="20" @current-change="get_commits_by_page(currentPage)" />
    </ContentBase>
</template>