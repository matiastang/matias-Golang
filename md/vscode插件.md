<!--
 * @Author: matiastang
 * @Date: 2022-04-18 16:15:34
 * @LastEditors: matiastang
 * @LastEditTime: 2022-04-18 16:17:54
 * @FilePath: /matias-Golang/md/vscode插件.md
 * @Description: 插件
-->
# 插件

VSCode 必须安装以下插件：
* 首先你必须安装 Golang 插件，然后再给 Go 安装工具包。
* 在 VS Code 中，使用快捷键：command+shift+P，然后键入：go:install/update tools，将所有 16 个插件都勾选上，然后点击 OK 即开始安装。
All tools successfully installed. You’re ready to Go 😃.

## 修改默认配置的方法：

在 Preferences -> Setting 然后输入 go，然后选择 setting.json，填入你想要修改的配置

### 自动完成未导入的包。

"go.autocompleteUnimportedPackages": true,

### VSCode 的一些插件需要配置代理，才能够正常安装。(配合VPN一起使用)

 "http.proxy": "192.168.0.100:1087",

### 如果你遇到使用标准包可以出现代码提示，但是使用自己的包或者第三方库无法出现代码提示，你可以查看一下你的配置项。

 "go.inferGopath": true,

### 如果引用的包使用了 ( . “aa.com/text”) 那这个text包下的函数也无法跳转进去，这是为什么？

修改 “go.docsTool” 为 gogetdoc，默认是 godoc。

"go.docsTool": "gogetdoc",

### 其他
当我们在使用 import 功能的时候，如果无法通过 lint 检查，则不会执行自动 import。
如果你需要自动 import 的前提是你必须把要导入的包的函数写完整。

settings.json
{
  "go.goroot": "",
  "go.gopath": "",
  "go.inferGopath": true,
  "go.autocompleteUnimportedPackages": true,
  "go.gocodePackageLookupMode": "go",
  "go.gotoSymbol.includeImports": true,
  "go.useCodeSnippetsOnFunctionSuggest": true,
  "go.useCodeSnippetsOnFunctionSuggestWithoutType": true,
  "go.docsTool": "gogetdoc",
}