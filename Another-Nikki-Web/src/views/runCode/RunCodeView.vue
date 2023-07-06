<script setup lang="ts">
import ContentBase from '@/components/ContentBase.vue';
import RunCode from '@/components/RunCode.vue';
</script>

<template>
    <ContentBase>
        <RunCode messageDefault="（可以生成默认代码与随机用例）"
        problemName=""
    :inputAreaMinRow=8
    :inputAreaMaxRow=8
    :run=true
    :defaultCode=true
    :clearCode=true
    :submitCode=false
    :testCode=false
            />
    </ContentBase>
</template>