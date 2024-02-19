<script setup lang="ts">
import { ref } from 'vue'
import { Article } from '@/types/Article';
import { ArticleApi } from '@/api';
import router from '@/router';
const articles = ref<Article[]>([])

const article_page_que = async (page_num: number, page_size: number) => {
    const ret = await ArticleApi.page_que({page_num:page_num, page_size: page_size})
    articles.value = ret.data.data.articles
}
article_page_que(1, 10)

const toArticleDetail = (id:number) => {
    console.log(id)
    router.push({path: `/article/${id}`})
}

</script>

<template>
    <el-card class="article-card" v-for="article in articles" @click="toArticleDetail(article.article_id)">
        <el-row :gutter="24">
            <el-col :span="18">
                <h3>
                    {{ article.article_title }}
                </h3>
                <p class="article-description">{{article.article_description}}</p>
                <ul class="article-footer">
                    <li class="article-footer-li">
                        <span class="li-title">{{article.created_time}}</span>
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