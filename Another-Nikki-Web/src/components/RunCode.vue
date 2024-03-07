<script setup lang="ts">
const props = defineProps<{
    id: number,
    problemName: string,
    messageDefault: string,
    inputAreaMinRow: number,
    inputAreaMaxRow: number,
    run: boolean,
    defaultCode: boolean,
    clearCode: boolean,
    submitCode: boolean,
    testCode: boolean,
}>()
import { ElMessage } from 'element-plus';

import { ref } from 'vue'

import { CodeRet } from '@/types/runCode';
import { RunCode } from '@/api'
import { getCookies } from "@/hooks/useCookies";
import { Codemirror } from "vue-codemirror";
import { python } from "@codemirror/lang-python";
import { go } from "@codemirror/lang-go";
import { cpp } from "@codemirror/lang-cpp";
import { java } from "@codemirror/lang-java";
import { oneDark } from '@codemirror/theme-one-dark'
import router from '@/router';

const extensions = ref([cpp(), oneDark]);

const textarea = ref('')
const language = ref('c++')
const codeLoading = ref(false)

const user_id = getCookies("user_id")
const user_name = getCookies("username")

const codeMsg = ref<CodeRet>({
    state: "",
    message: "空",
    cpu_time_used: "-",
    memory_used: "-",
    exit_code: "",
})

const options = [
    {
        value: 'Cpp',
        label: 'c++11 (g++ 12.2.1)',
    },
    {
        value: 'Python',
        label: 'python3 (3.11.4)',
    },
    {
        value: 'Java',
        label: 'Java 11',
    },
    {
        value: 'Golang',
        label: 'golang (1.20.5)',
    },
]

const judgeCode = async (code: string) => {
    if (typeof user_id === 'undefined' || user_id == -1) {
        router.push(`/auth/login`)
        return
    }
    codeLoading.value = true
    codeMsg.value = {
        state: "",
        message: "提交中...",
        cpu_time_used: "-",
        memory_used: "-",
        exit_code: "",
    }
    const ret = await RunCode.judgeCode({ user_id: user_id, user_name: user_name, problem_id: props.id, problem_name: props.problemName, language: language.value, code: code })

    codeLoading.value = false
    codeMsg.value.message = "等待编译中..."
    if (ret.data.code == 200) {
        return ElMessage.success(ret?.data?.message ?? 'success')
    }
    else {
        codeMsg.value.message = ret.data.message
        codeMsg.value.state = "编译错误"
        return ElMessage.error("运行失败！")
    }
}

const changeLanguage = (newLang: string) => {
    if (newLang === 'Python') {
        extensions.value = [python(),oneDark];
    }
    else if (newLang === "Java") {
        extensions.value = [java(),oneDark];
    }
    else if (newLang === "Golang") {
        extensions.value = [go(),oneDark];
    }
    else if (newLang === "Cpp") {
        extensions.value = [cpp(),oneDark];
    }
};
</script>

<template>
    选择语言：
    <el-select v-model="language" class="m-2" placeholder="Select" size="large" @change="changeLanguage">
        <el-option v-for="item in options" :key="item.value" :label="item.label" :value="item.value" />
    </el-select>
    {{ messageDefault }}
    <el-divider />
    <el-row>
        <el-col :span="2">
        </el-col>
        <el-col :span="20">
            <Codemirror class="code" v-model="textarea" :style="{ height: '400px' }" :autofocus="true"
                :extensions="extensions" :tabSize=4 style="font-size: 15px;"/>
        </el-col>
    </el-row>

    <div style="margin-top: 10px;">
        <el-row>
            <el-col :span="10"></el-col>
            <el-col :span="10">
                <el-row v-if="submitCode" style="margin-top: 30px;">
                    <el-button class="button" type="primary" :loading="codeLoading" @click="judgeCode(textarea)">
                        提交代码
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
