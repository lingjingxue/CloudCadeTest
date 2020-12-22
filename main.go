package main

import (
	"bytes"
	"fmt"
	"net"
	"strings"
)
func main() {

	ProfanityInit()
	//TestProfanityInit()
	TestProfanityFilter()

	MapUser = make(map[string]GameUser)
	QueueMessage = make([]ChatMessage, 0)

	fmt.Println("start server...")

	listen, err := net.Listen("tcp", "0.0.0.0:6666")
	if err != nil {
		fmt.Println("listen failed, err:", err)
		return
	}
	for {
		conn, err := listen.Accept()//监听是否有连接
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		go process(conn)//创建goroutine,处理连接
	}
}
func process(conn net.Conn) {

	defer conn.Close()
	for {
		buf := make([]byte, 1024)
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read err:", err)
			return
		}
		buf = bytes.Trim(buf,"\x00")
		var text = string(buf)
		//fmt.Println("read: ", text)
		var user = TryJoinRoom(conn)
		if text == "/popular" {
			user.Popular()
		} else if strings.HasPrefix(text, "setname ") {
			var name = strings.TrimPrefix(text,"setname ")
			user.SetName(name)
		} else if strings.HasPrefix(text, "stats ") {
			var name = strings.TrimPrefix(text,"stats ")
			user.Stats(name)
		} else {
			BroadcastMessage(user, text)
		}
	}
}
