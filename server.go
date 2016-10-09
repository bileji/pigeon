package main

import (
    "fmt"
    "net"
)

const (
    //绑定IP地址
    ip = "127.0.0.1"
    //绑定端口号
    port = 3333
)

func main() {
    listen, err := net.ListenTCP("tcp", &net.TCPAddr{IP: net.ParseIP(ip), Port: port, Zone: ""})
    if err != nil {
        fmt.Println("监听端口失败:", err.Error())
        return
    }
    fmt.Println("已初始化连接，等待客户端连接...")
    Server(listen)
}
func Server(listen *net.TCPListener) {
    for {
        conn, err := listen.AcceptTCP()
        if err != nil {
            fmt.Println("接受客户端连接异常:", err.Error())
            continue
        }
        fmt.Println("客户端连接来自:", conn.RemoteAddr().String())
        go func() {
            data := make([]byte, 1024)
            for {
                i, err := conn.Read(data)
                fmt.Println("客户端", conn.RemoteAddr().String(), "发来数据:", string(data[0:i]))
                if err != nil {
                    fmt.Println("读取客户端数据错误:", err.Error())
                    break
                }
                if string(data[:i]) == "exit" {
                    conn.Write([]byte{'e', 'x', 'i', 't'})
                    conn.Close()
                }
                conn.Write(data[0:i])
            }
        }()
    }
}