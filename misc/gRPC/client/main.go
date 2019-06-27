package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	pb "golang-example/misc/gRPC/api"
	certPath "golang-example/misc/gRPC/cert"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"os"
)

const (
	addr        = "127.0.0.1:8000"
	defaultName = "world"
)

func main() {
	cert, err := tls.LoadX509KeyPair(certPath.ConfPath("client.crt"), certPath.ConfPath("client.key"))
	if err != nil {
		log.Fatalf("client: loadkeys: %s", err)
	}

	ca, err := ioutil.ReadFile(certPath.ConfPath("My_Root_CA.crt"))
	if err != nil {
		panic("Unable to read cert.pem")
	}

	certPool := x509.NewCertPool()
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		log.Fatalf("certPool.AppendCertsFromPEM err")
	}

	c := credentials.NewTLS(&tls.Config{
		Certificates: []tls.Certificate{cert},
		ServerName:   "mydomain.com",
		RootCAs:      certPool,
	})

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(c))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	hello := pb.NewGreeterClient(conn)

	req := pb.HelloRequest{
		Name: defaultName,
	}
	resp, err := hello.SayHello(context.Background(), &req)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "SayHello error, %v\n", err)
		return
	}

	log.Printf("SayHello result: %s", resp.GetMessage())
}
