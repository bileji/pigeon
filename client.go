package main

import (
    "fmt"
    "net"
)

const (
    addr = "127.0.0.1:3333"
)

func main() {
    conn, err := net.Dial("tcp", addr)
    if err != nil {
        fmt.Println("连接服务端失败:", err.Error())
        return
    }
    fmt.Println("已连接服务器")
    defer conn.Close()
    Client(conn)
}
func Client(conn net.Conn) {
    go func() {
        sms := make([]byte, 1024)
        for {
            _, err := fmt.Scan(&sms)
            if err != nil {
                fmt.Println("数据输入异常:", err.Error())
            }
            conn.Write(sms)

        }
    }()

    go func() {
        for {
            buf := make([]byte, 1024)
            i, err := conn.Read(buf)
            if err != nil {
                fmt.Println("读取服务器数据异常:", err.Error())
                conn.Close()
                break
            }
            fmt.Println("receive", string(buf[:i]))
        }
    }()

    select {
    }
}
