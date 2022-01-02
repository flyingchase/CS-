# Docker-learning

------



## 00 虚拟化容器介绍



- 虚拟机的不足之处在于对物理服务器资源的消耗，当我们在物理服务器创建一台虚拟机时，便需要虚拟出一套硬件并在上面运行完整的操作系统，每台虚拟机都占用许多的服务器资源。
- 传统虚拟机技术是虚拟出一套硬件后，在其上运行一个完整操作系统，在该系统上再运行所需应用进程；而容器内的应用进程直接运行于宿主的内核，容器内没有自己的内核，而且也没有进行硬件虚拟。
- `Docker`是使用时下很火的`Golang`语言进行开发的，其技术核心是`Linux`内核的`Cgroup`,`Namespace`和`AUFS`类的`Union FS`等技术，这些技术都是`Linux`内核中早已存在很多年的技术，所以严格来说`Docker`并不是一个完全创新的技术，`Docker`通过这些底层的`Linux`技术，对`Linux`进程进行封装隔离，而被隔离的进程也被称为容器，完全独立于宿主机的进程。
  
- ![aTm088](https://cdn.jsdelivr.net/gh/flyingchase/Private-Img@master/uPic/aTm088.png)
  
- Docker & 虚拟机 比较

    ![image-20210401175039725](https://cdn.jsdelivr.net/gh/flyingchase/Private-Img@master/uPic/image-20210401175039725.png)



## 01 Docker基本概念

### 01.01 镜像 Image

- Root 文件系统 提供容器运行时的程序 库 资源 配置等文件和配置参数 

    ```bash
    # 列出镜像
    docker image ls
    # 拉取镜像
    docker pull [选项] [Docker Registry 地址[:端口号]/]仓库名[:标签]
    # 拉取一个镜像，需要指定Docker Registry的地址和端口号，默认是Docker Hub，还需要指定仓库名和标签，仓库名和标签唯一确定一个镜像，而标签是可能省略，如果省略，则默认使用latest作为标签名，另外，仓库名则由作者名和软件名组成。
    # 运行镜像
    docker run -it centos /bin/bash
    # image_name表示镜像名，image_id表示镜像id
    dockere image rm image_name/image_id
    ```
    
- `docker images` ：列出 docker host 机器上的镜像，可以使用 `-f` 进行过滤

- `docker build`：从 Dockerfile 中构建出一个镜像

- `docker history`：列出某个镜像的历史

- `dockerimport`：从 tarball 中创建一个新的文件系统镜像

- `docker pull`：从 docker registry 拉去镜像

- `docker push`：把本地镜像推送到 registry

- `docker rmi`：删除镜像

- `docker save`：把镜像保存为 tar 文件

- `docker search`：在 docker hub 上搜索镜像

- `docker tag`：为镜像打上 tag 标记

### 01.02 容器 container

镜像是生成容器的模版 

**镜像&&容器** 

​	面向对象中的 类与对象的关系 

```bash
# 查看容器
docker container ls
docker ps
# 删除容器
docker rm container_id # (docker ps 可以看见container_id)
# 进入容器
docker exec -it container_id command
```



### 01.03 仓库 Repository

类似GitHub进行集中储存和分发镜像的服务 

默认Docker Hub `hub.docker.com`	



### 01.04 Docker架构

使用C/S模型 通过客客户端调用服务端 Docker 客户端与 Docker 服务器进行交互，Docker服务端负责构建、运行和分发 Docker 镜像。

![image-20210401182157890](https://cdn.jsdelivr.net/gh/flyingchase/Private-Img@master/uPic/image-20210401182157890.png)



Docker Daemon 是服务器组建 为守护进程 响应客户端的请求并翻译为系统调用完成容器的管理操作



## 02 打包程序

**普通打包方式: **	

以js打包为例

- start with an OS
- Install Node
- Copy app files
- Run node app.js



**Dockerfileff打包:**

- 首字母D大写 其他均小写 无后缀
- FROM 镜像文件



```shell
docker build -t hellodocker .
docker images
docker image ls //列举出来的tag可以进行版本公职 versioning 
```



```shell
# 回顾基础Linux操作commad
# 查看文件前2行
head -n 2 *.txt
# 查看文件最后2行
tail -n 2 *.txt
# less more查看长文件
# 重定向 >
cat # 串联+合并 
```

















