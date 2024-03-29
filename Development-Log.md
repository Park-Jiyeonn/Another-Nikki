# Another-Nikki

## 记录
### 2024-03-07
$systemctl$ 不会使用 $shell$ 对话中的环境变量，想要使用环境变量的话，得配置一下 $EnvironmentFile$ 这个东西。

### 2024-02-09
#### 18:47
开始引入 $sqlx$ 了。注意 $wire$ 的使用，这个确实很方便，但是请注意不能循环依赖，否则会报错。

还有， $sqlx$ 要引入官方的 $sql$，否则会报错无法识别 $mysql driver$。

### 2024-02-07
#### 16:43
按照知乎老哥的建议，确实优化了 $Golang$ 的编译速度不少。

#### 11:43
$kratos$ 确实好用，对于目前我的个人开发来说。我在 $pkg$ 中写好了自己封装的 $log$，只要在项目已开始填充到某个位置，日志就生效了。还可以异步地写进 $es$，方便得很。

接下来研究下怎么在代码里打上日志，让代码也能够进入 $es$。

### 2024-01-29
#### 11:39
换 $kratos$ 框架用，比较齐全，很多中间件不用我自己去配置了。

### 2024-01-20
#### 16:33
今天看到 $go-common$ 里面的源码，$trace-id$ 原来是根据时间和随机数生成的（在一个叫 $genId$ 的函数里面），之后还会有相关的全局 $map$，会在起服务的时候帮忙把 $trace-id$ 注入，$http$ 层会用 $metadata.Pairs$ 把 $trace-id$ 放进 $context$。

### 2024-01-13
#### 14:00
装了这么多天的配置，最大的感悟就是，不要去网上直接搜别人的配置怎么配的，因为他们的配置可能早就过时了，就比如这次的 $kafka$，要不是搜到了最新的文章（下次还是不要搜），我根本不知道新版本的会自带 $raft$ 算法。

感觉最靠谱的方法是，去官方文档吧，$kafak$， $canal$ 全都是官方文档的指引，才一步步配好的。

$kafka$：我在众多 $google$ 到的配置里面不知道选哪个 $image$ 好，其实完全可以去 [docker hub](https://hub.docker.com/) 去搜索。而且我很早就在这个网站里面搜索到 [bitnami/kafka](https://hub.docker.com/r/bitnami/kafka) 了，却受网上杂七杂八的配置的影响，没有耐下心去看这个下载数量最多的镜像。

$kafka-ui$：本来也是下载了个杂牌的国产镜像，在自己的 $MacBook$ 上都跑不起来，也是去官方的英文文档去看有什么 $ui$，才搞到了这个镜像。

#### 13:31
又在服务器跑了很多次 $kafka$，阿里服务器的内存太小，稍微多跑几个容器，内存就满了。。得去后台重启。。

现在也算是搞定了，新版本的 $kafka$ 可以像 $zookeeper$ 一样管理集群了，所以可以少安装一个 $zookeeper$ 还是很香的。

$kafka-ui$ 可以在页面上管理这些集群、$topic$，还能直接看到 $topic$ 里面的消息，比原来还要去容器里手动执行命令，看有什么 $topic$，$topic$ 里面的 $message$ 方便多了。

### 2024-01-12 17:37
因为以后 $docker$ 有自己专属的文件夹了，数据库就要迁移哈哈哈哈，记一下迁移的指令吧：
```
docker cp Another-Nikki-data.sql mysql:/
docker exec -it mysql /bin/bash
mysql -uroot -p1234 Another_Nikki<Another-Nikki-data.sql
```

刚刚又想了下，其实不用迁移，$docker-compose.yaml$ 可以写在 $docker$ 文件夹外面。

### 2024-01-12 11:21
昨晚的结论是正确的，尽管 $canal$ 的容器被删除了，但是一些运行时的元数据还保存在本地，就在 ```canal/example/meta.dat``` 里面。

于是我昨晚运行容器的时候，时好时坏。。

### 2024-01-11 23:24
经过了无数次的 $docker-compose.yaml$ 失败（虽然能够安装 $canal$，但是无法成功捕捉到 $binlog$），最终选择在本地安装 $canal$。

其实本地安装不难的，先下载好官方的包，然后跟着文档一步一步来就好了：解压、修改配置（地址，用户名，密码）、启动、查看日志看看有没有启动成功。

我想我现在稍微懂了一点了，之前一直不成功的原因尚不清楚，但极大可能是因为账号没有获取 $binlog$ 的权限。

本地安装以后，本人过于手贱，为了熟悉这个过程，在成功一次以后，还要 $docker-compose down$，这样 $docker$ 里面的 $mysql$ 容器以及数据就被毁了。但是本地的 $canal$ 并不是无状态的啊。。。他会去自动获取 $binlog$ 文件和偏移量，然后保存下来（具体保存在哪里还需要继续一探究竟），所以新启动一个 $mysql$ 容器的时候，我这样就寄了。。。

### 2024-01-09 19:05
$docker$ 容器访问容器的端口，是有讲究的。只要他们在同一个网络中（$eg$.在一个 $docker-compose$ 里面），通过容器名称和端口访问另一个容器的端口通常是可行的。

但是 canal 容器连接 kafka 还是连不上，还需要继续排查问题。

### 2024-01-09 12:09
$working$ 的时候摸鱼ing...拿阿里的服务器测 $docker-compose.yaml$ 写没写对，结果把服务器的 cpu 打满了。。半天才重启完。

现在小领导和大领导在讨论自动化测试调该不该在线上写接口，感觉小领导其实早就知道讨论的结果，但是为了体现自己的思考，总要辩论一番，最后说，好，那我 $OKR$ 改一改。学到一点工作小技巧，这样可以在 leader 面前体现自己的思考？

### 2024-01-07 15:39
这个地方记录一些配置在哪学习的，怎么安装配置也比较合适。

[使用Docker部署Kafka和Zookeeper单机版](https://ken.io/note/kafka-standalone-docker-install#H3-7)

奇怪的是我已经注释了 ```docker-compose.yaml``` 中的 ```volume```，但是最后生成的容器里面，还是有这个东西，不过本地没有。可能是我已经忘记了 ```volume``` 的相关知识了吧。

### 7.8 ~ 12.21
主要加了上传图片的功能，一开始图片上传是没有 $md$ 命名的，很容易重名了，女朋友还发现了这个 $bug$ 哈哈哈。后来都修复了，还花二十多块买了腾讯云的半年的云存储。

十二月底主要是趁着上班摸鱼时间，把编译和运行这两个功能装成一个 $rpc$ 接口。

### 7.7 12:22
今天在往 $oj$ 加以前出的数据的时候，出现了神秘 $bug$，原因是，我为了省事，直接复制了以前在 $windows$ 下生成的数据，但是我的评测机生成的输出文件，却是 $linux$ 的。

这意味着什么呢，$windows$ 下，换行符是在 $ascii$ 码表示是 $10$，而 $linux$ 下换行是 $13$，我的评测机一开始只会忽略 $linux$ 的空格和换行，不会忽略 $windows$ 的回车。于是尴尬的事情就出现了，我提交代码之后，一直显示 $Wrong$ $Answer$。。

还有一个事情，如果加了新数据，$docker$ 是要重新 $build$ 的，因为数据是从本地复制进 $docker$ 镜像的，不是共享的。

这些 $bug$ 让我查了半个上午。。。

### 7.6
今天加上了最想要的功能，$oj$。

前几天加上了留言，关于留言的留言，其实现方式是，设置每条留言的 $rootID$ 与 $parentID$，这样子就可以支持二级留言了。

### 6.28
今天域名终于备案好了，可以启用 $Caddy$ 了。

最近加上了真的博客，后端没什么，前端的主要修改：
* 使用 $markdown-it$ 插件，在 $package.json$ 中添加的依赖，有 $@types$ 前缀的是适配了 $TypeScript$ 的：
  ```
  dependencies: {
    highlight.js
    markdown-it
    markdown-it-katex  // 数学公式
    markdown-it-attrs  // markdown 标签
  }
  devDependencies {
    @types/markdown-it
    @types/markdown-it-attrs
  }
  ```
* 如果是没有适配 $TypeScript$，可以在 $src$ 目录下的 `vite-env.d.ts` 新增，（其实 $VsCode$ 本身也会这样提示）：
  ```
  declare module 'markdown-it-katex';
  ```
* 如果有，要在 $tsconfig.json$ 的 `compilerOptions` 加：
  ```
  "allowSyntheticDefaultImports": true
  ```
* $markdown-it-katex$ 的使用踩坑了，没有仔细看官方文档，以为 `md.use()` 之后就可以了，其实还缺少 $css$ 的引入，否则显示的效果会让人很头痛，以为是代码哪里写错了。
* `.env.production` 修改为 `VITE_API_URL=https://jiyeon.club`，这样流量才能经过 $caddy$ 转发，而且是 $https$ 的。

### 6.23
#### 上午11:21
修好前端的响应了，同时后端的接口返回消息格式都规范了一下。以后都这样写，开发就方便了。

#### 凌晨0:36
部分页面需要特定用户才能查看。

后端 $easy$。

前端的代码风格大大的改进了：
 * 封装了 $axios$，可以很方便的定义一个请求，并且要求后端返回的内容具有特定的格式。这样开发更方便了吧。

    其中，新增的目录有：
    ```
    api : 定义请求，返回值的格式
    hooks : 目前是关于 cookkie 的存储与读取
    http : 封装 axios
    types : 返回值的具体格式，View 中也会用到
    ```
 * 如果本地的 $cookie$ 中包含 $token$，那么会自动带进请求头。

但是前端的消息响应还没做好。

### 6.22
#### 14:31
前端加了路由，但是还没鉴权。

#### 12:36
添加了日志功能，写了个 $gin$ 的中间件，得到请求与响应，然后写进数据库。

### 6.20
#### 16:03
新支持了 $Java$ 和 $Go$，修复了看不到错误信息的问题。

但是 $GO$ 的编译非常慢，要花 $4$ 秒才能编译完。不知道为什么。

#### 13:19
修复了昨晚的两个 $bug$，$python$ 的错误输出也会重定向到文件。主要是利用类似这样的命令：
```
python3 -u py.py > a.out 2>&1
```
后端的输出带上回车了，```join``` 的时候加上 ```\n```即可，前端的 $css$ 样式也做了调整，可以支持回车。

### 6.19
#### 晚23:55
支持运行 $python$ 了，但是还检查不出 $python$ 的编译错误。同时还发现一个 $bug$，输出不会带上回车，嗯，这个有点麻烦。

#### 晚20:49
修了修前端，稍微简化了下代码，主要是调样式，然后配置了```@```指向根目录，主要操作是：
```
在 tsconfig.json 中的编译选项中添加：

"baseUrl": ".",
"paths": {
    "@/*": [ "./src/*" ]
}

在 vite.config.ts 中添加：

import { resolve } from "path"
resolve: {
    alias:{
      '@': resolve(__dirname, 'src'),
    }
}
然后耐心等一会儿，就会好了。
```

### 6.18
#### 晚23:22
搞定输入了，这个做起来居然是最简单的（其实是因为所有的坑前面都踩过了）。就是前端有点丑，不过没事。

#### 晚22：43
下午偶然发现，如果有个人几乎在同时按下运行代码的按钮，由于```docker container```的名字是不变的，所以后按下去的那个会立马返回编译失败，于是给容器名字加上了纳秒时间，这样应该就不重复了吧。

新的 $bug$ 很快又发现了，由于输出的文件名是同一个```a.out```，又会有并发的问题，于是创建新的文件夹，是```tmp-time```的格式，这样子，算是解决问题了吧。然后退出前会删除文件夹。

#### 下午14:53
休息了一会，好像又有动力写下去了，现在可以利用主目录下那个 $calc.cpp$ ，对一个程序进行测量，同时也会限制运行的时间。

关于内存的限制，修改了 $docker$ 的 $run$ 语句，加上 ```-m 256m```即可。```docker```无法限制容器的运行时间，所以这个要靠自己写代码才能做到。

后端的代码稍微改了改，很多重复的地方，都用了函数的封装，这样代码漂亮些吧。比如 $err$ 之后 $c.JSON$
。
#### 中午12:18
在我的 $windows$ 弄了一上午的 ```docker build```，最终还是放弃了，期间尝试过各种国内镜像源，也尝试过 ```form ubuntu:20.04, lastest, 18.04```，但是都寄了，其中```lastest```可以装完镜像，但是无法编译一些特定的头文件。。。

### 6.17
#### 晚 1:29
天哪，为什么会有这种 $bug$，服务器上的系统服务的```os.Getenv("PWD")```的输出，和普通的```go run .```的```os.Getenv("PWD")```输出不一样。。。然后我就一直 ```compile failed```。。。

关键是，我不删除容器，进去看日志，也是毛消息都没得。

最后的解决办法是：
```
os.Getwd()
```

#### 晚0:20
服务器上的这些搞来搞的也太恶心了。。。

首先是解决了 $docker$ 每次都要 ```sudo```的问题，这个 $gpt$ 教我的，比 $google$ 的简单：
```
sudo usermod -aG docker $USER
```
执行完退出再登录就完事了。

#### 晚21:27
服务器在```build``` $docker$ 镜像的时候，一直很慢，我觉得不对啊，我配置了 $clash$ 的，然后一检查 ```curl google.com```也能通，为什么下载的这么慢呢，然后 $google$ 打算给 $docker$ 换源，但是换了还是很慢。。。，

过了很久才意识到，我加载出了一层 $alpine$，然后 ```RUN apt-install```都是在 $alpine$ 里的了，所以 $alpine$ 得换源。。。光是给 $docker$ 换源是不够的。

现在开始着手打包后端的 ```go run .```为一个系统服务：

又踩坑了反正，没写```User=ubuntu Group=ubuntu```，服务就以```root```身份启动。。。然后```root```没配置 $golang$ 的环境，寄。
```
[Unit]
Description=Another Nikki Server
After=network.target

[Service]
ExecStart=/usr/local/go/bin/go run main.go
WorkingDirectory=/home/ubuntu/dox/Another-Nikki/Another-Nikki-Server
Restart=always
User=ubuntu
Group=ubuntu

[Install]
WantedBy=multi-user.target
```

#### 下午 16:42
接口实现起来很 $easy$，主要是用 $go$ 的读写文件。这里用的是 $io$。

关键代码：
```
f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
_, err = io.WriteString(f, code.Code)
```
其中```OpenFile```填的第二个参数很有趣，用了或运算，代表着这三种需求同时满足，估计是为了代码的可读性考虑的吧。```O_TRUNC```表示如果文件已存在，则将其截断为空。如果没有它，那么只会从开头写数据，后面的内容不会被删除。真有趣。

接下来该把服务器上的后端打包成系统服务了，否则他是无法进行```docker```的系列命令的。

#### 下午 15:31
在网上看到，可以用 $linux$ 的系统调用做到。思路是 $fork$ 一个子进程，然后再用 $wait4$ 函数去获取 $CPU$ 时间和内存占用。不过内存使用并不准确，因为会预分配一些空间，但是对于 $acm$ 来说，这个测量精度并不需要太高。

我在本地预先编译了一个 ```.c``` 文件，文件名是 ```calc```，暂时用它来获得运行时间。

到这里，可以开始写后端的接口，和前端的界面了。虽然暂时只支持编译 $cpp$。不过 $py$ 和 $Java$ 也无非是在 $Dockerfile$ 里面 $install$ 一下。

```
docker run --rm --name cpp_run -v $(pwd)/code:/dox cpp_env:1 sh -c "./calc ./cpp > a.out"
```

#### 下午 14：12
看着窗外下着雨，一辆辆车从马路上开过，想起了上周这时候打完蓝桥杯，从上大赶去上海站，这会儿应该还没开始检票。

嗯，经过进一步的学习，发现 ```docker-compose``` 并不是那么适合我的需求。因为之前写的那些配置，都会创建一个新的镜像，然后再创建容器，这样并不好。

现在的实现是：
先创建镜像：
```
docker build -t cpp_env:1 .
```
再启动容器：
```
docker run --rm --name cpp_compile -v $(pwd)/code:/dox cpp_env:1 sh -c "g++ 'c++.cpp' -o 'cpp' -O2 -std=c++11 2> compile.log"
```
稍微解释一下：
 * ```--rm``` 是容器运行完就删除
 * ```--name cpp_compile``` 指定使用的镜像
 * ```-v $(pwd)/code:/dox``` 相当于原来的 $docker-compose.yaml$ 中的 $volume$。
 * ```cpp_env:1``` 容器名称与版本号
 * 编译指令，很显然，$O2$ 优化，程序 $64$ 位，编译错误的信息重定向到 $compile.log$ 

编译完成后，另起一个容器，运行程序：
```
docker run --name cpp_run -v $(pwd)/code:/dox cpp_env:1 sh -c "./cpp > a.out"
```

#### 早 11:22:

昨晚拆分出了编译和运行，用了两个不同的容器，然后关于运行时间和内存使用 ```docker inspect``` 就可以做到。

现在的问题是，每次都用```docker-compose up -d``` 来启动服务，这样子不好，一是因为以后肯定要加其他的语言的容器，每次都全部启动一遍肯定不行；二来是 ```docker-compose```并没有保证启动的顺序。。。很有可能编译没完，运行的容器就启动了，这样肯定不行。

关于第二点，查了下要结合 ```depends_on``` 和自己写命令行，等待编译完成。可这样既麻烦，又让我测量出的容器运行时间远大于程序运行时间。


### 6.16
今天公司的活少，下午领导也走了，于是开始放心摸鱼。请教了旁边一起和我实习的交大的实习生，他告诉我正确使用 $Dockerfile$ 的姿势：
```
docker-compose build
``` 
然后再
```
docker-compose up -d
```
后者只是运行一个构建好的容器（基于一个构建好的镜像）。要加上前者，先 $build$ 出来，才好运行。

不过在生产环境上，肯定是先 $build$ 完，然后每次再 ```docker-compose up -d``` 吧。

其实昨晚弄了一半的 $Dockerfile$ 和 $docker-compose.yaml$ 了，但是囫囵吞枣的，并没有完全理解到底怎么回事。今天才有种悟了的感觉。

然后昨晚也 $compile$ 出一个可执行文件，但是忘记写进 $.gitignore$ 了，已经 ```git add .``` 过的文件，再写进去就来不及了。今天问了旁边的香港大学的实习生为什么 $.gitignore$ 没有忽略文件，他说也不知道为什么，然后就马上 $google$ 帮我查，他是查 ```gitignore absolute path```，有点厉害，然后看了一下，告诉了我为什么。

目前虽然可以编译一个文件了，但是还不会计时和计算程序运行的时候占用的内存。

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

晚 $23:10$，$docker-compose.yaml$ 似乎配置好了，但是 $Dockerfile$ 发现可以去掉不用。用和不用的区别就是，如果用了，创建的镜像名是 $another-nikki-gcc$，现在则直接 $gcc$。

嗯，就这些

### 6.14
买了域名 [jiyeon.club](jiyeon.club)。给服务器配了 $caddy$，但是网站似乎要备案，目前正在备案中，感觉有点麻烦。不知道是不是这个原因，所以现在还不能通过域名访问网站。

还有，$webhook$ 正式启用了，脚本用了 $fastapi$ 框架，本来想用 $gin$ 的，但是还是因为 $fastapi$ 更方便，所以放弃 $gin$ 了，亏我还和 $gpt$ 交流了好久，问他怎么用 $go$ 执行命令行方便。

前几天前端加了操作消息播报，虽然能跑起来，但是 $vs\,code$ 会报错，网上查到的方法，都很繁琐。。

今天看了下 $inavix$ 的代码，发现只要在 $tsconfig.json$ 里这样改就行了：
```
"moduleResolution": "Node"
```

### 6.11
在铜陵的酒店里对项目的一些微调，后端加了随机留言，前端的随机留言按钮会禁用 $100ms$，空留言不能上传。

### 6.9
目前还不是 $build$ 得到静态文件，然后 $caddy$，所以前端部分我配置了环境变量：

$.env.production$ 是线上的环境变量，其优先级高于 $.env$，低于 $.env.local$。但是
 $.env.local$ 会被 $git$ 忽略，所以线上运行的前端服务也就不会受到影响。

 在 $package.json$ 中的 $scripts$，修改 $dev$ 为：

```
"dev": "vite --host 0.0.0.0"
```
这样跨域也能访问了。

随着留言越来越多，展示所有留言并不美观，于是只展示了最后七条留言，主要用到的技术是 $offset$ 配合 $limit$ 一起用。偏移量 + 数量限制。

### 6.6 ~ 6.8
项目从最初的只能在 $IP$ 地址下看到一条 $message$ 到拥有前端，数据库跑在 $docker$ 中。

$clash$ 配置了一下，大概是在官方的 $github$ 下载好，然后 $scp$ 到服务器上去，接着解压到 $usr/local/bin$ 目录里去好像，然后执行了个命令，给了 $clash$ 权限，可以直接 $clash$ 启动。最后配置了一个文件，将其打包为系统服务。

$docker$ 的安装比较简单，用 $ubuntu$ 的 $apt$ 就装好了。

$docker-compose$ 没那么容易，几经辗转，最终用 $windows$ 在官方的 $github$ 下载了文件（就算我配置了 $clash$ 下载起来还是很慢。。。）然后 $scp$ 到服务器上去。像 $clash$ 一样，给了权限之后就可以直接 $docker-compose$ 了。这个不必打包成系统服务。