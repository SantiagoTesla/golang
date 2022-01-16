package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthz(res http.ResponseWriter, req *http.Request) {

	io.WriteString(res, "hello, world!\n")
	//1.接收客户端 request，并将 request 中带的 header 写入 response header
	for k, v := range req.Header {
		//fmt.Fprintln(res, k+":", v)
		for _, res_v := range v {
			fmt.Printf("Header key: %s, Header value: %s \n", k, v)
			res.Header().Set(k, res_v)
		}
	}

	//2.读取当前系统的环境变量中的 VERSION 配置，并写入 response header
	os.Setenv("VERSION", "21.1.1")
	version := os.Getenv("VERSION")
	res.Header().Set("VERSION", version)
	fmt.Printf("VERSION = [%q]\n", version)

	//3.Server 端记录访问日志，包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
	ClientIPAddr := ClientIP(req)
	HttpCode := 200
	log.Printf("Success! clientip: %s", ClientIPAddr)
	log.Printf("Success! Response code: %d", HttpCode)

	//4.当访问 localhost/healthz 时，应返回 200
	res.WriteHeader(200)

}

// ClientIP 尽最大努力实现获取客户端 IP 的算法。
// 解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作。
func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}
	return ""
}
