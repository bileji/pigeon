package yaml

import (
    "sync"
    "os"
    "io/ioutil"
    "gopkg.in/yaml.v2"
    "github.com/bileji/pigeon/libary/config"
)

type Config struct{}

type Container struct {
    Data map[string]interface{}
    sync.Mutex
}

// 读取数据
func (yaml *Config) Reader(filename string) (handler config.Handler, err error) {
    data, err := ReadYaml(filename)
    if err != nil {
        return
    }
    handler = &Container{
        Data: data,
    }
    return
}

// 写入数据
func (yaml *Config) Writer(data []byte) (handler config.Handler, err error)  {
    // todo
    return
}

// 读取yaml数据
func ReadYaml(filename string) (data map[string]interface{}, err error) {
    handler, err := os.Open(filename)
    if err != nil {
        return
    }
    defer handler.Close()

    bytes, err := ioutil.ReadAll(handler)
    if err != nil || len(bytes) < 3 {
        return
    }
    err = yaml.Unmarshal(bytes, data)
    if err != nil {
        return nil, err
    }
    return
}