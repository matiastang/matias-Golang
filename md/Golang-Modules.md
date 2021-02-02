<!--
 * @Author: tangdaoyong
 * @Date: 2021-02-02 09:29:42
 * @LastEditors: tangdaoyong
 * @LastEditTime: 2021-02-02 09:53:41
 * @Description: Modules
-->
# Modules

[【重要】go mod 使用](https://www.cnblogs.com/dhcn/p/11321376.html)

[Go 语言中关于包导入必学的 8 个知识点](https://blog.csdn.net/pyf09/article/details/109165468)
[【重要】20. Go 语言中关于包导入必学的 8 个知识点](https://juejin.im/post/6844904167073382408)

导入的是路径还是包？
当我们使用 import 导入 testmodule/foo 时，初学者，经常会问，这个 foo 到底是一个包呢，还是只是包所在目录名？
import "testmodule/foo"复制代码为了得出这个结论，专门做了个试验（请看「第七点里的代码示例」），最后得出的结论是：

导入时，是按照目录导入。导入目录后，可以使用这个目录下的所有包。
出于习惯，包名和目录名通常会设置成一样，所以会让你有一种你导入的是包的错觉。

总结一下，使用相对导入，有两点需要注意
1. 项目不要放在 $GOPATH/src 下，否则会报错（比如我修改当前项目目录为GOPATH后，运行就会报错）
2. Go Modules 不支持相对导入，在你开启 GO111MODULE 后，无法使用相对导入。

最后，不得不说的是：使用相对导入的方式，项目可读性会大打折扣，不利用开发者理清整个引用关系。

所以一般更推荐使用绝对引用的方式。使用绝对引用的话，又要谈及优先级了