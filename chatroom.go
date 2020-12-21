package main

import (
	"fmt"
	"net"
	"time"
)
//用户列表
var MapUser map[string]GameUser
//消息列表
var QueueMessage []ChatMessage

//尝试加入房间
func TryJoinRoom(conn net.Conn) GameUser {

	//fmt.Println("TryJoinRoom ", conn.RemoteAddr())
	var addr = conn.RemoteAddr().String()
	user, ok := MapUser[addr]
	if !ok {
		user = GameUser{conn, addr, time.Now()}
		MapUser[addr] = user
		fmt.Println("TryJoinRoom New User", addr,len(QueueMessage))

		//发送最近50条消息
		var queened = QueueMessage
		if len(queened) > 50 {
			start := len(queened) - 50
			queened = queened[start:]
		}
		for _, v := range queened {
			user.ReceiveMessage(v)
		}
	}
	return user
}

func PopMessage() ChatMessage {
	element := QueueMessage[0]
	if len(QueueMessage) > 1 {
		QueueMessage = QueueMessage[1:]
	} else {
		QueueMessage = make([]ChatMessage, 0)
	}
	return element
}

//广播消息
func BroadcastMessage(user GameUser, text string) {

	text = ProfanityFilter(text)

	if text=="\r\n" {
		return
	}

	//fmt.Println("BroadcastMessage..." + user.UserName + " Say:" + text)
	dts, _ := time.ParseDuration("-5s")
	var drove = time.Now().Add(dts)
	var need = CheckMessageNeedPop(drove)
	for need {
		PopMessage()
		need = CheckMessageNeedPop(drove)
	}
	//添加到消息列表
	var msg = ChatMessage{text, time.Now(), user.UserName}
	QueueMessage = append(QueueMessage, msg)

	//广播
	for username := range MapUser {
		var u =MapUser [ username ]
		u.ReceiveMessage(msg)
	}

	//fmt.Println("QueueMessage...",len(QueueMessage))
}

//判断是否需要删除消息
func CheckMessageNeedPop(drove time.Time) bool {
	if len(QueueMessage) <= 50 {
		return false
	}
	if QueueMessage[0].SendTime.After(drove) {
		return false
	}
	return true
}