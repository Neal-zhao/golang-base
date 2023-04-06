package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

type World struct {

}

//func (w *World) Hello(par interface{})  {
//req string,res *string 参数必须这样写
/**	注意事项
		1、注意参数：只能有两个参数，第二个参数必须是指针类型
			req	表示客户端传过来的数据
			res 表示给客户端返回的数据
		2、注意方法：方法要有个错误返回，方法要是公开的，首字母要大写
		3、注意参数：类型不能是 channel、fun、需要可以进行序列化
*/
func (w *World) Hello(req string,res *string) error  {
	*res = req + "你好"
	fmt.Println("test info rpc :",17)
	return nil
}

type GoodsAbc struct {
}
type AddGoodReq1 struct {
	Id int
	Name string
	Price float64
	ProNum int
}
type AddGoodRes1 struct {
	Id int
	Content string
}
func (g *GoodsAbc)AllAddGoods(req AddGoodReq1,res *AddGoodRes1) error {
	fmt.Println("req: ", req)
	*res = AddGoodRes1{
		Id: 100,
		Content: "成功了",
	}
	return nil
}
func main()  {
	//1、注册 rpc 服务
	//？第一个参数是结构体名，第二个要是指针类型
	err := rpc.RegisterName("World",&World{}) //错误写法
	err = rpc.RegisterName("good",&GoodsAbc{}) //错误写法
	//err := rpc.RegisterName("Hello",new(World))
	//err := rpc.RegisterName("World",new(World)) //这里的参数 是结构体名称
	if err != nil {
		log.Printf("rpc.RegisterName error ：%s", err)
		//这里不用 return 退出
		//return
	}

	//2、监听端口
	listen,err := net.Listen("tcp","127.0.0.1:8000")
	if err != nil {
		log.Printf("net.Listen error ：%s", err)
		return
	}
	//3、应用退出后 关闭监听端口
	defer listen.Close()	//是监听关闭 还是conn关闭 ？？？？？？

	//循环读数据 阻塞式
	for  {
		//4、等待建立连接 有连接往下走 不然阻塞
		conn,err := listen.Accept()
		if err != nil {//这里的错误是退出呢  还是 continue 呢？？？
			log.Printf("listen.Accept error ：%s", err)
			return
		}

		//5、这里是需要绑定服务
		//rpc.ServeConn(conn)
		rpc.ServeCodec(jsonrpc.NewServerCodec(conn))	//采用 json 格式
		//????? 这里需要用协成么 ？？？ 不用这个了 go process(conn)
	}
}
func process(conn net.Conn)  {
	//这里是make  还是 直接类型 ？？？？？ 要make额么
	buf := make([]byte,0)
	n,err := conn.Read(buf)
	if err == io.EOF {
		log.Printf("读取结束 ：%s", err)
		return
	}
	//咋绑定服务
	fmt.Println(buf[:n])
}
