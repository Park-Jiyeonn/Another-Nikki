<script setup lang="ts">
import axios from 'axios';
import { ref } from 'vue'

const TargetPath = import.meta.env.VITE_API_URL;
interface Blog {
    ID: number;
    content: string;
    CreatedAt: string,
}
const blogs = ref<Blog[]>([])
const loading = ref(false);
const textarea = ref('')

const PostBlog = (content: string) => {
    console.log(content)
    loading.value = !loading.value
    axios({
        method: 'post',
        url: TargetPath + '/api/create_blog',
        data: {
            "content": content,
        }
    })
        .then(function (resp) {
            console.log(resp.data)
        })
        .catch(function (error) {
            console.log(error)
        })
    loading.value = !loading.value
}


axios({
    method: 'get',
    url: TargetPath + '/api/get_all_blogs',
})
    .then(function (resp) {
        console.log(resp.data.data)
        blogs.value = resp.data.data
        console.log(blogs.value)

        // 遍历 blogs 数组并修改 CreatedAt 值的格式
        blogs.value.forEach((blog) => {
            blog.CreatedAt = blog.CreatedAt.substring(0, 10) + " " + blog.CreatedAt.substring(11, 16)
        });
    })
    .catch(function (error) {
        console.log(error)
    })

</script>

<template>
    <div style="text-align: center; margin-top: 10px;">
        东拉西扯留言板：（目前后端支持发博客了，前端还没有支持哈哈哈）
    </div>
    <!-- <div>
        <textarea></textarea>
        <button>发表</button>
    </div> -->
    <el-row>
        <el-col :span="6">
        </el-col>
        <el-col :span="12">
            <el-table :data="blogs" stripe style="width: 100%; margin-top: 20px;">
                <el-table-column prop="ID" label="#" width="100" />
                <el-table-column prop="content" label="留言" width="300" />
                <el-table-column prop="CreatedAt" label="时间" />
            </el-table>
        </el-col>
    </el-row>

    <el-row>
        <el-col :span="6">
        </el-col>
        <el-col :span="12">
            <el-input v-model="textarea" :autosize="{ minRows: 2, maxRows: 4 }" type="textarea"
                placeholder="Please input" />
        </el-col>
    </el-row>

    <el-row>
        <el-col :span="11">
        </el-col>
        <el-col :span="12">
            <el-button class="button" 
                type="primary" 
                :loading="loading" 
                @click="PostBlog(textarea)">
                发布
            </el-button>
        </el-col>
    </el-row>
</template>

<style scoped>
.button {
    margin-top: 10px;
}
</style>
