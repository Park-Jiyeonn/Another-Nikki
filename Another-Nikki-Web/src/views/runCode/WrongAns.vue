<script setup lang="ts">
import ContentBase from '@/components/ContentBase.vue';
import { ref } from 'vue'
import { ElMessage } from 'element-plus';
import { User  } from '@/api';
import { Commits } from '@types/User';
import { useRoute } from 'vue-router'
import router from '@/router';
const route = useRoute();
const user_id: number = parseInt(String(route.params.user_id));
const commits = ref<Commits[]>([])
const get_user_wrong_answer = async (user_id : number) => {
    const ret = await User.get_user_wrong_answer({user_id : user_id,})
    if (ret.data.code == 200) {
        commits.value = ret.data.data.user_wrong_problem
    }
    else {
        return ElMessage.error(ret.data.message)
    }
}
const toProblemDetail = (id: number) => {
    router.push({ path: `/problem/${id}` })
}
get_user_wrong_answer(user_id)
</script>

<template>
    <ContentBase>
        <div class="container">
        <el-card class="problem-card" v-for="problem in commits" @click="toProblemDetail(problem.problem_id)">
            <span> {{ problem.problem_name }} </span>
        </el-card>
        </div>
    </ContentBase>
</template>
<style scoped>
.container {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-wrap: wrap;
}

.problem-card {
    cursor: pointer;
    width: 960px;
}
.problem-card:hover {
    color: chocolate;
}
</style>