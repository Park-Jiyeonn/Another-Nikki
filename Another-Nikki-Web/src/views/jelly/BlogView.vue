<script setup lang="ts">
import { ref } from 'vue'
import { ElMessage } from 'element-plus';

import HelloWorld from "../../components/HelloWorld.vue"
import { Blog } from '@/types/Blog';
import {BlogApi} from '@/api'

const blogs = ref<Blog[]>([])
const random_blogs = ref<Blog[]>([])
const loading = ref(false);
const laoding_random = ref(false)
const textarea = ref('')

const PostBlog = async (content: string) => {
    loading.value = !loading.value
    textarea.value = ''

    const ret = await BlogApi.PostBlog({content:content})

    loading.value = false
    get_last_seven_blogs()
    if (ret?.data?.code !== 200) return ElMessage.warning(ret?.data?.message ?? 'Something went wrong, please try again.')
    else return ElMessage.success(ret?.data?.message ?? 'success')
}

const get_last_seven_blogs = async () => {
    const ret = await BlogApi.get_last_seven_blogs({num:7})
    blogs.value = ret.data.data
    blogs.value.forEach((blog) => {
        blog.CreatedAt = blog.CreatedAt.substring(5, 10) + " " + blog.CreatedAt.substring(11, 16)
    });
}

const get_random_blogs = async () => {
    laoding_random.value = !laoding_random.value
    const ret = await BlogApi.get_random_blogs()
    laoding_random.value = false

    random_blogs.value = ret.data.data
    random_blogs.value.forEach((blog) => {
        blog.CreatedAt = blog.CreatedAt.substring(5, 10) + " " + blog.CreatedAt.substring(11, 16)
    });
}

get_last_seven_blogs()

</script>

<template>
    <div style="text-align: center;">
      <img src="/vite.svg" class="logo" alt="Vite logo" />
      <img src="../../assets/vue.svg" class="logo vue" alt="Vue logo" />
    </div>
    
    <HelloWorld msg="Another Nikki" />

    <div style=" margin-top: 10px; ">
        留言板：
        <el-table
        :data="blogs" 
        height="400" 
        stripe style="margin-top: 20px;" 
        >
            <el-table-column prop="ID" label="#" width="70" />
            <el-table-column prop="content" label="留言" width="300"/>
            <el-table-column prop="CreatedAt" label="时间" width="150"/>
        </el-table>
    </div>

    <el-row>
        <el-col :span="2">
        </el-col>
        <el-col :span="20">
            <el-input v-model="textarea" :autosize="{ minRows: 2, maxRows: 4 }" type="textarea" placeholder="Please input"
                style="margin-top: 10px;" />
        </el-col>
    </el-row>


    <div style="text-align: center; margin-top: 10px;">
        <el-button class="button" type="primary" :loading="loading" @click="PostBlog(textarea)">
            发布
        </el-button>
    </div>

    <div style=" margin-top: 10px; ">
        随机展示留言板：
        <el-table
        :data="random_blogs" 
        stripe style="margin-top: 20px;" 
        >
            <el-table-column prop="ID" label="#" width="70" />
            <el-table-column prop="content" label="留言" width="300"/>
            <el-table-column prop="CreatedAt" label="时间" width="150"/>
        </el-table>
    </div>

    <div style="text-align: center; margin-top: 10px;">
        <el-button class="button" type="primary" :loading="laoding_random" @click="get_random_blogs">
            点我随机
        </el-button>
    </div>
</template>

<style scoped>
.button {
    margin-top: 10px;
}

.logo {
  height: 4em;
  padding: 1.5em;
  will-change: filter;
  transition: filter 300ms;
}

.logo:hover {
  filter: drop-shadow(0 0 2em #646cffaa);
}

.logo.vue:hover {
  filter: drop-shadow(0 0 2em #42b883aa);
}
</style>
