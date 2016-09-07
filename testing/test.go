package main

import (
    "fmt"
    //"github.com/bileji/pigeon/libary/config"
    "github.com/bileji/pigeon/libary/config/yaml"
    //"encoding/json"
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
