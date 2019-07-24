# mock

## Mock 安装

1. 下载 mock
> go get -u -v github.com/golang/mock/gomock

2. 安装
> cd 到 `mock/mockgen` 目录
>
> go install

3. 运行
> mockgen

## 产生 mock 文件

先创建 `mock` 文件夹，然后执行：

> mockgen.exe -source=student.go -destination=./mock/mock_student.go

编写测试用例，然后执行：

> go test -v

生成覆盖率文件：

> go test -v -cover -coverprofile=cover.out

网页浏览覆盖率文件：

> go tool cover -html=cover.out
