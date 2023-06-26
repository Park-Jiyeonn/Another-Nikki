<script setup lang="ts">
import { useRoute } from 'vue-router';
const route = useRoute();

import { ref } from 'vue'
import { Article } from '@/types/Article';
import { ArticleApi } from '@/api';
const article = ref<Article>({
    ID:1,
    title: "title",
    content: "content",
    description:"description",
    CreatedAt: "CreatedAt",
})

const get_article = async (id:number) => {
    const ret = await ArticleApi.get_article({ID:id})
    article.value = ret.data.data
}

const id: string = String(route.params.id);
get_article(parseInt(id))

</script>
<template>
    <h2>
        {{article.title}}
    </h2>
    <div>
        {{ article.content }}
    </div>
</template>