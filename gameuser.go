package main

import (
	"fmt"
	"net"
	"time"
)

//用户
type GameUser struct {
	net.Conn//连接
	UserName string//username
	JoinTime time.Time
}
//设置name
func (user GameUser) SetName(text string) {
	user.UserName = text
	var addr = user.Conn.RemoteAddr().String()
	MapUser[addr] = user
}
//Stats
func (user GameUser) Stats(name string) {
	text := "user [" + name + "] not exist!"
	for username := range MapUser {
		var u =MapUser [ username ]
		if u.UserName == name {
			dur := time.Since(u.JoinTime)
			text = FormatDuration(dur)
			fmt.Println(FormatDuration(dur))
			break
		}
	}
	user.Conn.Write([]byte(text))
}
//收聊天
func (user GameUser) ReceiveMessage(msg ChatMessage) {
	//var text = msg.Text
	var text = msg.SendName + msg.Text
	user.Conn.Write([]byte(text))
}
//格式化时间间隔
func FormatDuration(d time.Duration) string {
	var second int = int(d/time.Second) % 60
	var minute int = int(d/time.Minute) % 60
	var hour int = int(d/time.Hour) % 24
	var day int = int(d/time.Hour/24)
	return fmt.Sprintf("%02dd %02dh %02dm %02ds", day, hour, minute, second)
}