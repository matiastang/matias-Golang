<!--
 * @Author: tangdaoyong
 * @Date: 2021-02-01 15:36:56
 * @LastEditors: tangdaoyong
 * @LastEditTime: 2021-02-01 18:04:19
 * @Description: Golang安装
-->
# Golang安装

## 安装golang

### 通过brew安装golang

1. 查看`golang`版本。
```
brew search go

==> Formulae
algol68g                                 google-authenticator-libpam
anycable-go                              google-benchmark
arangodb                                 google-java-format
argon2                                   google-sparsehash
aws-google-auth                          google-sql-tool
bogofilter                               googler
cargo-c                                  goolabs
cargo-completion                         goose
cargo-instruments                        gopass
certigo                                  gor
cgoban                                   goreleaser
clingo                                   goreman
django-completion                        gosu
forego                                   gotags
fuego                                    goto
gnu-go                                   gotop
go                                       gource
go-bindata                               govc
go-jira                                  govendor
go-jsonnet                               gowsdl
go-md2man                                gox
go-statik                                gst-plugins-good
go@1.10                                  gx-go
go@1.11                                  hugo
go@1.12                                  jpegoptim
go@1.13                                  katago
go@1.9                                   lego
goaccess                                 lgogdownloader
goad                                     libgosu
gobby                                    mongo-c-driver
gobject-introspection                    mongo-cxx-driver
gobo                                     mongo-orchestration
gobuster                                 mongoose
gocloc                                   pango
gocr                                     pangomm
gocryptfs                                powerline-go
godep                                    protoc-gen-go
goenv                                    pygobject3
gofabric8                                ringojs
goffice                                  spaceinvaders-go
golang-migrate                           spigot
gollum                                   svgo
golo                                     tengo
gom                                      wego
gomplate                                 wireguard-go
goocanvas                                write-good
goofys
==> Casks
acslogo
algodoo
farrago
gogs
goldendict
google-ads-editor
google-backup-and-sync
google-drive-file-stream
google-japanese-ime
google-photos-backup-and-sync
google-trends
google-web-designer
googleappengine
gopanda
indigo
kugoumusic
lego-digital-designer
logos
maltego
marshallofsound-google-play-music-player
mikogo
moefe-google-translate
mongodb-compass
mongodb-compass-isolated-edition
mongodb-compass-readonly
mongodbpreferencepane
netlogo
nosqlbooster-for-mongodb
paragon-extfs
password-gorilla
sogouinput
x2goclient
homebrew/cask-fonts/font-go
homebrew/cask-fonts/font-go-mono-nerd-font
homebrew/cask-fonts/font-inconsolata-go-nerd-font
```
2. 使用`brew`安装对应版本，比如安装`go@1.13`
```
brew install go@1.13
```

### 下载语言包安装

1. 到[GO中文网](https://studygolang.com/dl)选择对应的`Golang`版本下载。
2. 下载完成之后，双击安装。
3. 安装完成之后，查看版本。
```
go version

go version go1.15.7 darwin/amd64
```

### 从源码安装

[github golang](https://github.com/golang/go.git)
[修改并编译golang源码](https://www.jianshu.com/p/3ca7d1a649a8)

## 配置环境变量

`Golang`包含两个重要的环境变量：`GOROOT`和`GOPATH`，`GOROOT`存储了`Go`官方的源码和可执行文件，`GOPATH`存储了第三方的源码和可执行文件（自己的项目代码建议放在该目录下）。`GOROOT`在安装时已自动配置好，我们只需要配置`GOPATH`即可。

### GOPATH

1. 创建`GOPATH`文件夹，打开终端：
```
mkdir -p ~/gopath/{bin,pkg,src}
```
采用`zsh`作为默认的`shell`，故编辑zsh的配置文件：
```
vi ~/.zshenv
vim ~/.zshrc
```
新增如下代码：
```
# Golang PATH
export GOPATH=$HOME/gopath
# Golang BIN
export PATH=$PATH:$GOPATH/bin
```
我们会将`GOPATH/bin`文件夹加入系统环境变量，这样才能保证第三方库的可执行文件可以正常运行。

保存之后，重启终端(或则执行`source ~/.zshrc`)，运行`go env`指令即可验证`GOPATH`是否设置成功。
```
$ go env

```

### Go Modules配置

从`1.11`版本开始，`Golang`引入了新的依赖管理机制`Go Modules`解决长期以来`Go`语言依赖包没有版本控制的缺陷，`Go Modules`依赖的环境变量为`GOPROXY`和`GOSUMDB`，`GOPROXY`用于检索依赖包的信息，`GOSUMDB`用于校验，默认的配置为：
```
GOPROXY="https://proxy.golang.org,direct"
GOSUMDB="sum.golang.org"
```
由于国内屏蔽`google`，故导致这两个域名都无法访问。

对于`GOPROXY`，七牛云做了一个镜像，方便国内开发者使用，项目地址：

```
https://github.com/goproxy/goproxy.cn
​```

对于`GOSUMDB`，`google`官方提供了国内可访问的域名：`http://sum.golang.google.cn`（参见：`https://github.com/golang/go/issues/31755`）。因此，需要重新配置，同样是修改.zshenv文件：
```
vi ~/.zshenv
vim ~/.zshrc
```
加入如下代码：
```
export GOPROXY=https://goproxy.cn,direct
export GOSUMDB=sum.golang.google.cn
```
或者直接通过go指令修改：
```
go env -w GOPROXY=https://goproxy.cn,direct
go env -w GOSUMDB=sum.golang.google.cn
```