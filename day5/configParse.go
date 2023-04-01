package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"reflect"
	"strconv"
	"strings"
)

type config struct {
	FileName string `conf:"file_name"`
	FilePath string `conf:"file_path"`
	MaxSize int64 `conf:"max_size"`

}
var err error
func parseConfig(filePath string,config interface{})  {
	//打开文件
	t := reflect.TypeOf(config)
	if t.Kind() != reflect.Ptr {
		err = errors.New("config要是指针")
		return
	}
	tElem := t.Elem()
	if tElem.Kind() != reflect.Struct {
		err = errors.New("必须要是结构体")
	}

	tv := reflect.ValueOf(config)
	fmt.Println(t,tv)
	file,err := os.Open(filePath)
	if err != nil {
		//panic( "os.open error " )
		fmt.Println(err)
		return
	}
	newFile := bufio.NewReader(file)
	var i int
	for {
		i++
		str, err := newFile.ReadString('\n')
		if err == io.EOF {
			fmt.Println("读取完毕",str,err != nil)
			break
		}
		//读取校验
		//#开头的跳出
		has := strings.HasPrefix(str,"#")
		if has {
			fmt.Println("#开头注释 下一个")
			continue
		}
		index := strings.Index(str,"=")
		fmt.Println(str,index)
		if index == -1 {
			fmt.Println("没有=号 不符合格式要求 获取错误配置")
			break
		}
		//sliceStr := strings.Split(strings.TrimSpace(str), "=")
		key := strings.TrimSpace(str[:index])
		value := strings.TrimSpace(str[index+1:])
		if len(key) == 0{
			err = fmt.Errorf("第%d行语法错误 %s",i,key)
			return
		}

		//fmt.Println("行内容",sliceStr,key,value)
		for i:=0;i<tElem.NumField();i++ {
			field := tElem.Field(i)
			tag := field.Tag.Get("conf")
			if key == tag {
				fieldType := field.Type
				//fmt.Println("tElem.Field：",field,tag,fieldType)
				//fmt.Println("tElem.Field：",tv.Field(i),fieldType.Kind())
				switch fieldType.Kind() {
				case reflect.String:
					fieldValue := tv.Elem().FieldByName(field.Name)
					fmt.Println(" sss",field.Name,tv.Elem().Field(i),"	fieldValue\n")
					fieldValue.SetString(value)
					//tv.Elem().Field(i).SetString(value)
				case reflect.Int64,reflect.Int,reflect.Int8,reflect.Int16,reflect.Int32:
					//tv.Field(i)
					value64,_ := strconv.ParseInt(value,10,64)
					tv.Elem().Field(i).SetInt(value64)
				}
			}
		}
	}
}
func main()  {
	var config = &config{}
	parseConfig("./config.conf",config)
	fmt.Println("配置是：",*config)
}
