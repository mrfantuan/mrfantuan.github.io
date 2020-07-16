---
layout: post
title: 什么是GitHub
date:   2017-05-01 18:00:00
categories: GitHub
tag: GitHub
---

# 前言
大家可能都听说过Git以及Github的大名，并且我认为Git以及Github是每一个程序员必备的技能。

# Git
## Git简介
咱们先来了解一下什么是Git？

Git官网中对Git的介绍如下：

>Git is a free and open source distributed version control system designed to handle everything from small to very large projects with speed and efficiency.


这段话翻译过来其实就是：
>Git是一个免费并且开源的分布式版本控制系统，被设计用来快速、高效的管理一切从小到大的项目。

由此我们可以发现Git的几个特性：

>分布式
快速
高效

然而这也正是它与传统的SVN这类集中式版本控制系统的区别所在。

## 集中式VS分布式

那么什么是集中式版本控制系统，什么又是分布式版本控制系统呢？这两者又有什么区别呢？

首先先说集中式版本控制系统，它的版本库是存放在中央服务器，但是开发人员用的都是自己的电脑，所以在干活之前，开发人员需要先从中央服务器获取最新的代码，然后才可以工作。工作完了，再把新代码传到中央服务器去。

它最大的缺点就是必须联网才可以干活，没网就可以回家睡觉了。而且如果网速慢的话，你想要把代码从服务器下载下来或者提交代码到服务器都会非常慢。

那么分布式版本控制系统又有何不同呢？首先所谓的分布式版本控制系统，没有所谓的“中央服务器”一说，因为每一台电脑上都有一个完整的版本库，每一台电脑理论上都可以是“中央服务器”。因为每一台电脑都有一个完整的版本库，所以工作时并不需要联网。如果是团队协作的话，只需要把修改的文件推送给对方即可。

那么有的人会说，既然如此，为何还有Git服务器？其实这个服务器只是非常稳定，24小时开机，为了方便团队之间不同的人交换大家的修改而已。没有它一样可以正常的工作，而集中式便不行。所以在安全性上也是分布式的更好，如果某一台电脑坏了，只需要拷贝一份版本库即可。而集中式的服务器如果出了故障那就是很大的问题了。

当然Git相比于SVN这种集中式版本控制系统，并不仅有这一点优势，Git强大的分支管理，快速、高效的处理，便捷的使用，这些优势在使用的过程中你会慢慢感受到！

# GitHub
## GitHub简介

引用WIKI中对Github的介绍：

>GitHub是一个利用Git进行版本控制、专门用于存放软件代码与内容的共享虚拟主机服务。它由GitHub公司（曾称Logical Awesome）的开发者Chris Wanstrath、PJ Hyett和Tom Preston-Werner使用Ruby on Rails编写而成。

Github社区-[https://github.com](https://github.com),也就是我们常说的Github网站。它于2007年10月1日开始开发，网站于2008年2月以Beta版本开始上线，4月份正式上线。从此它成为了全球最大的开源社区，无数优秀的开源代码在这里存放，无数的技术人才在这里汇聚。所以作为一个程序员，Github是必备的技能，这里将是你最好的学习殿堂！


# Git and Github

在分别了解了Git以及Github是什么了之后，很多人可能还有一个疑问，那么就是这两者之间是什么关系？

有些人可能认为Git就是Github，其实这是一个错误的观念。

在上面的介绍中也说了，Github是利用Git来进行版本控制的。所以Git对于Github来说只是一个用来管理项目的工具。Github的功能远远不止于此。

# Git命令速查表
![image](/images/gitcommand.jpg)