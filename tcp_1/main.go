package main

import (
	"fmt"
	"net"
)
func main(){
	for i := 21; i <120;i++ {
		address := fmt.Sprintf("20.194.168.28:%d",i)
		conn, err := net.Dial("tcp",address)
		if err != nil {
			fmt.Printf("%s 关闭了\n",address)
			continue
		}
		conn.Close()
		fmt.Printf("%s 打开了！！！ \n",address)
	}
}