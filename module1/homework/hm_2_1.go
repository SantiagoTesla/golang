package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"
)

func main() {
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthz(res http.ResponseWriter, req *http.Request) {

	//1.接收客户端 request，并将 request 中带的 header 写入 response header
	for k, v := range req.Header {
		fmt.Fprintf(res, "Header[%q] = %q\n", k, v)
	}

	//2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	fmt.Fprintf(res, "Go Version = [%q]\n", runtime.Version())

	//3.Server 端记录访问日志，包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	ClientIPAddr := req.RemoteAddr
	HttpCode := 200
	fmt.Printf("ClientIPAddr = %s\nStatus = %d\n", ClientIPAddr, HttpCode)

	//4.当访问 localhost/healthz 时，应返回 200
	res.WriteHeader(200)

}
