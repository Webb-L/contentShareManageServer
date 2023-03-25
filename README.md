# 内容共享管理服务端

内容共享管理服务端是一个基于HTTP协议的服务器，提供了多个API，方便用户将内容添加到服务器中进行管理和分享。

## 项目简介

内容共享管理服务端是一个功能强大而灵活的服务器，可帮助用户轻松地管理和分享内容。通过访问API，用户可以将内容添加到服务器中，并随时发布和管理它们。

## 主要功能和特点

以下是内容共享管理服务端的主要功能和特点：

+ 多API支持：提供多个API接口，包括添加、发布和管理内容等，以满足用户的各种需求。

+ 简单易用：提供友好的API接口，使用户能够轻松地添加、发布和管理内容。

+ 可扩展性：用户可以自行编写脚本将您喜欢的内容添加到服务器中。

+ 安全可靠：提供身份验证、访问控制和加密传输等多种安全机制，确保用户的内容不受未授权访问和攻击。

## 安装和使用说明

安装和使用内容共享管理服务端非常简单：

1. 下载编译好的版本
   
2. 自行编译

3. docker

4. docker-compose 

### 下载编译好的版本

打开[releases](https://github.com/Webb-L/contentShareManageServer/releases)页面找到您需要的版本。下载成功后启动应用。

```bash
./contentShareManage 

[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

...
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8000
```

### 自行编译

+ 编译

```bash
go build
```

+ 启动

```bash
./contentShareManage 

[GIN-debug] [WARNING] Creating an Engine instance with the Logger and Recovery middleware already attached.

[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

...
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :8000
```

### docker

+ 编译docker镜像

```bash
docker build --tag content_share_manage .

Sending build context to Docker daemon  32.61MB
...
Successfully built acc6b1892ba0
Successfully tagged content_share_manage:latest
```

+ 运行docker镜像

```bash
docker run -d -p 8000:8000 --name content_share_manage content_share_manage

7de1d517bc0258de75229ef5c83e6cd2853be0374139d7db139f20e1cf464150
```

+ 检测是运行成功

```bash
docker ps

CONTAINER ID   IMAGE                  COMMAND   CREATED              STATUS          PORTS                                       NAMES
7de1d517bc02   content_share_manage   "app"     About a minute ago   Up 59 seconds   0.0.0.0:8000->8000/tcp, :::8000->8000/tcp   content_share_manage
```

### docker-compose

+ 编译和启动镜像

```bash
docker-compose up -d 

Building web
Step 1/10 : FROM golang:1.20
...
Removing intermediate container 4fb706158ef6
 ---> 84c5662b8fb1

Successfully built 84c5662b8fb1
Successfully tagged contentsharemanage_web:latest
Creating contentShareMange ... done
```

+ 检测是运行成功

```bash
docker ps

CONTAINER ID   IMAGE                    COMMAND   CREATED         STATUS         PORTS                                       NAMES
00592f236094   contentsharemanage_web   "app"     2 minutes ago   Up 2 minutes   0.0.0.0:8000->8000/tcp, :::8000->8000/tcp   contentShareMange
```

更详细的安装和使用说明，请参阅项目文档。

## 贡献指南

如果您希望为内容共享管理服务端做出贡献，您可以考虑以下几点：

+ 提交问题和建议：如果您在使用中遇到问题或有建议，请提交问题和建议。

+ 编写文档：为项目编写文档，使用户更容易地了解和使用该项目。

+ 提交代码：如果您有编程技能，可以提交代码并改进项目。

更多的贡献指南，请参阅项目贡献指南。

## 许可证

内容共享管理服务端采用Apache许可证开源，更多信息请查看项目许可证。

## 致谢

## JetBrains Support

该项目由 JetBrains 提供支持，JetBrains 为开源项目提供免费的 IDE 许可。

[![JetBrains](https://img.shields.io/badge/Powered%20by-JetBrains-%230000ff.svg)](https://www.jetbrains.com/)