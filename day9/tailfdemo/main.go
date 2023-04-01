package main

import (
	"fmt"
	"github.com/hpcloud/tail"
)

func main()  {
	filename := "./20230401.log"
	config := tail.Config{
		ReOpen: true,
		Follow: true,
		Location: &tail.SeekInfo{Offset: 0,Whence: 2},
		MustExist: false,
		Poll: true,
	}
	//打开文件
	tails,err := tail.TailFile(filename,config)
	if err != nil {
		fmt.Printf("tail %s failed, err:%v\n",filename,err)
		return
	}

	//开始读取数据
	var (
		msg *tail.Line
		ok bool
	)
	for  {
		msg,ok = <-tails.Lines
		if !ok {
			fmt.Printf("tail file close reopen,filename:%s\n",tails.Filename)
			continue
		}
		fmt.Println("msg：",msg.Text)
	}

}