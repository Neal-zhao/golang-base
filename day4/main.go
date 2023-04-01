package main

import (
	"fmt"
	//"fmt"
	"day4/myLog"
	//"myLog"
	"time"
)

var logger myLog.MyLogger
func main()  {
	//分割日志 大小判断 + 重新open + close旧的
	tm := time.Now()
	//todo time.Now().Date 使用
	//fmt.Sprintf("%d%d%d.txt",tm.Year(),tm.Month(),tm.Day())
	fileName := myLog.FormatFileName(0)
	filePath := fmt.Sprintf("./%d%d/%d",tm.Year(),tm.Month(),tm.Day())
	filePath = "./"
	format := "%s [%s] [%s]  人间烟火"
	nowStr := time.Now().Format("2006-01-02 15:04:05")
	//nowStr := time.Now().Format("2006-01-02 15:04:05.000")
	fmt.Println(filePath)
	//ml := myLog.NewLogger(fileName,filePath,myLog.INFO)
	//ml.Info(fileName,filePath,format,nowStr)
	//ml := myLog.NewLogger(fileName,filePath,myLog.DEBUG)
	logger = myLog.NewLogger(fileName,filePath,myLog.DEBUG)
	logger.Debug(format,nowStr)
	logger.Close()
	//defer ml.File.Close()

	logger = myLog.NewConsoleLogger( myLog.DEBUG)
	format = "%s [%s] [%s]  程响 人间烟火"
	logger.Debug(format,nowStr)
}
