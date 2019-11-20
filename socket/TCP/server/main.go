package main

import (
	"bufio"
	"fmt"
	proto "github.com/Ronnri/studygo/socket/TCP"
	"io"
	"net"
)

//
//func process(conn net.Conn)  {
//	defer conn.Close()//处理完之后关闭连接
//	//处理当前连接
//	for ; ; {
//		reader := bufio.NewReader(conn)
//		var buf [128]byte
//		n,err:=reader.Read(buf[:])
//		if err!=nil{
//			fmt.Println("read failed",err)
//			break
//		}
//		recv := string(buf[:n])
//		fmt.Println("收到数据:",recv)
//		conn.Write([]byte("ok"))
//
//	}
//
//}
//
//func main() {
//	//1.开启服务
//	listen,err:=net.Listen("tcp","127.0.0.1:20000")
//	if err!=nil{
//		fmt.Println("listen failed",err)
//		return
//	}
//	for ; ; {
//		// 2.等待客户端连接
//		conn,err:= listen.Accept()
//		if err!=nil{
//			fmt.Println("accept failed",err)
//			continue
//		}
//		//3.启动一个单独的goroutine去处理
//		go  process(conn)
//	}
//
//}

func process(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)
	for {
		msg, err := proto.Decode(reader)
		if err == io.EOF {
			return
		}
		if err != nil {
			fmt.Println("decode msg failed, err:", err)
			return
		}
		fmt.Println("收到client发来的数据：", msg)
	}
}

func main() {

	listen, err := net.Listen("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)
	}
}
