<script setup lang="ts">
import { ref } from 'vue'
import { ElMessage } from 'element-plus';

import HelloWorld from "../../components/HelloWorld.vue"
import { Comment } from '@/types/Comment';
import { CommentApi } from '@/api'
import { getCookies } from "@/hooks/useCookies";

const comments = ref<Comment[]>([])
const random_comments = ref<Comment[]>([])
const loading = ref(false);
const laoding_random = ref(false)
const textarea = ref('')


const post_comment = async (content: string) => {
    loading.value = !loading.value
    textarea.value = ''

    const user_id = getCookies("user_id")
    const user_name = getCookies("user_name")
    const user_avatar = getCookies("user_avatar")

    const ret = await CommentApi.post_comment({
        content:            content,
        author_id:          user_id,
        author_name:        user_name,
        parent_id:          0,
        root_id:            0,
        parent_name:        "",
        author_avatar:      user_avatar,
    })

    loading.value = false
    get_last_seven_comments()
    if (ret?.data?.code !== 200) return ElMessage.warning(ret?.data?.message ?? 'Something went wrong, please try again.')
    else return ElMessage.success(ret?.data?.message ?? 'success')
}

const reply_comment = async (content: string, parentID: number, root_id: number, parent_name:string) => {
    const user_id = getCookies("user_id")
    const user_name = getCookies("user_name")
    const user_avatar = getCookies("user_avatar") 

    const ret = await CommentApi.reply_comment({
        content:                content, 
        author_id:              user_id,
        author_name:            user_name,
        parent_id:              parentID,
        root_id:                root_id,
        parent_name:            parent_name,
        author_avatar:          user_avatar,
    })
    if (ret?.data?.code !== 200) ElMessage.warning(ret?.data?.message ?? 'Something went wrong, please try again.')
    else ElMessage.success(ret?.data?.message ?? 'success') 

    loading.value = false
    await get_last_seven_comments()
}

const get_last_seven_comments = async () => {
    const ret = await CommentApi.get_last_seven_comments({ num: 7 })
    comments.value = ret.data.data
    comments.value.forEach((blog) => {
        blog.CreatedAt = blog.CreatedAt.substring(5, 10) + " " + blog.CreatedAt.substring(11, 16)
        blog.children.forEach((chl) => {
            chl.CreatedAt = chl.CreatedAt.substring(5, 10) + " " + chl.CreatedAt.substring(11, 16)
        })
    });
}

const get_random_comment = async () => {
    laoding_random.value = !laoding_random.value
    const ret = await CommentApi.get_random_comment()
    laoding_random.value = false

    random_comments.value = ret.data.data
    random_comments.value.forEach((blog) => {
        blog.CreatedAt = blog.CreatedAt.substring(5, 10) + " " + blog.CreatedAt.substring(11, 16)
    });
}

get_last_seven_comments()

</script>

<template>
    <div style="text-align: center;">
        <img src="/vite.svg" class="logo" alt="Vite logo" />
        <img src="../../assets/vue.svg" class="logo vue" alt="Vue logo" />
    </div>

    <HelloWorld msg="Another Nikki" />

    <div style="display: flex; padding: 20px" v-for="item in comments">
        <div style="text-align: center; flex: 1">
            <el-image :src="item.author_avatar"
                style="width: 60px; height: 60px; border-radius: 50%"></el-image>
        </div>
        <div style="padding: 0 10px; flex: 5">
            <div><b style="font-size: 14px">{{ item.author_name }}</b></div>
            <div style="padding: 10px 0; color: #888">
                {{ item.content }}
            </div>
            <div style="color: #888; font-size: 12px; margin-bottom: 10px;">
                <span>{{ item.CreatedAt }}</span>
                <el-button link type="primary" style="margin-left: 20px" @click="item.replyIsVisible = !item.replyIsVisible">回复</el-button>
            </div>
            <div style="background-color: #eee; padding: 10px" v-if="item.children.length"> 
                <div v-for="chl in item.children">
                    {{ chl.author_name }} 
                    reply
                    {{ chl.parent_name }}
                    :
                    {{ chl.content }}
                    <div style="color: #888; font-size: 12px; margin-top: 5px; margin-bottom: 10px;">
                        <span>{{ chl.CreatedAt }}</span>
                        <el-button link type="primary" style="margin-left: 20px" @click="chl.replyIsVisible = !chl.replyIsVisible">回复</el-button>
                        
                        <div v-if="chl.replyIsVisible">
                            <el-input   v-model="chl.replyText"
                                        :autosize="{ minRows: 2, maxRows: 4 }" 
                                        type="textarea" 
                                        placeholder="Please input"
                                        style="margin-top: 10px;" />
                            <el-button link type="primary" style="margin-left: 20px" @click="chl.replyIsVisible = false">取消</el-button> 
                            <el-button link type="primary" style="margin-left: 20px" @click="chl.replyIsVisible = false; reply_comment(chl.replyText, chl.ID, item.ID, chl.author_name)">确定</el-button>
                        </div>
                    </div>
                </div>
            </div>
            
            <div v-if="item.replyIsVisible">
                <el-input   v-model="item.replyText"
                            :autosize="{ minRows: 2, maxRows: 4 }" 
                            type="textarea" 
                            placeholder="Please input"
                            style="margin-top: 10px;" />
               <el-button link type="primary" style="margin-left: 20px" @click="item.replyIsVisible = false">取消</el-button> 
               <el-button link type="primary" style="margin-left: 20px" @click="item.replyIsVisible = false; reply_comment(item.replyText, item.ID, item.ID, item.author_name)">确定</el-button>
            </div>
        </div>
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
        <el-button class="button" type="primary" :loading="loading" @click="post_comment(textarea)">
            发布
        </el-button>
    </div>

    <div style=" margin-top: 10px; ">
        随机展示留言板：
        <el-table :data="random_comments" stripe style="margin-top: 20px;">
            <el-table-column prop="ID" label="#" width="70" />
            <el-table-column prop="content" label="留言" width="300" />
            <el-table-column prop="CreatedAt" label="时间" width="150" />
        </el-table>
    </div>

    <div style="text-align: center; margin-top: 10px;">
        <el-button class="button" type="primary" :loading="laoding_random" @click="get_random_comment">
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
