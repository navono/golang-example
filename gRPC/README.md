# gRPC
## 环境设置
### 安装go插件
获取源码：
> go get -u github.com/golang/protobuf/protoc-gen-go

安装：
> go install github.com/golang/protobuf/protoc-gen-go

### 编译.proto文件
如果`GOPATH`的bin没有加入到环境变量：
> protoc --plugin=protoc-gen-go=$env:GOPATH\bin\protoc-gen-go.exe -I . helloworld.proto --go_out=.

如果加入了环境变量：
> protoc -I . helloworld.proto --go_out=.
> protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld

### 使用 Docker
> docker pull navono007/proto-compiler:0.1.0

Windows (bash):
> proto = proto-compiler:0.1.0
>
> docker run -v `pwd -W`:/repo $(proto) \
  		protoc -Irepo/proto -Irepo/third_party \
  			--js_out=import_style=commonjs,binary:/repo/frontend/react-grpc/proto \
  			--ts_out=service=true:/repo/frontend/react-grpc/proto \
  			--go_out=plugins=grpc:/repo/backend/go/pkg/api \
  			--grpc-gateway_out=logtostderr=true:/repo/backend/go/pkg/api \
  			--swagger_out=logtostderr=true:/repo/proto/swagger \
  		repo/proto/v1/ping_pong.proto
## TLS
Self signed CA

### Get certstrap
> go get -u -v github.com/square/certstrap

### Create a CA, server cert, and private key
> certstrap init --common-name "My Root CA"
>
> certstrap request-cert --domain mydomain.com

If you’re generating a cert for an IP, use the –ip flag, e.g. --ip 127.0.0.1.

> certstrap sign --CA "My Root CA" mydomain.com # or the IP

At this point you can choose to create a second CA for the client, or just use the same CA to sign another csr. We’ll use the same one for this example.

### Create client cert and private key
> certstrap.exe request-cert --common-name client --ip 127.0.0.1
>
> certstrap sign --CA "My Root CA" client