<script setup lang="ts">
import ContentBase from '../../components/ContentBase.vue';

import axios from 'axios';
import { ref } from 'vue'
import { ElMessage } from "element-plus"

const textarea = ref('')
const loading = ref(false);

const TargetPath = import.meta.env.VITE_API_URL;
const compileMsg = ref("空")

const open_warning = (message:string) => {
  ElMessage({
    showClose: true,
    message: message,
    type: 'warning',
  })
}

const open_success = (message:string) => {
  ElMessage({
    showClose: true,
    message: message,
    type: 'success',
  })
}

const runCode = (code: string) => {
    loading.value = !loading.value
    compileMsg.value = "docker正在努力工作中..."
    axios({
        method: 'post',
        url: TargetPath + '/api/runcode',
        data: {
            "lang":"c++",
            "code": code,
        }
    })
        .then(function (resp) {
            console.log(resp.data)
            compileMsg.value = resp.data.message
            if (resp.data.state == 'success') {
                open_success("编译运行成功～")
            } else {
                open_warning(resp.data.message)
            }
            loading.value = false
        })
        .catch(function (error) {
            console.log(error)
            open_warning("发生了未知错误，请联系开发者～")
            loading.value = false
        })
}
</script>

<template>
    <ContentBase>
        这里可以编译代码：（目前只支持c++）
        <el-row>
            <el-col :span="2">
            </el-col>
            <el-col :span="20">
                <el-input v-model="textarea" :autosize="{ minRows: 20, maxRows: 50 }" type="textarea"
                    placeholder="Please input" style="margin-top: 10px;" />
            </el-col>
        </el-row>

        <div style="margin-top: 10px;">
            <el-button class="button" type="primary" :loading="loading" @click="runCode(textarea)">
                运行
            </el-button>
            输出：{{ compileMsg }}
        </div>
    </ContentBase>
</template>

<style scoped></style>
