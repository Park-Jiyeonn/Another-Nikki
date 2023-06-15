# Another-Nikki

## 记录

### 6.15
给前端加了路由，注意事项：
 * 在 $web$ 文件夹里面执行
 ```
 yarn add vue-router
 ```
 * 在 $App.vue$ 一定要有
 ```
 <router-view></router-view>
 ```
 * 在 $main.ts$ 中是
 ```
 import Router from './router'
 app.use(router)
 ```

$webhook.py$ 修改了下，现在返回消息给 $github$ 比较及时了，用了 $python$ 的 $asyncio$，这个用来创建一个异步的任务。

 嗯，就这些

### 6.14
买了域名 [jiyeon.club](jiyeon.club)。给服务器配了 $caddy$，但是网站似乎要备案，目前正在备案中，感觉有点麻烦。不知道是不是这个原因，所以现在还不能通过域名访问网站。

还有，$web\,hook$ 正式启用了，脚本用了 $fastapi$ 框架，本来想用 $gin$ 的，但是还是因为 $fastapi$ 更方便，所以放弃 $gin$ 了，亏我还和 $gpt$ 交流了好久，问他怎么用 $go$ 执行命令行方便。

前几天前端加了操作消息播报，虽然能跑起来，但是 $vs\,code$ 会报错，网上查到的方法，都很繁琐。。

今天看了下 $inavix$ 的代码，发现只要在 $tsconfig.json$ 里这样改就行了：
```
"moduleResolution": "Node"
```

### 6.11
在铜陵的酒店里对项目的一些微调，后端加了随机留言，前端的随机留言按钮会禁用 $100\,ms$，空留言不能上传。

### 6.9
目前还不是 $yarn run build$ 得到静态文件，然后 $caddy$，所以前端部分我配置了环境变量：

$.env.production$ 是线上的环境变量，其优先级高于 $.env$，低于 $.env.local$。但是
 $.env.local$ 会被 $git$ 忽略，所以线上运行的[http://110.42.239.202:5173/](http://110.42.239.202:5173/) 也就不会受到影响。

 在 $package.json$ 中的 $scripts$，修改 $dev$ 为：

```
"dev": "vite --host 0.0.0.0"
```
这样跨域也能访问了。

随着留言越来越多，展示所有留言并不美观，于是只展示了最后七条留言，主要用到的技术是 $offset$ 配合 $limit$ 一起用。偏移量 + 数量限制。

### 6.6 ~ 6.8
项目从最初的只有 [http://110.42.239.202:8888/](http://110.42.239.202:8888/) 能看到一条 $message$ 到拥有前端，数据库跑在 $docker$ 中。

$clash$ 配置了一下，大概是在官方的 $github$ 下载好，然后 $scp$ 到服务器上去，接着解压到 $usr/local/bin$ 目录里去好像，然后执行了个命令，给了 $clash$ 权限，可以直接 $clash$ 启动。最后配置了一个文件，将其打包为系统服务。

$docker$ 的安装比较简单，用 $ubuntu$ 的 $apt$ 就装好了。

$docker-compose$ 没那么容易，几经辗转，最终用 $windows$ 在官方的 $github$ 下载了文件（就算我配置了 $clash$ 下载起来还是很慢。。。）然后 $scp$ 到服务器上去。像 $clash$ 一样，给了权限之后就可以直接 $docker-compose$ 了。这个不必打包成系统服务。