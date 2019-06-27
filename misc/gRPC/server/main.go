package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	pb "golang-example/misc/gRPC/api"
	certPath "golang-example/misc/gRPC/cert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"net"
	"os"
	"os/signal"
)

const (
	port = "8000"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.Name)
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}

func main() {
	// gRPC server startup options
	var opts []grpc.ServerOption

	// 从证书相关文件中读取和解析信息，得到证书公钥、密钥对
	cert, err := tls.LoadX509KeyPair(certPath.ConfPath("mydomain.com.crt"), certPath.ConfPath("mydomain.com.key"))
	if err != nil {
		log.Fatalf("tls.LoadX509KeyPair err: %v", err)
	}

	// 创建一个新的、空的 CertPool
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(certPath.ConfPath("My_Root_CA.crt"))
	if err != nil {
		log.Fatalf("ioutil.ReadFile err: %v", err)
	}

	// 尝试解析所传入的 PEM 编码的证书。如果解析成功会将其加到 CertPool 中，便于后面的使用
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("certPool.AppendCertsFromPEM err")
	}

	// 构建基于 TLS 的 TransportCredentials 选项
	cred := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},        // 设置证书链，允许包含一个或多个
		ClientAuth:   tls.RequireAndVerifyClientCert, // 要求必须校验客户端的证书
		ClientCAs:    certPool,                       // 设置根证书的集合，校验方式使用 ClientAuth 中设定的模式
	})

	opts = append(opts, grpc.Creds(cred))

	s := grpc.NewServer(opts...)
	pb.RegisterGreeterServer(s, &server{})

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			// sig is a ^C, handle it
			log.Println("shutting down gRPC server...")

			s.GracefulStop()
		}
	}()

	// start gRPC server
	log.Println("starting gRPC server...")

	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalf("failed to Listen: %v", err)
	}

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
