package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	userInfo "proto_demo/proto/proto"
)

func main()  {
	u := userInfo.UserInfo{
		Username: "张三",
		Age: 20,
		Hobby: []string{"aa","22"},
	}
	fmt.Println(u.GetUsername())

	//序列化 proto
	data,_ := proto.Marshal(&u)
	fmt.Println(data)
	//proto 反序列化

	user := userInfo.UserInfo{}
	proto.Unmarshal(data,&user)
	fmt.Printf("%#V",user)
	fmt.Println(user.GetUsername())
}

