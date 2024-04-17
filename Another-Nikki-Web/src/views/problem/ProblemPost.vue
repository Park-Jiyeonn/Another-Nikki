<script setup lang="ts">
import { ElMessage } from 'element-plus';

import { ref } from 'vue'
import { ProblemApi } from '@/api';

const problem_name = ref('')
const problem_content = ref('')
const loading = ref(false);
import { genFileId } from 'element-plus'
import type { UploadInstance, UploadProps, UploadRawFile, UploadFile } from 'element-plus'
let fileIn = ""
let fileOut = ""
const upload = ref<UploadInstance>();
const post_problem = async () => {
    const ret = await ProblemApi.post_problem({ 
        problem_title: problem_name.value, 
        problem_content:problem_content.value,
        data_in:fileIn,
        data_out:fileOut,
    })
    // console.log(ret.data)
    if (ret.data.code == 200) {
        ElMessage.success(ret?.data?.message ?? 'success')
    }
    else {
        return ElMessage.error("发布失败！")
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
    <div style="text-align: center; margin-top: 10px;">
        <el-button class="button" type="primary" :loading="loading" @click="post_problem()">
            发布
        </el-button>
    </div>
</template>