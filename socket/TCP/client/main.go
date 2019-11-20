package main

import (
	"fmt"
	proto "github.com/Ronnri/studygo/socket/TCP"
	"net"
)

//
//func main() {
//	//1.与服务端建立连接
//	conn,err:=net.Dial("tcp","127.0.0.1:20000")
//	if err!=nil{
//		fmt.Println("conn fail",err)
//		return
//	}
//	//2.利用该连接进行数据的发送和接受
//	input := bufio.NewReader(os.Stdin)
//	for ; ; {
//		str,_:=input.ReadString('\n')
//		str = strings.TrimSpace(str)
//		if strings.ToUpper(str)=="Q" {
//			return
//		}
//		//3.给服务端发消息
//		_,err:=conn.Write([]byte(str))
//		if err!=nil{
//			fmt.Println("send fail",err)
//			return
//		}
//		//4.从服务端接收
//		var buf [1024]byte
//		n ,err:= conn.Read(buf[:])
//		if err != nil{
//			fmt.Println("read from server fail",err)
//			return
//		}
//		fmt.Println("recv from server : ",string(buf[:n]))
//
//
//
//	}
//
//}

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		data, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode msg failed, err:", err)
			return
		}
		conn.Write(data)
	}
}
