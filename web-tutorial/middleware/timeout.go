package middleware

import (
	"context"
	"net/http"
	"time"
)

type TimeoutMiddleware struct {
	Next http.Handler
}

func (tm TimeoutMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if tm.Next == nil {
		tm.Next = http.DefaultServeMux
	}
	ctx := r.Context()                               // 拿出当前请求的ctx
	ctx, _ = context.WithTimeout(ctx, 3*time.Second) // "修改"当前的ctx，设置一个3秒超时
	r.WithContext(ctx)                               // 用修改后的ctx，代替当前的ctx
	// 如果请求在时间内完成的话，会接受一个信号
	ch := make(chan struct{}) // chan 接受一个struct{}类型
	// 使用gorountine,当serveHTTP执行完成后，就向ch里面发送一个信号，空struct{}
	go func() {
		tm.Next.ServeHTTP(w, r)
		ch <- struct{}{}
	}()
	select {
	// 按时正常处理完
	case <-ch:
		return
		// 超时
	case <-ctx.Done():
		w.WriteHeader(http.StatusRequestTimeout)
	}
	ctx.Done()
}
