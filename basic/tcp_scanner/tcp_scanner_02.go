package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports chan int, result chan int) {
	for p := range ports {
		// Sprintf 格式化返回字符串
		address := fmt.Sprintf("127.0.0.1:%d", p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			result <- 0
			continue
		}
		conn.Close()
		result <- p
	}
}

func main() {
	//make 创建容器，默认为1
	//创建通道容器，类型为int，容量为100
	ports := make(chan int, 100)
	results := make(chan int)
	//创建空切片 nil切片
	var openPorts []int
	var closePorts []int
	//cap 获得容量
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results)
	}

	go func() {
		for i := 1; i < 1024; i++ {
			ports <- 1
		}
	}()
	for i := 1; i < 1024; i++ {
		port := <-results
		if port != 0 {
			openPorts = append(openPorts, port)
		} else {
			closePorts = append(closePorts, port)
		}
	}
	close(ports)
	close(results)

	sort.Ints(openPorts)
	sort.Ints(closePorts)

	for _, port := range closePorts {
		fmt.Println(port)
	}
}
