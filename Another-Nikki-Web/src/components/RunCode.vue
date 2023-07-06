<script setup lang="ts">
const props = defineProps<{
    id?:                number,
    problemName:        string,
    messageDefault:     string,
    inputAreaMinRow:    number,
    inputAreaMaxRow:    number,
    run:                boolean,
    defaultCode:        boolean,
    clearCode:          boolean,
    submitCode:         boolean,
    testCode:           boolean,
}>()
import { ElMessage } from 'element-plus';

import { ref } from 'vue'

import {
    Delete,
} from '@element-plus/icons-vue'

import { CodeRet } from '@/types/runCode';
import { RunCode } from '@/api'

const textarea = ref('')
const loading = ref(false);
const inputArea = ref("")
const language = ref('c++')
const codeLoading = ref(false)


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


const runCode = async (code: string, input: string) => {
    loading.value = !loading.value
    codeMsg.value = {
        state: "",
        message: "docker正在努力工作中...",
        cpu_time_used: "-",
        memory_used: "-",
        exit_code: "",
    }

    const ret = await RunCode.runCode({ input: input, lang: language.value, code: code })

    loading.value = false
    if (ret.data.code == 200) {
        codeMsg.value = ret.data.data
        return ElMessage.success(ret?.data?.message ?? 'success')
    }
    else {
        codeMsg.value.message = ret.data.message
        codeMsg.value.state = "编译错误"
        return ElMessage.error("运行失败！")
    }
}

const codeInit1 = async () => {
    codeLoading.value = true

    const ret = await RunCode.codeInit({ lang: language.value })
    codeLoading.value = false
    console.log(ret.data)
    if (ret.data.code == 200) {
        textarea.value = ret.data.data.code
        inputArea.value = ret.data.data.input
    }
}

const codeClear = () => {
    textarea.value = ""
    inputArea.value = ""
}

const judgeCode = async (code: string, input: string, name: string) => {
    codeLoading.value = true

    codeMsg.value = {
        state: "",
        message: "docker正在努力工作中...",
        cpu_time_used: "-",
        memory_used: "-",
        exit_code: "",
    }
    const problem_name = String(props.id) + "." + name
    const ret = await RunCode.judgeCode({ input: input, lang: language.value, code: code, problem_name: problem_name })

    codeLoading.value = false
    console.log(ret.data)

    if (ret.data.code == 200) {
        codeMsg.value = ret.data.data
        return ElMessage.success(ret?.data?.message ?? 'success')
    }
    else {
        codeMsg.value.message = ret.data.message
        codeMsg.value.state = "编译错误"
        return ElMessage.error("运行失败！")
    }
}

</script>

<template>
    选择语言：
    <el-select v-model="language" class="m-2" placeholder="Select" size="large">
        <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value" />
    </el-select>
    {{ messageDefault }}
    <el-divider />
    <el-row>
        <el-col :span="2">
        </el-col>
        <el-col :span="20">
            <el-input v-model="textarea" :autosize="{ minRows: 20, maxRows: 20 }" type="textarea" placeholder="Please input"
                style="margin-top: 10px;" />
        </el-col>
    </el-row>

    <div style="margin-top: 10px;">
        <el-row>
            <el-col :span="2"></el-col>
            <el-col :span="10">
                <el-input v-model="inputArea" :autosize="{ minRows: inputAreaMinRow, maxRows: inputAreaMaxRow }" type="textarea" placeholder="自测输入"
                    style="margin-top: 10px; margin-bottom: 10px;" />
            </el-col>
            <el-col :span="1"></el-col>
            <el-col :span="10">
                <el-row v-if="run" style="margin-top: 20px;">
                    <el-button class="button" type="success" :loading="loading" @click="runCode(textarea, inputArea)">
                        运行
                    </el-button>
                </el-row>

                <el-row v-if="defaultCode" style="margin-top: 30px;">
                    <el-button class="button" type="primary" :loading="codeLoading" @click="codeInit1()">
                        默认代码
                    </el-button>
                </el-row>

                <el-row v-if="clearCode" style="margin-top: 30px;">
                    <el-button class="button" :icon="Delete" type="danger" @click="codeClear()">
                        清空
                    </el-button>
                </el-row>

                <el-row v-if="submitCode" style="margin-top: 30px;">
                    <el-button class="button" type="primary" :loading="codeLoading"
                        @click="judgeCode(textarea, inputArea, props.problemName)">
                        提交代码
                    </el-button>
                </el-row>

                <el-row v-if="testCode" style="margin-top: 20px;">
                    <el-button class="button" type="success" :loading="loading" @click="runCode(textarea, inputArea)">
                        自测输入
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
</template>

<style scoped>
.limit-text {
    max-height: 100px;
    /* 设置最大高度 */
    overflow: auto;
    /* 显示滚动条 */
    white-space: pre-wrap;
}
</style>