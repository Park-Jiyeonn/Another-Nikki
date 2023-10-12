<script setup lang="ts">
import { ref } from 'vue'

import { LogsType } from '@/types/Logs';
import { page_que, count, get_visit_time } from '@/api/logs'; 
import { getCookies } from "@/hooks/useCookies";

const logs = ref<LogsType[]>([])
const sum = ref(0)
const currentPage = ref(1)
const user_name = getCookies("user_name")
const user_id = getCookies("user_id")
const sum_visit_time = ref(0)
const today_visit_time = ref(0)

const get_page_que = async (page:number) => {
    const ret = await page_que({page:page})
    // console.log(ret.data.data.logs)
    logs.value =  ret.data.data.logs
    logs.value.forEach((log) => {
        log.CreatedAt = log.CreatedAt.substring(5, 10) + " " + log.CreatedAt.substring(11, 16)
    });
}

const get_count = async() => {
    const ret = await count()
    // console.log(ret.data)
    sum.value = ret.data.data.sum
}

const get_visitTime = async() => {
    const ret = await get_visit_time({user_id:user_id})
    sum_visit_time.value = ret.data.data.sum_visit_time
    today_visit_time.value = ret.data.data.today_visit_time
}

get_count()
get_page_que(1)
get_visitTime()

</script>

<template>
    <div> 为了督促某个人学习, 我在 2023.10.12 13:41 为网站加上新的内容： </div>
    <div>
        {{user_name}} 今日访问网站 {{today_visit_time}} 次, 累计访问网站 {{sum_visit_time}} 次
    </div>
    <div style=" margin-top: 10px; ">
        日志:
        <el-table
        :data="logs" 
        stripe style="margin-top: 20px;"
        height="500"  
        >
            <el-table-column prop="ID" label="#" width="70" />
            <el-table-column prop="ip" label="IP" width="150">
                <template #default="{ row }">
                    <a :href="`https://www.ip138.com/iplookup.php?ip=${row.ip}&action=2`">{{ row.ip }}</a>
                </template>
            </el-table-column>
            <el-table-column prop="response" label="客户端" width="120"/>
            <el-table-column prop="CreatedAt" label="时间" width="150"/>
            <el-table-column prop="api" label="链接" width="300"/>
            <el-table-column prop="status" label="状态"/>
        </el-table>
    </div>

    <el-pagination background layout="prev, pager, next" :total="sum" v-model:current-page="currentPage"  :page-size="20" @current-change="get_page_que(currentPage)" />
</template>

<style scoped>

</style>
