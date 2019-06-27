## proto
> go get -u github.com/golang/protobuf/protoc-gen-go
>
> protoc -I api/ api/helloworld.proto --go_out=plugins=grpc:api

## TLS
Self signed CA

### Get certstrap
> go get -u -v github.com/square/certstrap

### Create CA
> certstrap init --common-name "CA"

### Create server cert and self signed

> certstrap request-cert --common-name server --domain mydomain.com

If you’re generating a cert for an IP, use the –ip flag, e.g. --ip 127.0.0.1.

> certstrap sign --CA "CA" server

At this point you can choose to create a second CA for the client, or just use the same CA to sign another csr. We’ll use the same one for this example.

### Create client cert and self signed
> certstrap.exe request-cert --common-name client --domain mydomain.com
>
> certstrap sign --CA "CA" client