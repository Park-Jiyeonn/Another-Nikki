<script setup lang="ts">
import ContentBase from '@/components/ContentBase.vue';
import { reactive,ref } from 'vue'
import { ElMessage } from 'element-plus';
import { User  } from '@/api';
import { Commits } from '@types/User';
import { getCookies, setCookies } from "@/hooks/useCookies";
const commits = ref<Commits[]>([])
const sum = ref(0)
const currentPage = ref(1)
const dialogFormVisible = ref(false)
const formLabelWidth = '140px'
const user_avatar = ref(getCookies("avatar"))
const username = ref<string>(getCookies("username").toString())
const form = reactive({
  name: username.value,
  avatar_url: user_avatar.value,
})
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
const update_user = async() => {
    const ret = await User.update({username : form.name, avatar: form.avatar_url})
    dialogFormVisible.value = false
    if (ret.data.code == 200) {
        setCookies("username", form.name)
        setCookies("avatar", form.avatar_url)
        user_avatar.value = form.avatar_url
        username.value = form.name
        ElMessage.success("更新成功")
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
}
get_count()
get_commits_by_page(1)
</script>

<template>
    <el-dialog v-model="dialogFormVisible" title="编辑资料" width="500">
        <el-form :model="form">
            <el-form-item label="用户名" :label-width="formLabelWidth">
                <el-input v-model="form.name" autocomplete="off" />
            </el-form-item>
            <el-form-item label="头像链接" :label-width="formLabelWidth">
                <el-input v-model="form.avatar_url" autocomplete="off" />
            </el-form-item>
        </el-form>
        <template #footer>
          <div class="dialog-footer">
            <el-button @click="dialogFormVisible = false">取消</el-button>
            <el-button type="primary" @click="update_user()">
              确认
            </el-button>
          </div>
        </template>
    </el-dialog>

    <ContentBase>
        <el-row :gutter="20">
            <el-col :span="2">
                <el-avatar :size="60" :src="user_avatar" />
            </el-col>
            <el-col :span="5" style="text-align: left;">
                {{username}}
                <div style="margin-top: 10px;"><el-button text @click="dialogFormVisible = true"> 编辑资料 </el-button></div>
            </el-col>
            <el-col :span="13"></el-col>
            <el-col :span="3" style="margin-top: 5px; text-align: center;">
                提交次数
                <div style="margin-top: 7px;">{{sum}}</div>
            </el-col>
        </el-row>
    </ContentBase>
    <ContentBase>
        <el-table :data="commits" height="500" @row-click="handleRowClick">
            <el-table-column prop="judge_id" label="#" width="70" />
            <el-table-column prop="problem_name" label="题目" width="150" class-name="custom-hover-column"/>
            <el-table-column prop="compile_status" label="编译结果" width="100"/>
            <el-table-column prop="judge_status" label="运行结果" width="100"/>
            <el-table-column prop="cpu_time_used" label="运行时间" width="100"/>
            <el-table-column prop="memory_used" label="使用内存" width="100"/>
            <el-table-column prop="language" label="使用语言" width="100"/>
            <el-table-column prop="created_time" label="提交时间" width="180"/>
        </el-table>
        <el-pagination background layout="prev, pager, next" :total="sum" v-model:current-page="currentPage"  :page-size="20" @current-change="get_commits_by_page(currentPage)" />
    </ContentBase>
</template>
<style>
.custom-hover-column:hover {
  cursor: pointer;
  color: chocolate;
}
</style>