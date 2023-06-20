FROM alpine:latest

RUN echo "https://mirrors.ustc.edu.cn/alpine/latest-stable/main" > /etc/apk/repositories \
    && echo "https://mirrors.ustc.edu.cn/alpine/latest-stable/community" >> /etc/apk/repositories \
    && apk update \
    && apk add --no-cache g++ python3 openjdk11 go

COPY ./code /calc

WORKDIR /dox

# 这样会直接让 shell解释器来运行，但是考虑到日后还要运行其他的容器，所以就注释掉吧。
# CMD 命令有且只能有一个，在 Dockerfile 中
# CMD g++ -o cpp c++.cpp && ./cpp > a.out   
