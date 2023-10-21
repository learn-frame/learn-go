# Learning Go

## CLI

- `go build` 编译并安装自身包和依赖包
- `go install` 主要用于安装非标准库的包文件, 类似于 npm, yarn 这些
- `go fix` 类似于 codemod, 用于将你的 Go 代码从旧的发行版迁移到最新的发行版
- `go test` 单测

## Debug

- `print` / `println` 和 `fmt.Print`, `fmt.Println` 和 `fmt.Printf`
- 在 `fmt.Printf` 中使用下面的说明符来打印有关变量的相关信息:
  - `%+v`: 打印包括字段在内的实例的完整信息
  - `%#v`: 打印包括字段和限定类型名称在内的实例的完整信息
  - `%T`: 打印某个类型的完整说明
- 使用 `panic` 语句来获取栈跟踪信息
- 使用关键字 `defer` 来跟踪代码执行过程

## Go Doc

- `go doc [package]`: 生成 package 包文档
- `go doc [package]/[subpackage]`: 生成 subpackage 子包文档
- `go doc [package] [function]`: 生成 function 函数文档

通过 `godoc -http=:6060` 打开文档.