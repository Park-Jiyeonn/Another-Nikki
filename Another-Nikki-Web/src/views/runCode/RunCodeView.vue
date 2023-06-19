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
    label: 'c++',
  },
  {
    value: 'python',
    label: 'python',
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
    language.value = "c++"
    textarea.value = `#include <stdio.h>
#include <stdlib.h>
#include <time.h>

int roll_dice() {
    return rand() % 6 + 1;
}

int main() {
    int dice1, dice2, sum;
    srand(time(0));
    dice1 = roll_dice();
    dice2 = roll_dice();
    sum = dice1 + dice2;
    printf("你掷的两个骰子点数分别为 %d 和 %d，点数之和为 %d\\n", dice1, dice2, sum);
    return 0;
}`
    inputArea.value = ""
}

const codeInit2 = () => {
    language.value = "c++"
    textarea.value = `#include <bits/stdc++.h>
using namespace std;
#define int long long
const int N = 1e6 + 5, mod = 998244353;
int a[N];
void solve()
{
    int n;
    cin >> n;
    for (int i = 1; i <= n; i++) {
        cin >> a[i];
    }
    for (int i = 1; i <= n; i++) {
        cout << a[i] * 10 << " ";
    }
    cout << "\\n";
}
signed main()
{
    ios::sync_with_stdio(0);
    cin.tie(0);
    int tt = 1;
    // cin >> tt;
    while (tt--) solve();
    return 0;
}`
    inputArea.value = `5
5 4 3 2 1`
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
                    <el-row style="margin-top: 10px;">
                        <el-button class="button" type="success" :loading="loading" @click="runCode(textarea, inputArea)">
                            运行
                        </el-button>
                    </el-row>

                    <el-row style="margin-top: 10px;">
                        <el-button class="button" type="primary" @click="codeInit1()">
                            默认代码 -- 骰子
                        </el-button>
                    </el-row>

                    <el-row style="margin-top: 10px;">
                        <el-button class="button" type="primary" @click="codeInit2()">
                            默认代码 -- 自测输入
                        </el-button>
                    </el-row>

                    <el-row style="margin-top: 10px;">
                        <el-button class="button" :icon="Delete" type="danger" @click="codeClear()">
                            清空
                        </el-button>
                    </el-row>
                </el-col>
            </el-row>


            <el-form label-width="100px" style="max-width: 460px">
                <el-form-item label="运行结果: " class="limit-text">
                    {{ codeMsg.message }}
                </el-form-item>
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
.limit-text {
    max-height: 100px; /* 设置最大高度 */
    overflow: auto; /* 显示滚动条 */
}
</style>
