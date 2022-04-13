<!--
 * @Author: your name
 * @Date: 2021-09-17 10:17:54
 * @LastEditTime: 2022-04-13 16:29:10
 * @LastEditors: matiastang
 * @Description: In User Settings Edit
 * @FilePath: /matias-Golang/md/Go多版本管理.md
-->
# Go多版本管理

使用[`gvm`](https://github.com/moovweb/gvm)进行`golang`的多版本管理。

## 安装gvm

* bash
```
bash < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
```

* zsh
```
zsh < <(curl -s -S -L https://raw.githubusercontent.com/moovweb/gvm/master/binscripts/gvm-installer)
```

## 下载go

```
$ gvm install go1.4
$ gvm use go1.4 [--default]
```
## gvm指令

### gvm list

查看已下载的`go`版本。

### gvm listall

查看所有可下载的`go`版本。

**编译 Go 1.5+ 的注意事项**
`Go 1.5+` 从工具链中删除了 `C` 编译器，取而代之的是用 `Go` 编写的编译器。显然，如果你还没有一个可用的 `Go` 安装，这会产生一个引导问题。为了编译 `Go 1.5+`，请确保首先安装 `Go 1.4`。
```
$ gvm install go1.4 -B
$ gvm use go1.4
export GOROOT_BOOTSTRAP=$GOROOT
$ gvm install go1.5
```
但是执行`gvm install go1.4 -B`在一些`macOS`上报错。主要是现在需要有一个`go`的环境，现在还没有，所以要想版本安装一个版本的`go`。[issues问题](https://github.com/moovweb/gvm/issues/287)
可以使用`brew`先安装一版`go`
## 解决方法

使用 `Homebrew` 来获取 `Go` 的初始版本：
```
$ brew install go
```
确认您的 `shell` 正在使用 `Homebrew` 版本的路径：
```
$ which go
/usr/local/bin/go
```
使用`gvm`安装`go`版本：
```
$ gvm install go1.11.4
$ gvm use go1.11.4 --default
```
确认您的 `shell` 正在使用`gvm`安装的新版本`go`的路径：
```
$ which go
~/.gvm/gos/go1.11.4/bin/go
```
卸载 `brew` 版本的`go`：
```
$ brew uninstall go
```
