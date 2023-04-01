package myLog

import (
	"path"
	"fmt"
	"runtime"
	"time"
)
func FormatFileName(second int) (fileName string) {
	tm := time.Now()
	fileName = fmt.Sprintf("%d%d%d",tm.Year(),tm.Month(),tm.Day())
	if second > 0 {
		fileName = fmt.Sprintf("%s%d",fileName,tm.Second())
	}
	fileName = fmt.Sprintf("%s.txt",fileName)
	return fileName
}
func GetCaller(skip int)  (funName,fileName string,line int){
	pc, fileName, line, ok := runtime.Caller(skip)
	if !ok {
		return
	}
	funName = runtime.FuncForPC(pc).Name()
	funName = path.Base(funName)
	fileName = path.Base(fileName)
	fmt.Println(funName,fileName,line)
	return funName,fileName,line
}
