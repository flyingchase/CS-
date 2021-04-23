# Go 环境配置

平台：macOS(10.15.5)

```bash
brew install go //下载go
go version	//检查go安装
/*
brew默认安装地址 /usr/local/Cellar/go/10.15.2
*/
```

```bash
//配置环境变量
vim ~/.bash_profile
export GOPATH=~/go/
export GOROOT=/usr/local/opt/go/libexec
export PATH=$PATH:$GOPATH/bin
export PATH=$PATH:$GOROOT/bin
//刷新配置文件
source ~/.bash_profile

//查看go的信息
brew info go

//环境变量配置成功与否
go env
```

`source .bash_profile	//终端刷新配置文件，使之生效`

VS Code 中无法install依赖包 

```bash
export http_proxy=http://10.8.0.1:8118;export https_proxy=http://10.8.0.1:8118;  

go get -u -v github.com/mdempsky/gocode
go get -u -v github.com/rogpeppe/godef
go get -u -v golang.org/x/tools/cmd/guru
go get -u -v golang.org/x/tools/cmd/gorename
go get -u -v golang.org/x/tools/cmd/goimports
go get -u -v github.com/uudashr/gopkgs/cmd/gopkgs
go get -u -v github.com/golangci/golangci-lint/cmd/golangci-lint
go get -u -v github.com/ramya-rao-a/go-outline
go get -u -v github.com/acroca/go-symbols
go get -u -v github.com/zmb3/gogetdoc
go get -u -v github.com/fatih/gomodifytags
go get -u -v github.com/cweill/gotests/...
go get -u -v github.com/josharian/impl
go get -u -v github.com/uudashr/gopkgs/cmd/gopkgs
go get -u -v github.com/davidrjenni/reftools/cmd/fillstruct
```



```bash
//无需在bash_profile或者zshsrc中配置环境变量 使用go env命令即可
go env -w GO111MODULE=on //开启 go mod管理包工具
go env -w GOPATH=/User/<username>/go	//配置GOPATH
go env -w GOPROXY=https://goproxy.cn,direct	//配置代理

```

`go env |grep GOPROXY   //具体参数展示`

进入*.go所在文件：

`go build *.go`在同目录下生成exec文件 `go run .go`在终端即可运行



对于go在本地的root目录有：

| **名称** | **说明**          | **目录** | **说明**                |
| -------- | ----------------- | -------- | ----------------------- |
| misc     | 其他文件          | pkg      | 中间文件                |
| api      | 各版本之间api变更 | bin      | 编译器(go) godoc、gofmt |
| src      | 标准库            | test     | 测试                    |
| doc      | go文档            | lib      | 引用库文件              |

go mod命令：

| 命令         | 说明                                                         |
| :----------- | :----------------------------------------------------------- |
| **download** | download modules to local cache(下载依赖包)                  |
| **edit**     | edit go.mod from tools or scripts（编辑go.mod）              |
| **graph**    | print module requirement graph (打印模块依赖图)              |
| **init**     | initialize new module in current directory（在当前目录初始化mod） |
| **tidy**     | add missing and remove unused modules(拉取缺少的模块，移除不用的模块) |
| **vendor**   | make vendored copy of dependencies(将依赖复制到vendor下)     |
| **verify**   | verify dependencies have expected content (验证依赖是否正确） |
| **why**      | explain why packages or modules are needed(解释为什么需要依赖) |

**错误信息：**

- Q:手动本地安装插件失败 不是module或者package

  A：设置GO111MODULE=aoto 配合手动安装插件即可 同时可解决vs code安装出现`GOBIN is not in absolute path` 错误（已经配置好proxy） 

  参考：https://zhuanlan.zhihu.com/p/146970464

  ​			  https://stackoverflow.com/questions/61921282/golang-cannot-find-module-providing-package-package-name-working-directory-is

- Q:开启GO111MODULE 后goland在所配置的gopath/src下新建项目 运行报错  package *** is not in GOROOT 

  A: 使用 `go mod init **` 在Project内 代码中导入 `import "**/moudle"` 即可 或者 `go env -w CO111MOUDLE="off"` 

  参考：https://blog.csdn.net/fbbqt/article/details/105965896

  ​			  https://blog.csdn.net/AdolphKevin/article/details/105480530

  

  

  

  

  

