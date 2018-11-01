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

#### Topology

tasks构成的DAG。

#### Scheduler

调度器。

#### Worker

一个goroutine。


## 安装教程

1. xxxx
2. xxxx
3. xxxx

## 使用说明

1. xxxx
2. xxxx
3. xxxx

## 参与贡献

1. Fork 本项目
2. 新建 Feat_xxx 分支
3. 提交代码
4. 新建 Pull Request


## 特技

1. 使用 Readme\_XXX.md 来支持不同的语言，例如 Readme\_en.md, Readme\_zh.md
