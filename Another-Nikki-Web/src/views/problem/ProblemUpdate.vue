<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useRoute } from 'vue-router';
import { ElMessage } from 'element-plus';

import { ref } from 'vue'
import { Problem } from '@/types/Problem';
import { ProblemApi } from '@/api';


const route = useRoute();
const router = useRouter();
const problem = ref<Problem>({
    problem_id: 1,
    problem_title: "name",
    problem_content: "content",
    created_time: "CreatedAt",
})
const problem_name = ref('')
const problem_content = ref('')
const loading = ref(false);

import { genFileId } from 'element-plus'
import type { UploadInstance, UploadProps, UploadRawFile, UploadFile } from 'element-plus'
let fileIn = ""
let fileOut = ""
const upload = ref<UploadInstance>();

const get_problem = async (id: number) => {
    const ret = await ProblemApi.get_problem({ problem_id: id })
    if (ret.data.code == 200) {
        problem.value = ret.data.data
        problem_name.value = problem.value.problem_title
        problem_content.value = problem.value.problem_content
    }
    else {
        return ElMessage.error("获取内容失败！")
    }
}

const id: number = parseInt(String(route.params.id));
get_problem(id)

const update_problem = async () => {
    const ret = await ProblemApi.update_problem({ 
        problem_id: id, 
        problem_title: problem_name.value, 
        problem_content:problem_content.value,
        data_in:fileIn,
        data_out:fileOut,
    })
    // console.log(ret.data)
    if (ret.data.code == 200) {
        ElMessage.success(ret?.data?.message ?? 'success')
        return router.push(`/problem/${id}`)
    }
    else {
        return ElMessage.error("更新失败！")
    }
}
const handleExceed: UploadProps['onExceed'] = (files) => {
  upload.value!.clearFiles()
  const file = files[0] as UploadRawFile
  file.uid = genFileId()
  upload.value!.handleStart(file)
}
const handleChange = async (uploadFile: UploadFile, uploadFiles: UploadFile[]) => {
    const file_in = uploadFiles.find(file => file.name.endsWith('.in'));
    const file_out = uploadFiles.find(file => file.name.endsWith('.out'));

    if (file_in) fileIn = await fileToBase64(file_in.raw);
    if (file_out) fileOut = await fileToBase64(file_out.raw);
};

const fileToBase64= async (file: File): Promise<string> => {
  return new Promise((resolve, reject) => {
    const reader = new FileReader();
    reader.onload = () => {
      resolve(reader.result as string);
    };
    reader.onerror = () => {
      reject(reader.error);
    };
    reader.readAsDataURL(file);
  });
}
</script>
<template>
    <el-input v-model="problem_name" :autosize="{ minRows: 2, maxRows: 2 }" type="textarea" placeholder="Please input"
        style="margin-top: 10px;" />
    <el-input v-model="problem_content" :autosize="{ minRows: 10, maxRows: 10 }" type="textarea" placeholder="Please input"
        style="margin-top: 10px;" />
    <el-upload
        ref="upload" class="upload-demo" action=""
        :limit="2"
        :on-exceed="handleExceed"
        :on-change="handleChange"
        :auto-upload="false"
        >
        <template #trigger>
          <el-button type="primary">select file</el-button>
        </template>
    </el-upload>
    <span>如果上传了对应的文件, 则会更新对应的数据</span>
    <div style="text-align: center; margin-top: 10px;">
        <el-button class="button" type="primary" :loading="loading" @click="update_problem()">
            更新
        </el-button>
    </div>
</template>