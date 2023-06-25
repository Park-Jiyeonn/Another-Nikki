<script setup lang="ts">

</script>

<template>
<el-card class="article-card" >
    <h2 class="article-title">{{ 12 }}</h2>

    <ul class="article-title-footer">
      <li class="article-title-footer-li" >
        <el-tag type="success" effect="dark" size="mini">article-title-footer</el-tag>
      </li>
      <li class="article-title-footer-li" @click="">
        <el-avatar class="avatar li-icon" ></el-avatar>
        <span class="li-title">{{ 4 }}</span>
      </li>
      <li class="article-title-footer-li" @click="">
        <el-icon class="li-icon"><Menu /></el-icon>
        <span class="li-title">{{ 4 }}</span>
      </li>
      <li class="article-title-footer-li">
        <el-icon class="li-icon"><Clock /></el-icon>
        <span class="li-title">{{ 4 }}</span>
      </li>
      <li class="article-title-footer-li">
        <el-icon class="li-icon"><View /></el-icon>
        <span class="li-title">{{ 4 }}</span>
      </li>
    </ul>

    <div class="tag-row">
      <el-tag
        class="tag"
        type="primary"
        effect="dark"
        @click=""
        >{{ 12 }}</el-tag
      >
    </div>

    <div
      element-loading-text="正在获取文章内容..."
      class="article-content"
      v-highlight
      @click=""
    ></div>

    <div
      class="reprint"
      
    >
      本文最后更新于
      {{ 12 }}
    </div>

    <div  class="reprint">
      本文为
      <span class="reprint-author" @click=""
        >格姗知识圈 - {{ 12 }}
      </span>
      原创文章。
      可自由转载、引用，但需署名作者且注明文章出处。如转载至微信公众号，请微信与我联系。<br />
      微信号：geshanintell
    </div>

    <div class="artilce-operate">
      <el-tag
        class="artilce-operate-tag"
        type="primary"
        effect="dark"
        size="large"
        @click=""
        ><svg-icon icon-name="good" />
        <span class="artilce-operate-title">点赞</span>
        <span > ({{ 12 }})</span>
      </el-tag>
      <el-tag
        class="artilce-operate-tag"
        type="warning"
        effect="dark"
        size="large"
        @click=""
        ><el-icon><Star /></el-icon>
        <span class="artilce-operate-title">收藏</span>
        <span >
          ({{ 12 }})</span
        >
      </el-tag>
      <el-tag
        class="artilce-operate-tag"
        type="danger"
        effect="dark"
        size="large"
        @click=""
        ><svg-icon icon-name="money" />
        <span class="artilce-operate-title">打赏</span>
      </el-tag>
    </div>

    <!-- 相关文章推荐 -->
    <about-article :id="1" />

    <!-- 评论 -->
    <div>
      <h3>文章评论</h3>
      <el-divider></el-divider>

      <Comment :article-id="1" />
    </div>

    <!-- 赞赏弹框 -->
    <appreciation ref="appreciationRef" />

    <!-- 正文图片点击显示 -->
    <el-dialog
      class="article-content-image-preview-dialog"
      
      fullscreen
    >
      <div class="dialog-content">
        <img
          style="max-width: 100%"
        />
      </div>
    </el-dialog>
  </el-card>
</template>

<style scoped>
.article-card {
  .article-title {
    color: #009688;
  }

  .article-title-footer {
    width: 100%;
    list-style: none;
    padding: 0px;
    color: var(--el-text-color-secondary);
    font-size: var(--el-font-size-small);
    display: inline-block;
    flex-direction: row;

    .article-title-footer-li {
      float: left;
      padding-right: 15px;
      height: 25px;
      line-height: 25px;
      cursor: pointer;
    }

    .avatar {
      height: 20px;
      width: 20px;
      border-radius: 50%;
    }

    .li-icon {
      vertical-align: middle;
      margin-right: 3px;
    }

    .li-title {
      vertical-align: middle;
    }
  }

  .tag-row {
    clear: both;
    .tag {
      margin: 5px 5px 0px 0px;
      cursor: pointer;
    }
  }

  .article-content {
    margin-top: 30px;
  }

  /*设置图片，以防过大*/
  .article-content >>> img {
    max-width: 100%;
    height: auto;
    cursor: pointer;
  }

  .reprint {
    color: #3c763d;
    background-color: #dff0d8;
    border-color: #d6e9c6;
    padding: 10px 5px;
    margin-top: 20px;

    .reprint-author {
      text-decoration: none;
      color: var(--el-color-primary);
      cursor: pointer;
    }
  }

  .artilce-operate {
    text-align: center;
    margin-top: 20px;

    .artilce-operate-tag {
      font-size: 18px;
      cursor: pointer;
      margin-right: 10px;
    }

    .artilce-operate-title {
      margin-left: 5px;
    }
  }
}
</style>