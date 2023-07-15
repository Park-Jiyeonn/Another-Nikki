<script setup lang="ts">
import { ref } from 'vue'

import { LogsType } from '@/types/Logs';
import { page_que, count } from '@/api/logs'; 

const logs = ref<LogsType[]>([])
const sum = ref(0)
const currentPage = ref(1)

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

get_count()
get_page_que(1)

</script>

<template>
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
