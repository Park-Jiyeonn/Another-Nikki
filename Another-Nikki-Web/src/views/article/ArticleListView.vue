<script setup lang="ts">
import { ref } from 'vue'
import { Article } from '@/types/Article';
import { ArticleApi } from '@/api';
import router from '@/router';
const articles = ref<Article[]>([])

const article_page_que = async (page: number) => {
    const ret = await ArticleApi.page_que({page:page})
    articles.value = ret.data.data
    articles.value.forEach((article) => {
        article.CreatedAt = article.CreatedAt.substring(0, 10) + " " +
                            article.CreatedAt.substring(11, 16)
    });
}
article_page_que(1)

const toArticleDetail = (id:number) => {
    console.log(id)
    router.push({path: `/article/${id}`})
}

</script>

<template>
    <el-card class="article-card" v-for="article in articles" @click="toArticleDetail(article.ID)">
        <el-row :gutter="24">
            <el-col :span="18">
                <h3>
                    {{ article.title }}
                </h3>
                <p class="article-description">{{article.description}}</p>
                <ul class="article-footer">
                    <li class="article-footer-li">
                        <span class="li-title">{{article.CreatedAt}}</span>
                    </li>
                </ul>
            </el-col>
        </el-row>
    </el-card>
</template>

<style scoped>
.article-card {
  cursor: pointer;
}

.article-description {
    display: -webkit-box;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 2;
    overflow: hidden;
    margin-top: 5px;
    font-size: 14px;
}

.article-footer {
    list-style: none;
    padding: 0px;
    color: var(--el-text-color-secondary);
    font-size: var(--el-font-size-small);
    display: inline-block;
    flex-direction: row;

    .li-title {
        vertical-align: middle;
    }
}
</style>