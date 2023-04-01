package myLog

import (
	"fmt"
	"os"
	"time"
)
type MyLog struct {
	FileName string
	FilePath string
	Level Level
	File *os.File
	LoggerChan chan *LoggerContent
}

func (m *MyLog) NewLogger(FileName,FilePath string,Level Level) *MyLog  {
	ml := &MyLog{
		FileName: FileName,
		FilePath: FilePath,
		Level: Level,
	}
	m.initMyLogFileHandler(FileName,FilePath)
	return ml
}
func NewLogger(FileName,FilePath string,Level Level) *MyLog  {
	ml := &MyLog{
		FileName: FileName,
		FilePath: FilePath,
		Level: Level,
	}
	ml.initMyLogFileHandler(FileName,FilePath)
	fmt.Println(ml)
	return ml
}

func (m *MyLog)initMyLogFileHandler(fileName,filePath string) *os.File {
	fullFile := filePath + fileName
	file ,err := os.OpenFile(fullFile,os.O_WRONLY|os.O_APPEND|os.O_CREATE,0777)
	if err != nil {
		panic( ("os.OpenFile fail "))
	}
	m.File = file
	m.LoggerChan = make(chan *LoggerContent,10000)
	go m.ForChanWrite()
	return file
}

func (m *MyLog) Info(format string,args... interface{})  {
	//log := m.NewLogger(fileName,filePath,INFO)
	// 2023-03-03 15:15:15 [info] day4/main.go main.main 疯狂输出
	name,file,line := GetCaller(2)
	levelStr := getLevel(m.Level)
	fileFunInfo := fmt.Sprintf("%s:%s:%d",file,name,line)
	fmt.Println(format,fileFunInfo,levelStr,args)
	m.WriteLog(format,levelStr,fileFunInfo,args)
}
func (m *MyLog) Debug(format string,args... interface{})  {
	name,file,line := GetCaller(2)
	levelStr := getLevel(m.Level)
	fileFunInfo := fmt.Sprintf("%s [%s] %d",file,name,line)
	m.WriteLog(format,levelStr,fileFunInfo,args)
}
func (m *MyLog) Close()  {
	defer m.File.Close()
}

type LoggerContent struct {
	File *os.File
	Content string
}
var LoggerChan chan *LoggerContent
func (m *MyLog) WriteLog(format,levelStr,fileFunInfo string,args... interface{})  {
	if getLevel(m.Level) > levelStr {
		return
	}
	m.spliteFile()

	//初始化一个chan 将值传入
	//LoggerChan = make(chan *LoggerContent,10000)
	select {
	case m.LoggerChan <- &LoggerContent{
		File: m.File,
		Content: fmt.Sprintf(format,args,levelStr,fileFunInfo),
	} :
	default:
		//丢弃
	}

	//_,err := m.File.WriteString(fmt.Sprintf(format,args,levelStr,fileFunInfo))
	//if err != nil {
	//	fmt.Println("Info error ",err)
	//}
}
func (m *MyLog) ForChanWrite()  {
	for  {
		select {
		case logger := <-m.LoggerChan:
			_,_ = logger.File.WriteString(logger.Content)
		default:
			//chan len 满的时候  丢弃内容
			<-m.LoggerChan
		}
	}
}
func (m *MyLog) spliteFile()    {
	fileInfo,err := m.File.Stat()
	if err != nil {
		fmt.Printf("file.Stat err: %s",err)
		return
	}
	fileSize := fileInfo.Size()
	fmt.Println(fileSize,MAX_SIZE)
	if fileSize > MAX_SIZE{
		fmt.Println("fileSize")
		mutex.Lock()
		m.File.Close()
		fileName := m.File.Name()
		os.Rename(fileName,fmt.Sprintf("%s%d",m.File.Name(),time.Now().UnixNano()))
		//fileName := FormatFileName(time.Now().Second())
		m.File = m.initMyLogFileHandler(fileName,m.FilePath)
		mutex.Unlock()
	}
 }


