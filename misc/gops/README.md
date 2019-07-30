# gops

## 安装

> go get -u -v github.com/google/gops

## 增加 agent

参考 `main.go`，`agent` 其他选项参考文档

## 分析

列出所有 go 程序

> gops

显示以下内容：
- PID
- PPID
- 程序名称
- 构建该程序的 Go 版本号
- 程序所在绝对路径

程序名后面带 `*` 表示程序加入了 `gops` 的诊断分析代码。


其他的命令可以参考 `gops` 的帮助文档。
