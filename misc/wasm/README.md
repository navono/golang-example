# WASM

## 编译

根目录运行：

> GOOS=js GOARCH=wasm go build -o ./misc/wasm/main.wasm main.go

## 运行时库

从 `Golang` 的安装目录下的 `misc\wasm` 中，将 `wasm_exec.js` 拷贝到此目录

## 运行

编写入口文件 `index.html`，全局安装 `http-server`，运行

> http-server ./

> Note: http-server 目前有个 bug，访问 root 时会提示 [ERR_INVALID_REDIRECT](https://github.com/http-party/http-server/issues/525) 错误。
> 临时的解决办法是，访问时加上 `index.html`。

打开 `DevTools`，在 `console` 中可看到 `Go` 模块的输出。
