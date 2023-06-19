<script setup lang="ts">
import axios from 'axios';
import { ref } from 'vue'
import HelloWorld from "../../components/HelloWorld.vue"
import { send_warning, send_success } from "@/components/utils/sendElMsg"

const TargetPath = import.meta.env.VITE_API_URL;
interface Blog {
    ID: number;
    content: string;
    CreatedAt: string,
}
const blogs = ref<Blog[]>([])
const random_blogs = ref<Blog[]>([])
const loading = ref(false);
const laoding_random = ref(false)
const textarea = ref('')

const PostBlog = (content: string) => {
    console.log(content)
    loading.value = !loading.value
    axios({
        method: 'post',
        url: TargetPath + '/api/blog/create_blog',
        data: {
            "content": content,
        }
    })
        .then(function (resp) {
            console.log(resp.data)
            
            get_last_seven_blogs()
            if (resp.data.message == 'success') {
                send_success("发布成功咯～")
            } else {
                send_warning(resp.data.message)
            }
            
            loading.value = false
        })
        .catch(function (error) {
            console.log(error)
            send_warning("发生了未知错误，请联系开发者～")
            loading.value = false
        })
    textarea.value = ''
}

const get_last_seven_blogs = () => {
    axios({
        method: 'get',
        url: TargetPath + '/api/blog/get_last_seven_blogs',
        params:{
            num:7,
        },
    })
        .then(function (resp) {
            console.log(resp.data.data)
            blogs.value = resp.data.data
            console.log(blogs.value)

            // 遍历 blogs 数组并修改 CreatedAt 值的格式
            blogs.value.forEach((blog) => {
                blog.CreatedAt = blog.CreatedAt.substring(5, 10) + " " + blog.CreatedAt.substring(11, 16)
            });
        })
        .catch(function (error) {
            console.log(error)
        })
}

const get_random_blogs = () => {
    laoding_random.value = !laoding_random.value
    axios({
        method: 'get',
        url: TargetPath + '/api/blog/get_random_blog',
    })
        .then(function (resp) {
            console.log(resp.data.data[0].content)
            random_blogs.value = resp.data.data
            console.log(blogs.value)

            // 遍历 blogs 数组并修改 CreatedAt 值的格式
            random_blogs.value.forEach((blog) => {
                blog.CreatedAt = blog.CreatedAt.substring(5, 10) + " " + blog.CreatedAt.substring(11, 16)
            });

            laoding_random.value = false
        })
        .catch(function (error) {
            console.log(error)

            laoding_random.value = false
        })
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
