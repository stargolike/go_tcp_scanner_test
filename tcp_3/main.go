package main

import (
	"fmt"
	"net"
	"sort"
)
//worker工作
func worker(ports chan int, results chan int){
	for p:= range ports{
		address := fmt.Sprintf("20.194.168.28:%d",p)
		conn, err := net.Dial("tcp",address)
		if err != nil {
			results<- 0
			continue
		}
		conn.Close()
		results <- p

	}
}
//100个worker工作，100个channel，通过循环下达1024个任务。创建results channel收集任务结果。如果结果不为0就放入openports,为0就放入closedports
func main() {
	ports := make(chan int,100)
	results := make(chan int)
	var openports []int
	var closedports []int

	for i:=0;i < cap(ports);i++{
		go worker(ports,results)
	}
	//在func中单独运行
	go func() {
		for i := 1;i < 1024;i++ {
			ports <- i
		}
	}()

 	//main goroutine里运行
	for i := 1;i < 1024;i++ {
		port := <- results
		if port != 0{
			openports = append(openports,port)
		} else {
			closedports = append(closedports,port)
		}
	}
	close(ports)
	close(results)

	//使用sort包排序
	sort.Ints(openports)

	for _,port := range openports {
		fmt.Println("%d open\n",port)
	}
}