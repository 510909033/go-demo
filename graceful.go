package main

import (
	"net/http"
	"time"
)

func main() {
	// 创建一个 HTTP 服务
	http.HandleFunc("/", handler)

	// 设置 GracefulClose 超时时间
	time.Sleep(5 * time.Second)

	// 关闭 HTTP 服务
	// httputil.GracefulClose()
}

func handler(w http.ResponseWriter, r *http.Request) {
	// 处理 HTTP 请求
	// ...

	// 关闭连接
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello, world!"))
	w.(http.Flusher).Flush()
}
