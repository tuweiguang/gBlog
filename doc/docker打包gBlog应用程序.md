# docker打包gBlog应用程序

## 1. 安装docker

方法一：

> centos7自带docker rpm包直接安装

```bash
yum install docker
```

方法二：

[安装docker-ce社区版](https://developer.aliyun.com/article/110806)

## 2. 编写Dockerfile

[官方docker-library/golang](https://hub.fastgit.org/docker-library/docs/tree/master/golang)

```dockerfile
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

ENTRYPOINT ["./gblog"]
```

> - COPY：源路径需要是相对路径，不然失败？？
> - Dockerfile文件要和应用项目在一个目录

**<u>注意：每次重新构建,都需要修改`conf/app.toml`和`conf/db.toml`文件</u>**

```
//conf/app.toml
[app]
    httpAddr = "0.0.0.0:8080"
    pprofAddr = "0.0.0.0:6060"
    ...
[session]
    ...
    domain = "xxx.xxx.xxx.xxx" #ECS主机外网地址
    ...
...    
```

```
//conf/db.toml
[[db]]
    ...
    host = "xxx.xxx.xxx.xxx" #ECS主机本地私有地址
    ...

[[db]]
    dbType = "redis"
    host = "xxx.xxx.xxx.xxx" #ECS主机本地私有地址
    dbNum = 0 #数据库
...
```

## 3. 构建

```bash
docker build -t gblog_image:0.1 .
```

> -  -f /xxx/Dockerfile 指定dockerfile目录，若不指定默认在当前目录下的`Dockerfile`
> - 镜像名字不能有大写字母

出现`Successfully built xxxxxxxxx`则表示构建镜像成功。

## 4. 运行

```bash
docker run -d -it -p 8080:8080 -p 6060:6060 --name my-gblog -v /data/gBlog_log/:/data/gblog/log -v /usr/share/zoneinfo/Asia/Shanghai:/etc/localtime --privileged=true gblog_image:0.1
```

> - -d：后台运行
> - -p：端口映射
> - --name：指定容器名字，若不指定会默认分配随机名字
> - -v：数据卷映射
> - privileged=true：若docker容器的血权限不够，将其值设置为true

注意镜像不加版本默认使用最新版本latest。



## 5. 配置nginx

```
server {
        listen       80;
        server_name  xxx.xxx.xxx.xxx; #ECS公网ip或域名

        #charset koi8-r;

        #access_log  logs/host.access.log  main;

        location / {
            proxy_pass http://127.0.0.1:8080;
            root   html;
            index  index.html index.htm;
        }
        ...
}
```

配置nginx的反向代理，所以这个系统对外暴露80端口即可。

>  nginx刷新配置：`nginx -s reload`

# 6. 推送至阿里云

登陆阿里云搜`容器镜像服务`，然后按照提示创建`个人版实例`，创建完后创建镜像仓库并且与自己的GitHub账户绑定，然后设置自定义自动构建规则，当提交代码至GitHub将自动构建镜像。

1. 登录阿里云Docker Registry

```shell
sudo docker login --username=xxx registry.cn-shanghai.aliyuncs.com
```
> 然后输入登陆阿里云的密码
2. 从Registry中拉取镜像

```shell
sudo docker pull registry.cn-shanghai.aliyuncs.com/gblog/gblog:[镜像版本号]
```
3. 将镜像推送到Registry

```shell
sudo docker login --username=xxx registry.cn-shanghai.aliyuncs.com
sudo docker tag [ImageId] registry.cn-shanghai.aliyuncs.com/gblog/gblog:[镜像版本号]
sudo docker push registry.cn-shanghai.aliyuncs.com/gblog/gblog:[镜像版本号]
```

> 将本地生成的镜像推送至阿里云的私有仓库

## 7. 坑

1. web服务连接MySQL失败？

   [mysql允许外部访问](https://blog.csdn.net/qq_31930499/article/details/100802920)

2. redis服务设置外部访问？

    [redis允许外部访问](https://blog.csdn.net/babybabyup/article/details/85273859)

3. docker时间与本地时间差8小时？

   启动命令增加`-v /usr/share/zoneinfo/Asia/Shanghai:/etc/localtime`

4. 服务启动但是访问不通？

   需要将配置文件`app.toml`中`httpAddr = "localhost:8080"`地址改成`httpAddr = "0.0.0.0:8080"`。

5. 可以访问到登陆首页，但是服务登陆不上？

   将配置文件中设置session的`domain`设置成在浏览器输入的域名或ip。
   
6. 使用nginx反向代理,下游服务获取到的ip都是nginx的ip,而不是真实ip?

    ```
    location / {
                proxy_set_header X-Forward-For $remote_addr;
    	        proxy_set_header X-real-ip $remote_addr;
                ...
            }
    ```

    

## 8. todolist

1. 将标准输出日志打印到日志
2. 开启dump core文件
3. ~~打包镜像，并且提交到阿里云仓库~~
4. ~~docker本地时间设置~~
5. 完善Dockerfile





