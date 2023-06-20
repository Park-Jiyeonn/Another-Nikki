<script setup lang="ts">
import ContentBase from '@/components/ContentBase.vue';

import axios from 'axios';
import { ref } from 'vue'
import { send_warning, send_success } from "@/components/utils/sendElMsg"

import {
  Delete,
} from '@element-plus/icons-vue'

const textarea = ref('')
const loading = ref(false);
const inputArea = ref("")
const language = ref('c++')

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


const options = [
  {
    value: 'c++',
    label: 'c++11 (g++ 12.2.1)',
  },
  {
    value: 'python',
    label: 'python3 (3.11.4)',
  },
  {
    value: 'java',
    label: 'Java 11',
  },
  {
    value: 'go',
    label: 'golang (1.20.5)',
  },
]


const runCode = (code: string, input: string) => {
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
            "input": input,
            "lang": language.value,
            "code": code,
        }
    })
        .then(function (resp) {
            console.log(resp.data)
            codeMsg.value = resp.data
            if (resp.data.state == 'success') {
                send_success("编译运行成功～")
            } else {
                codeMsg.value.state = "编译错误"
                send_warning("编译失败！")
            }
            loading.value = false
        })
        .catch(function (error) {
            console.log(error)
            send_warning("发生了未知错误，请联系开发者～")
            loading.value = false
        })
}

const codeInit1 = () => {
    axios({
        method: 'get',
        url: TargetPath + '/api/runcode/default_code',
        params:{
            "lang":language.value,
        },
    })
        .then(function (resp) {
            console.log(resp.data)
            textarea.value = resp.data.code
            inputArea.value = resp.data.input
        })
        .catch(function (error) {
            console.log(error)
        })
}

const codeClear = () => {
    textarea.value = ""
    inputArea.value = ""
}

</script>

<template>
    <ContentBase>
        选择语言： 
        <el-select v-model="language" class="m-2" placeholder="Select" size="large">
            <el-option
            v-for="item in options"
            :key="item.value"
            :label="item.label"
            :value="item.value"
            />
        </el-select>
        <el-row>
            <el-col :span="2">
            </el-col>
            <el-col :span="20">
                <el-input v-model="textarea" :autosize="{ minRows: 20, maxRows: 20 }" type="textarea"
                    placeholder="Please input" style="margin-top: 10px;" />
            </el-col>
        </el-row>

        <div style="margin-top: 10px;">
            <el-row>
                <el-col :span="2"></el-col>
                <el-col :span="10">
                    <el-input v-model="inputArea" :autosize="{ minRows: 8, maxRows: 8 }" type="textarea" placeholder="自测输入"
                        style="margin-top: 10px; margin-bottom: 10px;" />
                </el-col>
                <el-col :span="1"></el-col>
                <el-col :span="10">
                    <el-row style="margin-top: 20px;">
                        <el-button class="button" type="success" :loading="loading" @click="runCode(textarea, inputArea)">
                            运行
                        </el-button>
                    </el-row>

                    <el-row style="margin-top: 30px;">
                        <el-button class="button" type="primary" @click="codeInit1()">
                            默认代码
                        </el-button>
                    </el-row>

                    <el-row style="margin-top: 30px;">
                        <el-button class="button" :icon="Delete" type="danger" @click="codeClear()">
                            清空
                        </el-button>
                    </el-row>
                </el-col>
            </el-row>

            <el-form label-width="100px" style="max-width: 460px">
                <el-form-item label="运行结果:" class="limit-text">
                    {{ codeMsg.message }}
                </el-form-item>
                <el-form-item label="运行时间:">
                    {{ codeMsg.cpu_time_used }}
                </el-form-item>
                <el-form-item label="占用内存:">
                    {{ codeMsg.memory_used }}
                </el-form-item>
                <el-form-item label="运行状态:">
                    {{ codeMsg.state }}
                </el-form-item>
            </el-form>

        </div>
    </ContentBase>
</template>

<style scoped>
.limit-text {
    max-height: 100px; /* 设置最大高度 */
    overflow: auto; /* 显示滚动条 */
    white-space: pre-wrap;
}
</style>
