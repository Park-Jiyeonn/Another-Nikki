<script setup lang="ts">
const props = defineProps<{
    article_id: number
}>()
import { ref } from 'vue'
import { ElMessage } from 'element-plus';
import type { UploadProps, UploadUserFile } from 'element-plus'

import { Comment } from '@/types/Comment';
import { CommentApi } from '@/api'
import { getCookies } from "@/hooks/useCookies";
import { renderMarkdown } from '@/utils/markdown'

const comments = ref<Comment[]>([])
const loading = ref(false);
const textarea = ref('')

const fileList = ref<UploadUserFile[]>([])

const dialogImageUrl = ref('')

const handleRemove: UploadProps['onRemove'] = (uploadFile, uploadFiles) => {
    console.log(uploadFile, uploadFiles)
}

const handlePictureCardPreview: UploadProps['onPreview'] = (uploadFile) => {
    dialogImageUrl.value = uploadFile.url!
}

const handleUpload: UploadProps['onSuccess'] = (response, uploadFile) => {
    console.log(response)
    console.log(uploadFile)
    uploadFile.url = response.data
}

const post_comment = async (content: string) => {
    fileList.value.forEach((file) => {
        content += `\n\n<img src="${file.url}" width="55%" height="55%">`
    })
    fileList.value = []
    loading.value = !loading.value
    textarea.value = ''

    const user_id = getCookies("user_id")
    const user_name = getCookies("username")
    const user_avatar = getCookies("avatar")

    const ret = await CommentApi.post_comment({
        content: content,
        article_id:props.article_id,
        user_id: user_id,
        username: user_name,
        parent_id: 0,
        root_id: 0,
        parent_name: "",
        user_avatar: user_avatar,
    })

    loading.value = false
    get_last_seven_comments()
    if (ret?.data?.code !== 200) return ElMessage.warning(ret?.data?.message ?? 'Something went wrong, please try again.')
    else return ElMessage.success(ret?.data?.message ?? 'success')
}

const get_last_seven_comments = async () => {
    const ret = await CommentApi.get_last_seven_comments({ num: 7, article_id:props.article_id, })
    comments.value = ret.data.data.comments
}

const reply_comment = async (content: string, parentcomment_id: number, root_id: number, parent_name: string) => {
    const user_id = getCookies("user_id")
    const user_name = getCookies("username")
    const user_avatar = getCookies("avatar")

    const ret = await CommentApi.reply_comment({
        content: content,
        article_id:props.article_id,
        user_id: user_id,
        username: user_name,
        parent_id: parentcomment_id,
        root_id: root_id,
        parent_name: parent_name,
        user_avatar: user_avatar,
    })
    if (ret?.data?.code !== 200) ElMessage.warning(ret?.data?.message ?? 'Something went wrong, please try again.')
    else ElMessage.success(ret?.data?.message ?? 'success')

    loading.value = false
    await get_last_seven_comments()
}

get_last_seven_comments()
</script>
<template>
    <div v-for="item in comments" :key="item.comment_id" class="comment-container">
        <div class="avatar-container">
            <el-avatar :src="item.user_avatar" class="avatar"></el-avatar>
        </div>
        <div class="content-container">
            <div><b style="font-size: 14px">{{ item.username }}</b></div>
            <div class="content" v-html="renderMarkdown(item.content)" />
            <div class="info">
                <span>{{ item.CreatedAt }}</span>
                <el-button link type="primary" style="margin-left: 20px"
                    @click="item.replyIsVisible = !item.replyIsVisible">回复</el-button>
            </div>
            <div v-if="item.children && item.children.length" class="reply-container">
                <div v-for="chl in item.children" :key="chl.comment_id" class="reply">
                    {{ chl.username }} reply {{ chl.parent_name }}: {{ chl.content }}
                    <div class="reply-info">
                        <span>{{ chl.CreatedAt }}</span>
                        <el-button link type="primary" style="margin-left: 20px"
                            @click="chl.replyIsVisible = !chl.replyIsVisible">回复</el-button>
                        <div v-if="chl.replyIsVisible" class="reply-input-container">
                            <el-input v-model="chl.replyText" :autosize="{ minRows: 2, maxRows: 4 }" type="textarea"
                                placeholder="Please input"></el-input>
                            <el-button link type="primary" style="margin-left: 20px"
                                @click="chl.replyIsVisible = false">取消</el-button>
                            <el-button link type="primary" style="margin-left: 20px"
                                @click="chl.replyIsVisible = false; reply_comment(chl.replyText, chl.comment_id, item.comment_id, chl.username)">确定</el-button>
                        </div>
                    </div>
                </div>
            </div>
            <div v-if="item.replyIsVisible" class="reply-input-container">
                <el-input v-model="item.replyText" :autosize="{ minRows: 2, maxRows: 4 }" type="textarea"
                    placeholder="Please input"></el-input>
                <el-button link type="primary" style="margin-left: 20px" @click="item.replyIsVisible = false">取消</el-button>
                <el-button link type="primary" style="margin-left: 20px"
                    @click="item.replyIsVisible = false; reply_comment(item.replyText, item.comment_id, item.comment_id, item.username)">确定</el-button>
            </div>
        </div>
    </div>

    <el-row>
        <el-col :span="2">
        </el-col>
        <el-col :span="20">
            <el-input v-model="textarea" :autosize="{ minRows: 2, maxRows: 4 }" type="textarea" placeholder="Please input"
                style="margin-top: 10px;" />
        </el-col>
    </el-row>

    <div style="display: flex; justify-content: center;  margin-top: 10px;">
        <el-upload v-model:file-list="fileList" class="upload-demo" action="https://jiyeon.club/api/upload"
            :on-preview="handlePictureCardPreview" :on-remove="handleRemove" :on-success="handleUpload" list-type="picture">
            <el-button type="primary">上传图片</el-button>
        </el-upload>
        <el-button class="button" style="margin-left: 10px;" type="primary" :loading="loading"
            @click="post_comment(textarea)">
            发布
        </el-button>
    </div>
</template>

<style scoped>
.comment-container {
    display: flex;
    padding: 20px;
}

.avatar-container {
    text-align: center;
    flex: 1;
}

.avatar {
    width: 60px;
    height: 60px;
    border-radius: 50%;
}

.content-container {
    padding: 0 10px;
    flex: 5;
}

.content {
    padding: 10px 0;
}

.info {
    color: #888;
    font-size: 12px;
    margin-bottom: 10px;
}

.reply-container {
    background-color: #eee;
    padding: 10px;
}

.reply {
    color: #423f3f;
    margin-bottom: 10px;
}

.reply-info {
    color: #888;
    font-size: 12px;
    margin-top: 5px;
    margin-bottom: 10px;
}

.reply-input-container {
    margin-top: 10px;
}
</style>