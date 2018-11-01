# stargo

## 项目介绍

STAR(Simultaneous Task Autotuning Runtime) implemented with Golang

## 软件架构

star是一个基于DAG调度的并行优化引擎，目前有共享内存（star\_tp）和分布式内存（star-d、star-thrift）的实现。

现有版本的稳定性不确定，共享内存版本和分布式内存版本的接口不一致，而且有一些BUG。

由于Golang对并行编程和异步通信的天然支持，我决定采用Golang实现star的共享内存版本和分布式内存版本，为用户提供统一的接口。

### stargo modules

#### Task

对任务的封装。

#### Dag

tasks构成的Dag。

#### Pool

Goroutine池，处理Dag。

## 使用说明

### example

#### demo_01

demo_01是一个简单的stargo示例程序，Dag图如下图所示：

![demo_01_dag](example/demo_01/demo_01.png)

运行结果：

![demo_01_gif](example/demo_01/demo_01.gif)

## 参与贡献

1. Fork 本项目
2. 新建 Feat_xxx 分支
3. 提交代码
4. 新建 Pull Request


## 特技

1. 使用 Readme\_XXX.md 来支持不同的语言，例如 Readme\_en.md, Readme\_zh.md
