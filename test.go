package main

import (
    "log"
    "time"
    "net"
    "github.com/pigeongo/monitor"
    "github.com/pigeongo/broadcast"
)

const (
    //绑定IP地址
    ip = "127.0.0.1"
    //绑定端口号
    port = 3333
    //超时时间
    timeout = 30
)

//func main() {
//    ticker := time.NewTicker(time.Millisecond * 500)
//    //go func() {
//    //    for t := range ticker.C {
//    //        fmt.Println("Tick at", t)
//    //    }
//    //}()
//
//    go hearBeat(ticker, func(t *time.Ticker) {
//        for ti := range t.C {
//            fmt.Println("Tick at", ti)
//        }
//    })
//
//
//    select {}
//}
//

//  .....
func Master(listener *net.TCPListener, fn func(b *broadcast.Broadcaster, buffer []byte)) {

    tt := time.NewTimer(5 * time.Second)
    select {
    case <- tt.C:
        log.Println("---------")
    }

    b := broadcast.NewBroadcaster()
    for {
        c, err := listener.AcceptTCP()
        if err != nil {
            continue
        }

        // 保持客户端连接
        go func(conn *net.TCPConn, b *broadcast.Broadcaster) {
            buffer := make([]byte, 512)
            for {
                monitor.Diagnose(conn, buffer, time.Second * timeout)
                n, err := conn.Read(buffer)
                if err != nil {
                    // todo 维护连接信息
                    log.Println(err)
                    break
                }
                fn(b, buffer[:n])
            }
        }(c, b)

        // 广播消息
        go func(conn *net.TCPConn, r broadcast.Receiver) {
            for {
                _, err := conn.Write((r.Reader()).([]byte))
                if err != nil {
                    break
                }
            }
        }(c, b.Listen())
    }
}

func main() {
    listener, err := net.ListenTCP("tcp", &net.TCPAddr{IP: net.ParseIP(ip), Port: port, Zone: ""})
    if err != nil {
        log.Println("监听端口失败:", err.Error())
        return
    }
    log.Println("已初始化连接，等待客户端连接...")
    Master(listener, func(b *broadcast.Broadcaster, buffer []byte) {
        log.Println("broadcast ", string(buffer))
        // todo
        b.Writer(buffer)
    })
}

