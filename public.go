package public 

import (
	"github.com/edunx/lua"
)

//消息
type Message interface {
	Byte() []byte
}

//自定义对象
type UserData interface { //用户再次读取修改配置的时候 需要转化成userdata
	ToUserData(*lua.LState) *lua.LUserData
}

type Logger interface {
	Err(format string, v ...interface{})
	Info(format string, v ...interface{})
	Debug(format string, v ...interface{})
}

//远程传输方法
type Transport interface {
	UserData

	Start() error
	Close()
	Reload()

	Push(interface{})
}