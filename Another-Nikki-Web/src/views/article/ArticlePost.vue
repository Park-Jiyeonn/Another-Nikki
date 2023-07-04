<script setup lang="ts">
import { ElMessage } from 'element-plus';

import { ref } from 'vue'
import { ArticleApi } from '@/api';

const article_title = ref('')
const article_description = ref('')
const article_content = ref('')
const loading = ref(false);

const post_article = async () => {
    const ret = await ArticleApi.post_article({ 
        title: article_title.value, 
        description: article_description.value,
        content:article_content.value
    })
    // console.log(ret.data)
    if (ret.data.code == 200) {
        ElMessage.success(ret?.data?.message ?? 'success')
    }
    else {
        return ElMessage.error("发布失败！")
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
        <el-button class="button" type="primary" :loading="loading" @click="post_article()">
            发布
        </el-button>
    </div>
</template>