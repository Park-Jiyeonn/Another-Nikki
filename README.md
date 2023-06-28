# Another-Nikki

<p align="center">
   <a target="_blank" href="#">
      <img style="display: inline-block;" src="https://img.shields.io/badge/Go-1.20-blue"/>
      <img style="display: inline-block;" src="https://img.shields.io/badge/Gin-v1.9.1-blue"/>
      <img style="display: inline-block;" src="https://img.shields.io/badge/mysql-8.0-blue"/>
      <img style="display: inline-block;" src="https://img.shields.io/badge/GORM-v1.25.1-blue"/>
      <img style="display: inline-block;" src="https://img.shields.io/badge/vue-v3.X-green"/>
    </a>
</p>


## What's inside
个人网站。我想要实现的东西都会放进来。
 * 在线编译器（目前支持 $Cpp$, $Python$, $Java$, $Golang$）
 * 博客
 * 日志（直观的看到哪些 $api$ 被使用了，没钱买 $ElasticSearch$，记录只能写进 $MySQL$）

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
    * 如果想使用在线编译器的功能，还要再 $build$ 出一个 $docker$ 镜像（大概 $1.2GB$）
        ```
        docker build -t oj:1 .
        ```
        嗯，这时候的镜像已经可以编译和运行代码了，但是还不够完整。因为缺少了计算程序运行时间和空间的工具。

        在项目的 $code$ 目录下，$calc.cpp$ 正是这个工具的源码。

        启动一个容器，在 $linux-alpine$ 的环境下编译出 $calc1$：
        ```
        docker run --rm --name cpp_compile -v $(pwd)/code:/dox cpp_env:1 sh -c "g++ 'calc.cpp' -o 'calc1' -O2 -std=c++11 2> compile.log"
        ```

        修改源码中第 $20$ 行为：
        ```
        rl.rlim_cur = rl.rlim_max = 2;
        ```
        然后编译 $calc2$：
        ```
        docker run --rm --name cpp_compile -v $(pwd)/code:/dox cpp_env:1 sh -c "g++ 'calc.cpp' -o 'calc2' -O2 -std=c++11 2> compile.log"
        ```

        最后，让我们的镜像变得完整吧～  
        ```
        docker build -t oj:1 .
        ```

        启动：
        ```
        go run .
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