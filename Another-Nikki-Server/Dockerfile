FROM alpine:latest

RUN echo "https://mirrors.ustc.edu.cn/alpine/latest-stable/main" > /etc/apk/repositories \
    && echo "https://mirrors.ustc.edu.cn/alpine/latest-stable/community" >> /etc/apk/repositories \
    && apk update \
    && apk add --no-cache g++ python3 openjdk11 go

WORKDIR /dox

# docker run -d -i -m 256m --name oj -v $(pwd)/onlineJudge:/dox oj:1