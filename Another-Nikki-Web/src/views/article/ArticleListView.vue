<template>
    <el-card class="article-card" @click="">
        <el-row :gutter="24">
            <el-col :span="6">
                <el-image fit="cover" class="article-card-image article-card-image-common"
                    src="https://cdn.acwing.com/media/user/profile/photo/104037_lg_820bcf5d13.jpg" lazy>
                </el-image>
            </el-col>
            <el-col :span="18">
                <h3 class="article-title">
                    标题
                </h3>
                <p class="article-description">描述</p>
                <ul class="article-footer">
                    <li class="article-footer-li" @click.stop="">
                        <span class="li-title">2023-6-23</span>
                    </li>
                </ul>
            </el-col>
        </el-row>
    </el-card>
</template>

<style scoped>
.article-card-image {
    width: 100%;
    height: 160px;
}

.article-title {
    display: -webkit-box;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 1;
    overflow: hidden;
}

.article-title:hover {
    color: #337ab7;
}

.article-description {
    display: -webkit-box;
    -webkit-box-orient: vertical;
    -webkit-line-clamp: 2;
    overflow: hidden;
    margin-top: 5px;
    font-size: 14px;
}

.article-footer {
    list-style: none;
    padding: 0px;
    color: var(--el-text-color-secondary);
    font-size: var(--el-font-size-small);
    display: inline-block;
    flex-direction: row;

    .li-title {
        vertical-align: middle;
    }
}
</style>