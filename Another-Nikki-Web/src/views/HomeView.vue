<template>
    <div style="margin-top: 10px; margin-bottom: 80px">
        <el-card>
            <div style="padding: 20px; color: #888">
                <div>
                    <el-input type="textarea" :rows="3" v-model="entity.content"></el-input>
                    <div style="text-align: right; padding: 10px"><el-button type="primary" @click="">留言</el-button>
                    </div>
                </div>
            </div>

            <div style="display: flex; padding: 20px" v-for="item in messages">
                <div style="text-align: center; flex: 1">
                    <el-image :src="item.avatar" style="width: 60px; height: 60px; border-radius: 50%"></el-image>
                </div>
                <div style="padding: 0 10px; flex: 5">
                    <div><b style="font-size: 14px">{{ item.username }}</b></div>
                    <div style="padding: 10px 0; color: #888">
                        {{ item.content }}
                        <el-button type="text" size="mini" @click=""
                            v-if="item.username === user.username">删除</el-button>
                    </div>
                    <div style="background-color: #eee; padding: 10px" v-if="item.parentMessage">{{ item.username }}:{{
                        item.parentMessage }}</div>
                    <div style="color: #888; font-size: 12px">
                        <span>2023-07-04</span>
                        <el-button type="text" style="margin-left: 20px" @click="">回复</el-button>
                    </div>
                </div>
            </div>

            <el-dialog title="回复信息" v-model="dialogFormVisible" width="30%">
                <el-form :model="entity" label-width="80px">
                    <el-form-item label="内容">
                        <el-input v-model="entity.reply" autocomplete="off" type="textarea" :rows="3"></el-input>
                    </el-form-item>
                </el-form>
                <template #footer>
                    <el-button @click="">取 消</el-button>
                    <el-button type="primary" @click="">确 定</el-button>
                </template>
            </el-dialog>
        </el-card>


    </div>
</template>
  
<script setup lang="ts">
import { ref } from 'vue';
// import request from "@/utils/request";

interface Entity {
  parentId?: number;
  content?: string;
  reply?: string;
}

interface User {
  username?: string;
}

interface Message {
  username: string,
  avatar:string,
content:string,
parentMessage:string,

}

const user = ref<User>({
    username:"jiyeon"
});
const messages = ref<Message[]>([
    {
        username: "string",
  avatar:"string",
content:"string",
parentMessage:"string",
    },
]);
const dialogFormVisible = ref(false);
// const isCollapse = ref(false);
const entity = ref<Entity>({
    parentId: 1,
  content: "string",
  reply: "string", 
});

</script>
