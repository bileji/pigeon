package main

import (
    "os"
    "fmt"
    "log"
    "github.com/pigeongo/config"
    "github.com/pigeongo/config/yaml"
)

var C yaml.Config

type Data struct {
    Name string `json:"name"`
    Age  int `json:"age"`
}

func main() {
    handler, err := C.Reader("./config/gate.yaml")
    if err != nil {
        fmt.Println(err)
    } else {
        //fmt.Println(handler.Slice("test1", *new([]interface{})))
        fmt.Println(handler.Slice("gade", *new([]interface{})))
    }

    h1, err := config.NewConfig("yaml", "./config/gate.yaml")
    fmt.Println(h1.Slice("gade", *new([]interface{})))

    file, _ := os.Getwd()
    log.Println("current path:", file)


    //bytes, _ := json.Marshal(&Data{Name: "shuchao", Age: 22})
    //h, err := C.Writer(bytes)
    //if err != nil {
    //    fmt.Println(err)
    //} else {
    //    fmt.Println(h.String("name", ""))
    //}

    //b, err := config.NewConfig("yaml", "./glide.yaml")
    //if err != nil {
    //    fmt.Println(err)
    //} else {
    //    fmt.Println(b.String("package", ""))
    //}
}
