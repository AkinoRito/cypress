package main

import (
	"net"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
	"fmt"
)

type Params struct {	// 注意字段必须是导出（大写）
	Width, Height int;
}

type Rect struct{}

func (r *Rect) Area(p Params, ret *int) error {
	*ret = p.Width * p.Height;
	return nil;
}

func (r *Rect) Perimeter(p Params, ret *int) error {
	*ret = (p.Width + p.Height) * 2
	return nil
}

func chkError(err error) {
	if err != nil {
		log.Fatal(err);
	}
}

func main() {
	fmt.Println("------server------")
	rect := new(Rect);
	// 注册rpc服务
	rpc.Register(rect);
	fmt.Println("完成rpc服务注册")
	// 获取tcpaddr
	tcpaddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:8080");
	chkError(err);
	fmt.Println("完成获取tcpaddr获取")
	//监听端口
	tcplisten, err2 := net.ListenTCP("tcp", tcpaddr)
	chkError(err2)
	fmt.Println("开始监听端口")
	for{
		conn, err3 := tcplisten.Accept();
		if err3 != nil {
			continue;
		}
		// 使用goroutine单独处理rpc连接请求
		// 治理使用jsonrpc进行处理
		go jsonrpc.ServeConn(conn);
	}
}