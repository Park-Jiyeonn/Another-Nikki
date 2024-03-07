<script setup lang="ts">
import { ref } from 'vue'
import HelloWorld from "../../components/HelloWorld.vue"
import { Comment } from '@/types/Comment';
import { CommentApi } from '@/api'
import CommentComponent from '@/components/CommentComponent.vue';
import ContentBase from '@/components/ContentBase.vue';

const random_comments = ref<Comment[]>([])
const laoding_random = ref(false)

const get_random_comment = async () => {
    laoding_random.value = !laoding_random.value
    const ret = await CommentApi.get_random_comment()
    laoding_random.value = false

    random_comments.value = [ret.data.data.comment]
    console.log(random_comments.value)
}

</script>

<template>
    <ContentBase>
        <div style="text-align: center;">
            <img src="../../assets/vite.svg" class="logo" alt="Vite logo" />
            <img src="../../assets/vue.svg" class="logo vue" alt="Vue logo" />
        </div>

        <HelloWorld msg="Another Nikki" />

        <CommentComponent :article_id="0"/>

        <div style=" margin-top: 10px; ">
            随机展示留言板：
            <el-table :data="random_comments" stripe style="margin-top: 20px;">
                <el-table-column prop="comment_id" label="#" width="70" />
                <el-table-column prop="content" label="留言" width="300" />
                <el-table-column prop="created_time" label="时间" width="300" />
            </el-table>
        </div>

        <div style="text-align: center; margin-top: 10px;">
            <el-button class="button" type="primary" :loading="laoding_random" @click="get_random_comment">
                点我随机
            </el-button>
        </div>
    </ContentBase>
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
