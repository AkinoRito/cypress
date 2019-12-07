package main

import (
	"fmt"
	"log"
	"net/rpc/jsonrpc"
)

type Params struct {	// 注意字段必须是导出（大写）
	Width, Height int;
}

func main() {
	fmt.Println("------client------")
	// 远程连接rpc服务
	//这里使用jsonrpc.Dial
	rpc, err := jsonrpc.Dial("tcp", "127.0.0.1:8080");
	if err != nil {
		log.Fatal(err);
	}
	fmt.Println("连接已建立")
	ret := 0;
	// 远程调用方法
	// 注意第三个参数是指针类型
	err2 := rpc.Call("Rect.Area", Params{50, 100}, &ret);
	if err2 != nil {
		log.Fatal(err2);
	}
	fmt.Println(ret);
	err3 := rpc.Call("Rect.Perimeter", Params{50, 100}, &ret);
	if err3 != nil {
		log.Fatal(err3);
	}
	fmt.Println(ret);
}