<script setup lang="ts">
import { useRoute } from 'vue-router';
import { useRouter } from 'vue-router'
import { ref } from 'vue'

import { Article } from '@/types/Article';
import { ArticleApi } from '@/api';
import { renderMarkdown } from '@/utils/markdown'
import { useIsLoggedIn } from '@/hooks/userIsLogin'
import CommentComponent from '@/components/CommentComponent.vue';
import ContentBase from '@/components/ContentBase.vue';
import { getCookies } from "@/hooks/useCookies";

const isLoggedIn = useIsLoggedIn()
const route = useRoute();
const router = useRouter();
const user_id : number = getCookies("user_id")

const article = ref<Article>({
    article_id: 1,
    article_title: "title",
    article_content: "content",
    article_description: "description",
    created_time: "CreatedAt",
})

const get_article = async (id: number) => {
    const ret = await ArticleApi.get_article({ article_id: id })
    if (ret.data.code < 200 || ret.data.code > 200) {
        article.value.article_id = -1
        article.value.article_title = "请求错误捏, 请不要填奇奇怪怪的参数哦"
        article.value.article_description = "空"
        article.value.article_content = "空"
        return
    }
    article.value = ret.data.data
}

const article_id: number = parseInt(String(route.params.id));
get_article(article_id)

const update_article = () => {
    router.push(`/article/update/${article_id}`)
}

</script>
<template>
    <ContentBase>
        <h2>
            {{ article.article_title }}
        </h2>
        <el-divider />
        <div>
            {{ article.article_description }}
        </div>
        <el-divider />
        <div class="content markdown-body" v-html="renderMarkdown(article.article_content)" />

        <div v-if="isLoggedIn && user_id == 1" style="text-align: center; margin-top: 10px;">
            <el-button class="button" type="primary" @click="update_article()">
                编辑
            </el-button>
        </div>

        <el-divider />
        <div>
            评论区：
        </div>

        <CommentComponent :article_id=article_id />
    </ContentBase>
</template>

<style scoped>
</style>