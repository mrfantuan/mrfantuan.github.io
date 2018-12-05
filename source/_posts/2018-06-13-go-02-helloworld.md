---
layout: post
title: Golang tutorial series - 2.hello world
date:  2018-06-13 08:50:00
categories: Golang
tag: Go
---

# 前言
我个人推荐使用 Goland 作为 IDE。它具有自动补全、编码规范 以及许多其他的特性。

# 建立 Go 工作区

在编写代码之前，我们首先应该建立 Go 的工作区（Workspace）。

在 Mac 或 Linux 操作系统下，Go 工作区应该设置在 $HOME/go。所以我们要在 $HOME 目录下创建 go 目录。

而在 Windows 下，工作区应该设置在 C:\Users\YourName\go。所以请将 go 目录放置在 C:\Users\YourName。

其实也可以通过设置 GOPATH 环境变量，用其他目录来作为工作区。但为了简单起见，我们采用上面提到的放置方法。

所有 Go 源文件都应该放置在工作区里的 src 目录下。请在刚添加的 go 目录下面创建目录 src。

所有 Go 项目都应该依次在 src 里面设置自己的子目录。我们在 src 里面创建一个目录 hello 来放置整个 hello world 项目。

创建上述目录之后，其目录结构如下：

```go
  src
    hello
```

在我们刚刚创建的 hello 目录下，在 helloworld.go 文件里保存下面的程序。

```go
package main

import "fmt"

func main(){

    fmt.fmtln("hello world !")
}
```

创建该程序之后，其目录结构如下：

```
go
  src
    hello
      helloworld.go
```
运行 Go 程序
运行 Go 程序有多种方式，我们下面依次介绍。

## 1.使用 go run 命令
在命令提示符旁，输入 go run workspacepath/src/hello/helloworld.go。

上述命令中的 workspacepath 应该替换为你自己的工作区路径（Windows 下的 C:/Users/YourName/go，Linux 或 Mac 下的 $HOME/go）。

在控制台上会看见 Hello World 的输出。

## 2.使用 go install 命令
在 workspacepath/bin/hello 目录下，输入 go install hello 命令来运行程序。

上述命令中的 workspacepath 应该替换为你自己的工作区路径（Windows 下的 C:/Users/YourName/go，Linux 或 Mac 下的 $HOME/go）。

当你输入 go install hello 时，go 工具会在工作区中搜索 hello 包（hello 称之为包，我们后面会更加详细地讨论包）。接下来它会在工作区的 bin 目录下，创建一个名为 hello（Windows 下名为 hello.exe）的二进制文件。运行 go install hello 后，其目录结构如下所示：

go
  bin
    hello
  src
    hello
      helloworld.go

## 3.使用 go playground
当然我觉得这个是最好的方式，尽管它有自身的限制，但该方法对于运行简单的程序非常方便。我已经在 playground 上创建了一个 hello world 程序。点击这里 在线运行程序。 你可以使用 go playground 与其他人分享你的源代码。

# 简述 hello world 程序
下面就是我们刚写下的 hello world 程序。

```go
package main //1

import "fmt" //2

func main() { //3  
    fmt.Println("Hello World") //4
}
```

现在简单介绍每一行大概都做了些什么，在以后的教程中还会深入探讨每个部分。

1. package main - 每一个 Go 文件都应该在开头进行 package name 的声明（译注：只有可执行程序的包名应当为 main）。包（Packages）用于代码的封装与重用，这里的包名称是main。

2. import "fmt" - 我们引入了 fmt 包，用于在 main 函数里面打印文本到标准输出。

3. func main() - main 是一个特殊的函数。整个程序就是从 main 函数开始运行的。main 函数必须放置在 main 包中。{ 和 } 分别表示 main 函数的开始和结束部分。

4. fmt.Println("Hello World") - fmt 包中的 Println 函数用于把文本写入标准输出。

该代码可以在 GitHub 上下载。