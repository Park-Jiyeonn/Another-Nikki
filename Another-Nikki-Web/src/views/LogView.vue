<script setup lang="ts">
import axios from 'axios';
import { ref } from 'vue'

const TargetPath = import.meta.env.VITE_API_URL;
interface Log {
    ID: number
    IP: number
    api: string
    status: number
    response:string
    CreatedAt:string
}
const logs = ref<Log[]>([])
const sum = ref(0)
const currentPage = ref(2)

const get_page_que = (page:number) => {
    axios({
        method: 'get',
        url: TargetPath + '/api/log/get_page_que',
        params:{
            page:page,
        },
    })
        .then(function (resp) {
            logs.value = resp.data.data 
            logs.value.forEach((log) => {
                log.CreatedAt = log.CreatedAt.substring(5, 10) + " " + log.CreatedAt.substring(11, 16)
            });
        })
        .catch(function (error) {
            console.log(error)
        })
}

const get_count = () => {
    axios({
        method:"get",
        url: TargetPath + '/api/log/count',
    })
        .then(function (resp) {
            sum.value = resp.data.sum
            // sum.value = ((sum.value + 20 - 1) / 20 | 0) * 20
            console.log(sum.value)
        })
        .catch(function (error) {
            console.log(error)
        })
}

get_count()
get_page_que(1)

</script>

<template>
    <div style=" margin-top: 10px; ">
        日志：
        <el-table
        :data="logs" 
        stripe style="margin-top: 20px;"
        height="500"  
        >
            <el-table-column prop="ID" label="#" width="70" />
            <el-table-column prop="ip" label="IP"/>
            <el-table-column prop="api" label="链接" width="300"/>
            <el-table-column prop="status" label="状态"/>
            <el-table-column prop="response" label="结果" width="300"/>
            <el-table-column prop="CreatedAt" label="时间" width="150"/>
        </el-table>
    </div>

    <el-pagination background layout="prev, pager, next" :total="sum" v-model:current-page="currentPage"  :page-size="20" @current-change="get_page_que(currentPage)" />
</template>

<style scoped>

</style>
