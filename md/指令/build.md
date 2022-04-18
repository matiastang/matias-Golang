<!--
 * @Author: matiastang
 * @Date: 2022-04-18 10:23:22
 * @LastEditors: matiastang
 * @LastEditTime: 2022-04-18 11:26:10
 * @FilePath: /matias-Golang/md/指令/build.md
 * @Description: go build
-->
# go build

> go build [-o 输出名] [-i] [编译标记] [包名]

如果参数为XX.go文件或文件列表，则编译为一个个单独的包。
当编译单个main包（文件），则生成可执行文件。
当编译单个或多个包非主包时，只构建编译包，但丢弃生成的对象（.a），仅用作检查包可以构建。
当编译包时，会自动忽略_test.go的测试文件。
**注意** `build`的目录中的`.go`文件，只能属于同一个包(`package *`相同)，不然将报错，如：`found packages carp (carp.go) and main (main.go)`
`package pakName` 此行必须写在第一行，且一个文件夹下的所有文件必须使用同一个包名，一个`.go`文件不能声明两个`package`

```
go build
go build .
go build main.go
# go build + 文件列表
go build test1.go test2.go
# go build + 包(如下，输出文件名为test，包名为main)
go build -o test main
```

## 参数

### -o

* `output` 指定编译输出的名称，代替默认的包名。

### -i

* `install` 安装作为目标的依赖关系的包(用于增量编译提速)。

### 编译标记

* `build` 参数可用在 `build, clean, get, install, list, run, test`
  
```
-v	编译时显示包名
-p n	开启并发编译，默认情况下该值为 CPU 逻辑核数
-a	强制重新构建
-n	打印编译时会用到的所有命令，但不真正执行
-x	打印编译时会用到的所有命令
-race	开启竞态检测
```
```
-a
    完全编译，不理会-i产生的.a文件(文件会比不带-a的编译出来要大？)
-n
    仅打印输出build需要的命令，不执行build动作（少用）。
-p n
    开多少核cpu来并行编译，默认为本机CPU核数（少用）。
-race
    同时检测数据竞争状态，只支持 linux/amd64, freebsd/amd64, darwin/amd64 和 windows/amd64.
-msan
    启用与内存消毒器的互操作。仅支持linux / amd64，并且只用Clang / LLVM作为主机C编译器（少用）。
-v
    打印出被编译的包名（少用）.
-work
    打印临时工作目录的名称，并在退出时不删除它（少用）。
-x
    同时打印输出执行的命令名（-n）（少用）.
-asmflags 'flag list'
    传递每个go工具asm调用的参数（少用）
-buildmode mode
    编译模式（少用）
    'go help buildmode'
-compiler name
    使用的编译器 == runtime.Compiler
    (gccgo or gc)（少用）.
-gccgoflags 'arg list'
    gccgo 编译/链接器参数（少用）
-gcflags 'arg list'
    垃圾回收参数（少用）.
-installsuffix suffix
    ？？？？？？不明白
    a suffix to use in the name of the package installation directory,
    in order to keep output separate from default builds.
    If using the -race flag, the install suffix is automatically set to race
    or, if set explicitly, has _race appended to it.  Likewise for the -msan
    flag.  Using a -buildmode option that requires non-default compile flags
    has a similar effect.
-ldflags 'flag list'
    '-s -w': 压缩编译后的体积
    -s: 去掉符号表
    -w: 去掉调试信息，不能gdb调试了
-linkshared
    链接到以前使用创建的共享库
    -buildmode=shared.
-pkgdir dir
    从指定位置，而不是通常的位置安装和加载所有软件包。例如，当使用非标准配置构建时，使用-pkgdir将生成的包保留在单独的位置。
-tags 'tag list'
    构建出带tag的版本.
-toolexec 'cmd args'
    ？？？？？？不明白
    a program to use to invoke toolchain programs like vet and asm.
    For example, instead of running asm, the go command will run
    'cmd args /path/to/asm <arguments for asm>'.
```

## 跨平台编译

编译跨平台的只需要修改`env`中的`GOOS、GOARCH、CGO_ENABLED`三个环境变量即可
`GOOS`：目标平台的操作系统(`darwin、freebsd、linux、windows`)
`GOARCH`：目标平台的体系架构32位还是64位(`386、amd64、arm`)
交叉编译不支持 `CGO` 所以要禁用它

### Mac 64位可执行程序
```
SET CGO_ENABLED=0
SET GOOS=darwin
SET GOARCH=amd64
go build main.go
```
### Linux 64位可执行程序
```
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64
go build main.go
```