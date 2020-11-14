package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

// 并发
func main() {
	//声明变量并赋值
	start := time.Now()
	//声明变量 wg 类型sync.WaitGroup
	/*
		main函数相当于启动一个goroutine，在main中启动其他goroutine时，会因为main执行结束其他goroutine直接关闭
		因此，使用sync.WaitGroup在main中做阻塞操作
		操作：启动goroutine之前wg.add(1) 数量+1，在goroutine执行结束时wg.Done() 数量-1。最后 wg.Wait() 计数器归零时执行
	*/
	//wg在每次启动一个goroutine之前+1，goroutine执行的时候-1,
	var wg sync.WaitGroup
	for i := 1; i < 65535; i++ {
		wg.Add(1)
		go func(j int) {
			//defer 在函数体执行结束后按照调用顺序的 相反顺序 逐个执行,发生 严重错误 也会执行
			defer wg.Done()
			address := fmt.Sprintf("20.194.168.28:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				fmt.Printf("%s 关闭了\n", address)
				return
			}
			conn.Close()
			fmt.Printf("%s 打开了\n", address)
		}(i)

	}
	wg.Wait()
	elapsed := time.Since(start) / 1e9
	fmt.Printf("\n\n%d s", elapsed)
}

// 单线程
//func main() {
//	for i := 21; i < 120; i++ {
//		address := fmt.Sprintf("20.194.168.28:%d", i)
//		conn, err := net.Dial("tcp", address)
//		if err != nil {
//			fmt.Printf("%s 关闭了\n", address)
//			continue
//		}
//		conn.Close()
//		fmt.Printf("%s 打开了\n", address)
//	}
//}
