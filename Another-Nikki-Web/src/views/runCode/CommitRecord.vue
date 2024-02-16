<script setup lang="ts">
import ContentBase from '@/components/ContentBase.vue';
import { ref } from 'vue'
import { ElMessage } from 'element-plus';
import { User } from '@/api';
import { Commits } from '@types/User';
const commits = ref<Commits[]>([])
const get_commits = async () => {
    const ret = await User.commit_records({ })
    if (ret.data.code == 200) {
        commits.value = ret.data.data.commits
    }
    else {
        return ElMessage.error(ret.data.message)
    }
}
get_commits()
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
            <el-table-column prop="problem_name" label="题目" width="100"/>
            <el-table-column prop="compile_status" label="编译结果" width="100"/>
            <el-table-column prop="judge_status" label="运行结果" width="100"/>
            <el-table-column prop="cpu_time_used" label="运行时间" width="100"/>
            <el-table-column prop="memory_used" label="使用内存" width="100"/>
            <el-table-column prop="language" label="使用语言" width="100"/>
            <el-table-column prop="created_time" label="提交时间" width="180"/>
        </el-table>
    </div>
    </ContentBase>
</template>