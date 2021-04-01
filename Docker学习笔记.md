# Docker-learning

------



## 00 虚拟化容器介绍



- 虚拟机的不足之处在于对物理服务器资源的消耗，当我们在物理服务器创建一台虚拟机时，便需要虚拟出一套硬件并在上面运行完整的操作系统，每台虚拟机都占用许多的服务器资源。
- `Docker`是使用时下很火的`Golang`语言进行开发的，其技术核心是`Linux`内核的`Cgroup`,`Namespace`和`AUFS`类的`Union FS`等技术，这些技术都是`Linux`内核中早已存在很多年的技术，所以严格来说`Docker`并不是一个完全创新的技术，`Docker`通过这些底层的`Linux`技术，对`Linux`进程进行封装隔离，而被隔离的进程也被称为容器，完全独立于宿主机的进程。
  - ![aTm088](https://cdn.jsdelivr.net/gh/flyingchase/Private-Img@master/uPic/aTm088.png)

- Docker & 虚拟机 比较

  ![image-20210401175039725](https://cdn.jsdelivr.net/gh/flyingchase/Private-Img@master/uPic/image-20210401175039725.png)

 0







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































