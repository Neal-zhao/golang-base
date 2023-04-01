package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type MyFile struct {}
//打开文件
func (m *MyFile) FileOpenFlag(filePath string) *os.File  {
	file,err := os.OpenFile(filePath,os.O_APPEND|os.O_WRONLY|os.O_CREATE,0776)
	if err != nil {
		fmt.Println(err)
	}
	//defer file.Close()
	return file
}
func (m *MyFile) WriteFileContent(filePath string)  {
	file := m.FileOpenFlag(filePath)
	defer file.Close()
	//bytes := make([]byte,128)
	//bytes := [128]byte{}
	//_,_ = file.Write([]byte{})
	_,_ = file.Write([]byte("1111111111"))
	_,_ = file.WriteString("222222222455555abcsss \n")
	m.ReadBufio(filePath)
}
func (m *MyFile) WriteFor(filePath string)  {
	file := m.FileOpen(filePath)
	defer file.Close()
	//bytes := make([]byte,128)
	bytes := [128]byte{}
	for  {
		n,err := file.Read(bytes[:])
		if err != nil {
			fmt.Printf("err %s\n",err)
		}
		if n == 0 {
			fmt.Printf("长度 %d\n",n)
		}
		if err == io.EOF {
			fmt.Println("结束",err)
			break
		}
		fmt.Println(string(bytes[:]))
	}



}
func (m *MyFile) WriteBufio(filePath string)  {
	file := m.FileOpenFlag(filePath)
	defer file.Close()
	//newFile := bufio.NewReader(file)
	newWriter := bufio.NewWriter(file)
	size,err := newWriter.WriteRune(rune('文'))
	size,err = newWriter.WriteString("rune('文')")
	fmt.Println("size err",size,err)
	newWriter.Flush()
	m.ReadBufio(filePath)
}
func (m *MyFile) WriteFile(filePath string)  {
	file := m.FileOpen(filePath)
	defer file.Close()
	err := ioutil.WriteFile(filePath,[]byte("啥顶顶顶顶顶\n"),0777)
	fmt.Println(err)
	m.ReadAll(filePath)
}
func (m *MyFile) FileOpen(filePath string) *os.File  {

	file,err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}
	//defer file.Close()
	return file
}
//打开文件 读取 || 循环读取
func (m *MyFile) ReadFileContent(filePath string)  {
	file := m.FileOpen(filePath)
	defer file.Close()
	//bytes := make([]byte,128)
	bytes := [128]byte{}
	n,err := file.Read(bytes[:])
	if err != nil {
		fmt.Printf("err %s",err)
	}
	if n == 0 {
		fmt.Printf("长度 %d",n)
	}
	fmt.Print(string(bytes[:]))

}
func (m *MyFile) ReadFor(filePath string)  {
	file := m.FileOpen(filePath)
	defer file.Close()
	//bytes := make([]byte,128)
	bytes := [128]byte{}
	for  {
		n,err := file.Read(bytes[:])
		if err != nil {
			fmt.Printf("err %s\n",err)
		}
		if n == 0 {
			fmt.Printf("长度 %d\n",n)
		}
		if err == io.EOF {
			fmt.Println("结束",err)
			break
		}
		fmt.Println(string(bytes[:]))
	}



}
func (m *MyFile) ReadBufio(filePath string)  {
	file := m.FileOpen(filePath)
	defer file.Close()
	//bytes := make([]byte,128)
	//bytes := [128]byte{}
	newFile := bufio.NewReader(file)
	for  {
		//_,err := file.Read(bytes[:])
		str,err := newFile.ReadString('\n')
		if err == io.EOF {
			fmt.Println("结束",err)
			break
		}
		if err != nil {
			fmt.Printf("err %s\n",err)
			break
		}
		//fmt.Println(string(bytes[:]))
		fmt.Println(str)
	}



}
func (m *MyFile) ReadAll(filePath string)  {
	file := m.FileOpen(filePath)
	defer file.Close()
	//bytes := make([]byte,128)
	//bytes := [128]byte{}
	//newFile := bufio.NewReader(file)
	bytes,_ := ioutil.ReadAll(file)
	fmt.Println(string(bytes))
}

func main()  {
	var f MyFile
	filePath := "./json3.go"
	//filePath = "./abc.txt"
	filePath = "./123.txt"
	//f.ReadFileContent(filePath)
	//f.ReadFor(filePath)
	//f.ReadBufio(filePath)
	//f.ReadAll(filePath)
	//f.WriteFileContent(filePath)
	f.WriteBufio(filePath)
	//f.WriteFile(filePath)
	//fmt.Fprintf()
}
