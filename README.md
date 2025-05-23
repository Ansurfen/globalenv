<h1 align="center">Welcome to globalenv 👋</h1>
<center>

[![Go](https://img.shields.io/badge/Go-1.23.5-%2300ADD8.svg?logo=go&logoColor=white&style=for-the-badge)](https://golang.org/)

</center>

---

English | [简体中文](README.zh-CN.md)

> GlobalEnv provides a simple and consistent way to manage global environment variables across different operating systems. It allows you to add, remove, update, and query environment variables, making it ideal for managing configuration values that need to persist across system sessions or applications.

## 📦 Install
```cmd
go get "github.com/ansurfen/globalenv"
```

## 🚀 Usage
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

## 🤝 Contributing

Contributions, issues and feature requests are welcome.<br />
Feel free to check [issues page](https://github.com/Ansurfen/globalenv/issues) if you want to contribute.<br />
[Check the contributing guide](./CONTRIBUTING.md).<br />

## 📝 License

This software is licensed under the MIT license, see [LICENSE](./LICENSE) for more information.

---

_This Markdown was generated with ❤️ by [docwiz](https://github.com/ansurfen/docwiz)_