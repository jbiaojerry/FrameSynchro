package main
import person "person.pb"
import proto "code.google.com/p/goprotobuf/proto"
import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("server start.")
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Println(err)
		return
	}
	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go echoFunc(conn)
	}
}

func echoFunc(c net.Conn) {
	buf := make([]byte, 1024)

	for {
		n, err := c.Read(buf)
		if err != nil {
			log.Println(err)
			return
		}

		//fmt.Println(string(buf[:n]))
		if len(buf[:n]) > 0{
			toPerObj := &person.Person{}
			proto.Unmarshal(buf[:n], toPerObj)
			fmt.Println("ID:", toPerObj.GetId())
			fmt.Println("名字:", toPerObj.GetName())
			fmt.Println("邮箱地址:", toPerObj.GetEmail())
			c.Write([]byte(toPerObj.GetName()))
			c.Write([]byte("欢迎你~"))
			c.Write([]byte(c.RemoteAddr().String()))
		}
	}
}
