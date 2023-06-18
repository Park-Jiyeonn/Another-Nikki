<script setup lang="ts">
import ContentBase from '../../components/ContentBase.vue';

import axios from 'axios';
import { ref } from 'vue'
import { ElMessage } from "element-plus"

const textarea = ref('')
const loading = ref(false);

const TargetPath = import.meta.env.VITE_API_URL;

interface CodeRet {
    "state": string,
    "message": string,
    "cpu_time_used": string,
    "memory_used": string,
    "exit_code": string,
}
const codeMsg = ref<CodeRet>({
    state: "",
    message: "空",
    cpu_time_used: "-",
    memory_used: "-",
    exit_code: "",
})

const open_warning = (message: string) => {
    ElMessage({
        showClose: true,
        message: message,
        type: 'warning',
    })
}

const open_success = (message: string) => {
    ElMessage({
        showClose: true,
        message: message,
        type: 'success',
    })
}

const runCode = (code: string) => {
    loading.value = !loading.value
    codeMsg.value = {
        state: "",
        message: "docker正在努力工作中...",
        cpu_time_used: "-",
        memory_used: "-",
        exit_code: "",
    }

    axios({
        method: 'post',
        url: TargetPath + '/api/runcode',
        data: {
            "lang": "c++",
            "code": code,
        }
    })
        .then(function (resp) {
            console.log(resp.data)
            codeMsg.value = resp.data
            if (resp.data.state == 'success') {
                open_success("编译运行成功～")
            } else {
                codeMsg.value.state = "编译错误"
                open_warning("编译失败！")
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
            <div class="message-container">{{ codeMsg.message }}</div>


            <div style="margin: 20px" />
            <el-form label-width="100px" style="max-width: 460px">
                <el-form-item label="运行时间: ">
                    {{ codeMsg.cpu_time_used }}
                </el-form-item>
                <el-form-item label="占用内存: ">
                    {{ codeMsg.memory_used }}
                </el-form-item>
                <el-form-item label="运行状态: ">
                    {{ codeMsg.state }}
                </el-form-item>
            </el-form>
        </div>
    </ContentBase>
</template>

<style scoped>
.message-container {
  white-space: pre-line;
}
</style>
