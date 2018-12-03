---
layout: post
title: redis 安装
date:   2017-05-11 17:00:00
categories: Redis
tag: Redis
---

# 下载地址

redis官方下载地址：
>[https://redis.io/download](https://redis.io/download)

redis64位下载地址：
>[https://github.com/ServiceStack/redis-windows](https://github.com/ServiceStack/redis-windows)

# Window环境安装

1. 解压完成后，目录结构如下:
![image](/images/redis/2-1.png "redis 2-1")

2. 进入文件夹修改 redis.windows.conf文件，设置maxmemory 大小:
![image](/images/redis/2-2.png "redis 2-2")

3. 设置redis密码
![image](/images/redis/2-3.png "redis 2-3")

4. 启动redis
> redis-server.exe redis.windows.conf

![image](/images/redis/2-4.png "redis 2-4")

成功后如图:

![image](/images/redis/2-5.png "redis 2-5")

5. 将redis加入到windows的服务中
> redis-server –-service-install redis.windows.conf -–loglevel verbose

# Unix环境安装

