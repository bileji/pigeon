package main

import (
    "log"
    "os"
    "fmt"
    "flag"
)

/**
 * 架框思路
 *  一主多从，master通过消息广播实时同步连接信息到slave，从而保证session一直性
 */
func main() {
    // 自定义Usage
    flag.Usage = func() {
        fmt.Fprintln(os.Stderr, "Usage: session-server [cmd [arg [arg ...]]]")
        flag.PrintDefaults()
    }
    // UDP监听端口
    port := flag.Int("p", 32000, "specifies the listener port")
    // 日志文件保存位置
    logPath := flag.String("l", "/data/logs/session.log", "log file save location")
    flag.Parse()

    log.Println(*port)
    log.Println(*logPath)
}
