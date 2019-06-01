# 安装go插件
获取源码：
> go get -u github.com/golang/protobuf/protoc-gen-go

安装：
> go install github.com/golang/protobuf/protoc-gen-go

# 编译.proto文件
如果`GOPATH`的bin没有加入到环境变量：
> protoc --plugin=protoc-gen-go=$env:GOPATH\bin\protoc-gen-go.exe -I . helloworld.proto --go_out=.

如果加入了环境变量：
> protoc -I . helloworld.proto --go_out=.
> protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld
