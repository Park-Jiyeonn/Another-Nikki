<script setup lang="ts">
</script>

<template>
    <div class="home">
        <!-- container是bootstrap定义好的一个类，放进去就漂亮了 -->
        <div class="container">
            <div class="card">
                <!-- 这一行是图片的意思 -->
                <!-- <img src="..." class="card-img-top" alt="..." />   -->
                <div class="card-body">
                    <!-- slot是别人填的东西，会渲染到这里来 -->
                    <slot>

                    </slot>
                </div>
            </div>
        </div>
    </div>
</template>
  
  
<style scoped>
.container {
    /* 距离顶端是20厘米 */
    margin-top: 20px;
}
</style>