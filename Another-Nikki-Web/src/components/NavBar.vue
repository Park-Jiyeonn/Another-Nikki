<template>
    <el-menu mode="horizontal" :ellipsis="false">
        <el-menu-item v-if="(user_id == 1 || user_id == 2)" index="1" @click="router.push('/jelly')">
            Jelly
        </el-menu-item>
        <el-menu-item index="5" @click="router.push('/message-board')">
            留言板
        </el-menu-item>
        <el-menu-item v-if="(user_id == 1 || user_id == 2)" index="6" @click="router.push('/logs')">
            日志
        </el-menu-item>
        <el-menu-item index="2" @click="router.push('/problems')">
            题目集
        </el-menu-item>
        <el-menu-item index="3" @click="router.push('/articles')">
            文章与题解
        </el-menu-item>
        <div class="flex-grow" />
        <el-menu-item index="4" @click="get_user()">
            <el-dropdown trigger="click">
                <el-avatar :src="user_avatar"/>
                <template #dropdown> <el-dropdown-menu>
                    <el-dropdown-item @click="dialogFormVisible=true"> 修改资料 </el-dropdown-item>
                    <el-dropdown-item @click="logout()"> 退出登录 </el-dropdown-item>
                    <el-dropdown-item @click="wrong_ans()"> 我的错题 </el-dropdown-item>
                  </el-dropdown-menu> </template>
            </el-dropdown>
        </el-menu-item>
    </el-menu>

    <el-dialog v-model="dialogFormVisible" title="编辑资料" width="500">
        <el-form :model="form">
            <el-form-item label="用户名" :label-width="formLabelWidth">
                <el-input v-model="form.name" autocomplete="off" />
            </el-form-item>
            <el-form-item label="头像链接" :label-width="formLabelWidth">
                <el-input v-model="form.avatar_url" autocomplete="off" />
            </el-form-item>
            <el-form-item label="个性签名" :label-width="formLabelWidth">
                <el-input v-model="form.description" autocomplete="off" />
            </el-form-item>
            <el-form-item label="旧密码" :label-width="formLabelWidth">
                <el-input v-model="form.old_password" autocomplete="off" show-password />
            </el-form-item>
            <el-form-item label="新密码" :label-width="formLabelWidth">
                <el-input v-model="form.new_password" autocomplete="off" show-password />
            </el-form-item>
        </el-form>
        <template #footer>
          <div class="dialog-footer">
            <el-button @click="dialogFormVisible = false">取消</el-button>
            <el-button type="primary" @click="update_user()">
              确认
            </el-button>
          </div>
        </template>
    </el-dialog>
</template>

<script lang="ts" setup>
import { reactive,ref } from 'vue'
import { useRouter } from 'vue-router'
// import { useIsLoggedIn } from '@/hooks/userIsLogin'
import { getCookies, setCookies } from "@/hooks/useCookies";
import { ElMessage } from 'element-plus';
import { User  } from '@/api';
const router = useRouter()
// const isLoggedIn = useIsLoggedIn()
const user_id : number = getCookies("user_id")
const user_avatar = ref(getCookies("avatar"))
const username = ref<string>(getCookies("username"))
const description = ref<string>(getCookies("description"))
const dialogFormVisible = ref(false)
const formLabelWidth = '140px'
const form = reactive({
  name: username.value,
  avatar_url: user_avatar.value,
  description: description.value,
  old_password:"",
  new_password:"",
})
function logout()  {
    // removeCookies("username")
    // removeCookies("avatar")
    // removeCookies("user_id")
    // removeCookies("token")
    setCookies("token", "")
    setCookies("user_id", -1)
    setCookies("username", "")
    setCookies("avatar", "")
    setCookies("description", "")
    location.reload();
}
const update_user = async() => {
    const ret = await User.update({username : form.name, avatar: form.avatar_url, description: form.description, old_password:form.old_password, new_password:form.new_password})
    dialogFormVisible.value = false
    if (ret.data.code == 200) {
        setCookies("username", form.name)
        setCookies("avatar", form.avatar_url)
        user_avatar.value = form.avatar_url
        username.value = form.name
        ElMessage.success("更新成功")
    }
    else {
        form.name = username.value
        form.avatar_url = user_avatar.value
        form.description = description.value
        return ElMessage.error(ret.data.message)
    }
}
const get_user = async() => {
    if (typeof user_id !== 'undefined' && user_id != -1) {
        router.push(`/profile/${user_id}`)
        return
    }
    router.push(`/auth/login`) 
}
const wrong_ans = async() => {
    if (typeof user_id !== 'undefined' && user_id != -1) {
        router.push(`/wrong_answer/${user_id}`)
        return
    }
    router.push(`/auth/login`) 
}
</script>

<style>
.flex-grow {
  flex-grow: 1;
}
</style>