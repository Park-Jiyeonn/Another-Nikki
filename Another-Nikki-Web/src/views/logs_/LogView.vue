<script setup lang="ts">
import { ref } from 'vue'

import { LogsType } from '@/types/Logs';
import { page_que} from '@/api/logs'; 

const logs = ref<LogsType[]>([])
const sum = ref(0)
const currentPage = ref(1)

const get_page_que = async (page_num:number, page_size: number) => {
    const ret = await page_que({page_num:page_num, page_size: page_size})
    // console.log(ret.data.data.logs)
    logs.value =  ret.data.data.logs
    sum.value = ret.data.data.sum_log
}

get_page_que(1, 20)

</script>

<template>
    <div style=" margin-top: 10px; ">
        日志:
        <el-table
        :data="logs" 
        stripe style="margin-top: 20px;"
        height="500"  
        >
            <el-table-column prop="log_id" label="#" width="70" />
            <el-table-column prop="ip" label="IP" width="136">
                <template #default="{ row }">
                    <a :href="`https://www.ip138.com/iplookup.php?ip=${row.ip}&action=2`">{{ row.ip }}</a>
                </template>
            </el-table-column>
            <el-table-column prop="platform" label="客户端" width="100"/>
            <el-table-column prop="ts" label="时间" width="170"/>
            <el-table-column prop="url" label="链接" width="200"/>
            <el-table-column prop="level" label="状态" width="80"/>
            <el-table-column prop="code" label="code" width="60"/>
            <el-table-column prop="msg" label="msg" width="250"/>
            <el-table-column prop="stack" label="stack" width="250"/>
        </el-table>
    </div>

    <el-pagination background layout="prev, pager, next" :total="sum" v-model:current-page="currentPage"  :page-size="20" @current-change="get_page_que(currentPage, 20)" />
</template>

<style scoped>

</style>
