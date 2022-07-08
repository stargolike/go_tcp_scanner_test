package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)
func main(){
	start := time.Now()
	var wg sync.WaitGroup
	for i := 21; i <120;i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			address := fmt.Sprintf("20.194.168.28:%d",j)
			conn, err := net.Dial("tcp",address)
			if err != nil {
				fmt.Printf("%s 关闭了\n",address)
				return
			}
			conn.Close()
			fmt.Printf("%s 打开了！！！ \n",address)
		}(i)
	}
	wg.Wait()
	elapsed := time.Since(start)/1e9
	fmt.Printf("\n\n%s seconds elapsed\n", elapsed)
}