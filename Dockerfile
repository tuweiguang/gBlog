FROM golang:1.14
MAINTAINER tuweiguang@foxmail.com

RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,https://goproxy.io,direct

#把宿主机当前上下文的xxx拷贝到容器yyy路径下
#将app代码放到宿主机/data/gblog目录下启动该Dockerfile
COPY . /data/gblog

#设置工作访问时候的WORKDIR路径，登录落脚点
ENV MYPATH /data/gblog
WORKDIR $MYPATH

RUN go build .

#服务端口
EXPOSE 8080
#pprof端口
EXPOSE 6060

#ENTRYPOINT ["./gBlog"]
ENTRYPOINT GOTRACEBACK=crash ./gBlog 1>>./log/$(date +%Y%m%d%H%M%S)"_stdout.log" 2>>./log/$(date +%Y%m%d%H%M%S)"_stderr.log"
