package main

import (
	"crypto/tls"
	"crypto/x509"
	"golang_example/gRPC/tls/cert"
	"io/ioutil"
	"log"
	"net/rpc"
)

func main() {
	certKeyPair, err := tls.LoadX509KeyPair(cert.ConfPath("client.crt"), cert.ConfPath("client.key"))
	if err != nil {
		log.Fatalf("client: loadkeys: %s", err)
	}
	//if len(cert.Certificate) != 2 {
	//	log.Fatal("client.crt should have 2 concatenated certificates: client + CA")
	//}

	//ca, err := x509.ParseCertificate(cert.Certificate[0])
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//certPool := x509.NewCertPool()
	//certPool.AddCert(ca)

	certBytes, err := ioutil.ReadFile(cert.ConfPath("My_Root_CA.crt"))
	if err != nil {
		panic("Unable to read cert.pem")
	}
	certPool := x509.NewCertPool()
	ok := certPool.AppendCertsFromPEM(certBytes)
	if !ok {
		panic("failed to parse root certificate")
	}

	config := tls.Config{
		Certificates: []tls.Certificate{certKeyPair},
		ServerName:   "mydomain.com",
		RootCAs:      certPool,
	}

	conn, err := tls.Dial("tcp", "127.0.0.1:8000", &config)
	if err != nil {
		log.Fatalf("client: dial: %s", err)
	}
	defer conn.Close()

	log.Println("client: connected to: ", conn.RemoteAddr())
	rpcClient := rpc.NewClient(conn)
	res := new(clientResult)
	if err := rpcClient.Call("Foo.Bar", "twenty-three", &res); err != nil {
		log.Fatal("Failed to call RPC", err)
	}
	log.Printf("Returned result is %d", res.Data)
}

type clientResult struct {
	Data int
}
