package master

import (
	"net"
	"net/http"
	"time"
)

// 定义apiserver

// ApiServer 任务的http接口
type ApiServer struct {
	httpServer *http.Server
}

// 单例对象
var (
	G_apiServer *ApiServer
)

// handleJobSave 保存任务接口
func handleJobSave(w http.ResponseWriter, r *http.Request){

}

// 初始化服务
func InitApiServer() (err error){
	var (
		mux *http.ServeMux
		listen net.Listener
		httpServer *http.Server
	)
	// 配置路由
	mux = http.NewServeMux()
	mux.HandleFunc("/job/save", handleJobSave)

	// 启动TCP监听
	if listen, err = net.Listen("tcp",":8070"); err != nil {
		return
	}

	// 创建一个Http服务
	// 解释： 当httpserver收到请求的时候会回调给handler方法 handler已经把路由传递进去了 路由内部根据请求的url遍历他自己的函数进行转发 路由本质上是一种代理模式
	httpServer = http.Server{
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 5 * time.Second,
		Handler: mux,
	}

	// 赋值单例
	G_apiServer = &ApiServer{
		httpServer: httpServer,
	}

	// 让服务跑到协程里
	// 启动了服务端
	go httpServer.Serve(listen)

	return

}