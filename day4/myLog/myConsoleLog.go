package myLog

import (
	"fmt"
	"os"
)

type MyConsoleLog struct {
	Level Level
}

func NewConsoleLogger( Level Level) *MyConsoleLog  {
	ml := &MyConsoleLog{
		Level: Level,
	}
	return ml
}

func (m *MyConsoleLog) Info( format string,args... interface{})  {
	m.WriteLog(format, args)
}
func (m *MyConsoleLog) Debug(format string,args... interface{})  {
	m.WriteLog(format,args)
}
func (m *MyConsoleLog) Close()  {
}
func (m *MyConsoleLog) WriteLog(format string,args... interface{})  {
	name,file,line := GetCaller(3)
	levelStr := getLevel(m.Level)
	fileFunInfo := fmt.Sprintf("%s [%s] %d",file,name,line)
	if getLevel(m.Level) > levelStr {
		return
	}
	_,err := fmt.Fprintf(os.Stdout,format,args,levelStr,fileFunInfo)
	if err != nil {
		fmt.Println("Fprintf error ",err)
	}
}


