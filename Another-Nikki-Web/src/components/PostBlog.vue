<script setup lang="ts">
import axios from 'axios';
import { ref } from 'vue'

interface Blog {
    id: number;
    title: string;
    content: string;
}
const blogs = ref<Blog[]>()

axios({
    method: 'get',
    url: 'http://localhost:8888/api/get_all_blogs',
})
    .then(function (resp) {
        console.log(resp.data.data)
        blogs.value = resp.data.data
        console.log(blogs.value)
    })
    .catch(function (error) {
        console.log(error)
    })

</script>

<template>
    <div>
        东拉西扯留言板：（目前后端支持发博客了，前端还没有支持哈哈哈）
    </div>
    <!-- <div>
        <textarea></textarea>
        <button>发表</button>
    </div> -->
    <table class="table table-hover">
        <thead>
            <tr>
                <th scope="col">#</th>
                <th scope="col">名字</th>
                <th scope="col">博客</th>
            </tr>
        </thead>
        <tbody>
            <tr v-for="blog in blogs" :key="blog.id">
                <th scope="row"> {{ blog.id }} </th>
                <!-- <td>{{ user.username }}</td> -->
                <td> Jiyeon </td>
                <td>{{ blog.content }}</td>
                <!-- <button @click="delete_a_post(blog.id)" type="button" class="btn btn-danger btn-sm">删除</button> -->
            </tr>
        </tbody>
    </table>
</template>

<style scoped></style>
