<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useRoute } from 'vue-router';
import { ElMessage } from 'element-plus';

import { ref } from 'vue'
import { Article } from '@/types/Article';
import { ArticleApi } from '@/api';


const route = useRoute();
const router = useRouter();
const article = ref<Article>({
    article_id: 1,
    article_title: "",
    article_content: "",
    article_description: "",
    created_time: "",
})
const article_title = ref('')
const article_description = ref('')
const article_content = ref('')
const loading = ref(false);

const get_article = async (id: number) => {
    const ret = await ArticleApi.get_article({ article_id: id })
    if (ret.data.code == 200) {
        article.value = ret.data.data
        article_title.value = article.value.article_title
        article_description.value = article.value.article_description
        article_content.value = article.value.article_content
    }
    else {
        return ElMessage.error("获取内容失败！")
    }
    
}

const id: number = parseInt(String(route.params.id));
get_article(id)

const update_article = async () => {
    const ret = await ArticleApi.post_article({ 
        article_id: id, 
        article_title: article_title.value, 
        article_description: article_description.value,
        article_content:article_content.value
    })
    // console.log(ret.data)
    if (ret.data.code == 200) {
        ElMessage.success(ret?.data?.message ?? 'success')
        return router.push(`/article/${id}`)
    }
    else {
        return ElMessage.error("更新失败！")
    }
}


</script>
<template>
    <el-input v-model="article_title" :autosize="{ minRows: 2, maxRows: 2 }" type="textarea" placeholder="Please input"
        style="margin-top: 10px;" />
    <el-input v-model="article_description" :autosize="{ minRows: 2, maxRows: 4 }" type="textarea" placeholder="Please input"
        style="margin-top: 10px;" />
    <el-input v-model="article_content" :autosize="{ minRows: 10, maxRows: 10 }" type="textarea" placeholder="Please input"
        style="margin-top: 10px;" />

    <div style="text-align: center; margin-top: 10px;">
        <el-button class="button" type="primary" :loading="loading" @click="update_article()">
            更新
        </el-button>
    </div>
</template>