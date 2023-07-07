<script setup lang="ts">
</script>

<template>
    <div class="container">
        <el-card class="box-card">
            <slot>
            </slot>
        </el-card>
    </div>
</template>
  
  
<style scoped>
.container {
    margin-top: 20px;
    display: flex;
    justify-content: center;
    align-items: center;
    .box-card{
        width: 960px;
    }
}
</style>