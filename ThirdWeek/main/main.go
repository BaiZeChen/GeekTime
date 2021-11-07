package main

import (
	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// 一些具体的逻辑
	})

	var group errgroup.Group

	var server = &http.Server{Addr: ":80", Handler: mux}

	//启动一个http服务
	group.Go(func() error {
		return server.ListenAndServe()
	})

	// 监听信号
	group.Go(func() error {

		signalChan := make(chan os.Signal)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-signalChan:
			return errors.New("接受到关闭信号，请求关闭服务")
		}

	})

	if err := group.Wait(); err != nil {

		// 做一些日志或者其他的事情

		// 关闭程序
		server.Shutdown(nil)

	}

}
