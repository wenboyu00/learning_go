package main

import (
	"fmt"
	"sync"
)

func worker(ports chan int, wg *sync.WaitGroup){
	for p := range ports{
		fmt.Println(p)
		wg.Done()
	}
}

func main(){
	//默认一个缓存
	ports := make(chan int)
}