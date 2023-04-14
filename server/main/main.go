package main

import (
	"fmt"
	"net"
)

func process(conn net.Conn) {

	//延迟关闭连接
	defer conn.Close()

	//交给processor处理
	ps := &Processor{
		Conn: conn,
	}
	ps.processing()

}

func main() {

	//监听8888端口
	fmt.Println("服务器在8888端口开始监听")
	ln, err := net.Listen("tcp", "127.0.0.1:8888")
	if err != nil {
		fmt.Println("服务器监听8888端口错误", err)
		return
	}

	//循环监听
	for {

		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("与客户端连接错误", err)
			return
		}

		//开启协程为客户端服务
		go process(conn)

	}

}
