# Another-Nikki

<p align="center">
   <a target="_blank" href="#">
      <img style="display: inline-block;" src="https://img.shields.io/badge/Go-1.20-blue"/>
      <img style="display: inline-block;" src="https://img.shields.io/badge/Kratos-v2.7.2-blue"/>
      <img style="display: inline-block;" src="https://img.shields.io/badge/mysql-8.0-blue"/>
      <img style="display: inline-block;" src="https://img.shields.io/badge/vue-v3.X-green"/>
    </a>
</p>


## What's inside
个人网站。
 * $OnlineJudge$（我出过的题大多都在里面。目前支持 $Cpp$, $Python$, $Java$, $Golang$）
 * 博客、一些题解
 * 留言
 * 日志（只对个人开放权限）

## How to run
* 首先，克隆项目到本地：
    ```
    git clone git@github.com:Park-Jiyeonn/Another-Nikki.git
    ```

 * 然后进入项目的目录：
    ```
    cd Another-Nikki
    ```
    因为有前端和后端，所以这时候要再起一个终端。

    分别进入 $Server$ 和 $Web$
    ```
    cd Another-Nikki-Server
    cd Another-Nikki-Web
    ```
 * 对于 $Server$
    * 先安装依赖：
        ```
        go mod tidy
        ```

    * 然后启动 $Docker$：
        ```
        docker-compose up -d
        ```
    * 如果想使用在线评测机的功能，还要再 $build$ 出一个 $docker$ 镜像（大概 $1.2GB$）
        ```
        docker build -t oj:1 .
        ```
        嗯，这时候的镜像已经可以编译和运行代码了，但是还不够完整。因为缺少了 $online$ $judge$ 的工具。

        在项目的 $online$ $judge$ 目录下，正是这个工具的源码。

        启动一个容器，在 $linux-alpine$ 的环境下编译出 $judger$：
        ```
        docker run --rm --name cpp_compile -v $(pwd)/onlineJudge:/dox oj:1 sh -c "g++ -o judger judger.cpp parse.cpp run.cpp limit.cpp result.cpp -O2 -std=c++11"
        ```

        最后，让我们的镜像变得完整吧～  
        ```
        docker build -t oj:1 .
        ```

        启动评测姬容器：
        ```
        docker run -d -i -m 256m --name oj -v $(pwd)/onlineJudge:/dox oj:1
        ```
    * 后端分为两个微服务：（$main$ 的版本作为朴素的毕设，有三个，必须依赖众多第三方中间件 $Canal$、$Kafka$、$Elastic-Search$、$Kibana$），新的 $mq-less$ 分支剔除了运行内存庞大的第三方中间件。
        启动 $Interact-Hub$：
        ```
        cd interact_hub/service
        go run ./cmd
        ```

        启动 $Judge$ 服务:
        ```
        cd judge/service
        go run ./cmd
        ```
 
 * 对于 $Web$，
    * 先安装依赖：
        ```
        yarn
        ```

    * 然后即可启动：
        ```
        yarn dev
        ```
