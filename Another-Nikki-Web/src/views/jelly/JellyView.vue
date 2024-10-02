<script setup lang="ts">
import { ref } from 'vue'
import HelloWorld from "../../components/HelloWorld.vue"
import { Comment } from '@/types/Comment';
import { CommentApi } from '@/api'
import CommentComponent from '@/components/CommentComponent.vue';
import ContentBase from '@/components/ContentBase.vue';
import { getCookies } from "@/hooks/useCookies";
import { useRouter } from 'vue-router'
import { onMounted, onUnmounted } from 'vue';

const router = useRouter()
const user_id : number = getCookies("user_id")
const get_user = async() => {
    if (user_id != 1 && user_id != 2) {
        router.push(`/auth/login`)
        return
    }
}
get_user()

const random_comments = ref<Comment[]>([])
const laoding_random = ref(false)

const get_random_comment = async () => {
    laoding_random.value = !laoding_random.value
    const ret = await CommentApi.get_random_comment({article_id:0})
    laoding_random.value = false

    random_comments.value = [ret.data.data.comment]
    console.log(random_comments.value)
}

// 控制按钮可见性
const isButtonVisible = ref(true);

const handleScroll = () => {
    const bottomThreshold = 200; // 距离页面底部 200 像素时隐藏按钮
    const scrollPosition = window.innerHeight + window.scrollY;
    const pageHeight = document.body.scrollHeight;

    if (pageHeight - scrollPosition < bottomThreshold) {
        isButtonVisible.value = false;
    } else {
        isButtonVisible.value = true;
    }
};

// 注册滚动事件
onMounted(() => {
    window.addEventListener('scroll', handleScroll);
});

onUnmounted(() => {
    window.removeEventListener('scroll', handleScroll);
});

const scrollToBottom = () => {
    window.scrollTo({
        top: document.body.scrollHeight,
        behavior: 'auto'
    });
};

</script>

<template>
    <ContentBase>
        <div style="text-align: center;">
            <img src="../../assets/vite.svg" class="logo" alt="Vite logo" />
            <img src="../../assets/vue.svg" class="logo vue" alt="Vue logo" />
        </div>

        <HelloWorld msg="Another Nikki" />
        <div v-if="isButtonVisible" class="scroll-to-bottom" @click="scrollToBottom">
            ↓
        </div>
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
.scroll-to-bottom {
    position: fixed;
    bottom: 40px;
    right: 20px;
    background-color: #42b883;
    color: white;
    width: 50px;          /* 设置按钮宽度 */
    height: 50px;         /* 设置按钮高度 */
    border-radius: 50%;   /* 圆形 */
    display: flex;        /* 使内容居中 */
    align-items: center;  /* 垂直居中 */
    justify-content: center; /* 水平居中 */
    cursor: pointer;
    box-shadow: 0 2px 5px rgba(0, 0, 0, 0.2);
    transition: background-color 0.3s ease, transform 0.3s ease;
}

.scroll-to-bottom:hover {
    background-color: #38a169;
}

.scroll-to-bottom:active {
    transform: scale(0.95);
}

</style>
