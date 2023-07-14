<script setup lang="ts">
import { ref } from 'vue'
import HelloWorld from "../../components/HelloWorld.vue"
import { Comment } from '@/types/Comment';
import { CommentApi } from '@/api'
import CommentComponent from '@/components/CommentComponent.vue';

const random_comments = ref<Comment[]>([])
const laoding_random = ref(false)

const get_random_comment = async () => {
    laoding_random.value = !laoding_random.value
    const ret = await CommentApi.get_random_comment()
    laoding_random.value = false

    random_comments.value = ret.data.data
    random_comments.value.forEach((blog) => {
        blog.CreatedAt = blog.CreatedAt.substring(5, 10) + " " + blog.CreatedAt.substring(11, 16)
    });
}

</script>

<template>
    <div style="text-align: center;">
        <img src="../../assets/vite.svg" class="logo" alt="Vite logo" />
        <img src="../../assets/vue.svg" class="logo vue" alt="Vue logo" />
    </div>

    <HelloWorld msg="Another Nikki" />

    <CommentComponent/>

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
