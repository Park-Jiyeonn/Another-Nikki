<script setup lang="ts">
import { useRoute } from 'vue-router';
import { useRouter } from 'vue-router'
import { ref } from 'vue'

import { Article } from '@/types/Article';
import { ArticleApi } from '@/api';
// import 'highlight.js/styles/github-dark-dimmed.css'
import { renderMarkdown } from '@/utils/markdown'
import { useIsLoggedIn } from '@/hooks/userIsLogin'
import CommentComponent from '@/components/CommentComponent.vue';
import ContentBase from '@/components/ContentBase.vue';

const isLoggedIn = useIsLoggedIn()
const route = useRoute();
const router = useRouter();

const article = ref<Article>({
    ID: 1,
    title: "title",
    content: "content",
    description: "description",
    CreatedAt: "CreatedAt",
})

const get_article = async (id: number) => {
    const ret = await ArticleApi.get_article({ ID: id })
    article.value = ret.data.data
}

const id: number = parseInt(String(route.params.id));
get_article(id)

const update_article = () => {
    router.push(`/article/update/${id}`)
}

</script>
<template>
    <ContentBase>
        <h2>
            {{ article.title }}
        </h2>
        <el-divider />
        <div>
            {{ article.description }}
        </div>
        <el-divider />
        <div class="content markdown-body" v-html="renderMarkdown(article.content)" />

        <div v-if="isLoggedIn" style="text-align: center; margin-top: 10px;">
            <el-button class="button" type="primary" @click="update_article()">
                编辑
            </el-button>
        </div>

        <el-divider />
        <div>
            评论区：
        </div>

        <CommentComponent :article_id=id />
    </ContentBase>
</template>

<style scoped>
</style>