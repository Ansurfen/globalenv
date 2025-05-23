<h1 align="center">欢迎使用 globalenv 👋</h1>
<center>

[![Go](https://img.shields.io/badge/Go-1.23.5-%2300ADD8.svg?logo=go&logoColor=white&style=for-the-badge)](https://golang.org/)

</center>

---

[English](README.md) | 简体中文

> GlobalEnv 提供了一种简单统一的方式来管理不同操作系统下的全局环境变量。它支持添加、删除、更新和查询环境变量，非常适合管理需要跨系统会话或应用程序持久化的配置值。

## 📦 安装
```cmd
go get "github.com/ansurfen/globalenv"
```

## 🚀 使用
```cmd
package main

import (
    "fmt"
    "github.com/ansurfen/globalenv"    
)

func main() {
    globalenv.Set("globalEnv", "Hello World")

    fmt.Println(globalenv.Get("globalEnv"))
}
```


<!-- description -->

## 🤝 贡献

我们欢迎各种形式的贡献，包括问题报告、功能请求和代码提交。<br />
您可以通过[issues页面](https://github.com/Ansurfen/globalenv/issues)参与项目。<br />
[查看贡献指南](./CONTRIBUTING.md).<br />

## 📝 许可证

本软件采用 MIT 许可，详见 [LICENSE](./LICENSE) 文件。

---

_This Markdown was generated with ❤️ by [docwiz](https://github.com/ansurfen/docwiz)_