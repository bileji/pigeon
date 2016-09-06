package yaml

import (
    "sync"
    "os"
    "fmt"
    "path"
    "time"
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
func (yaml *Config) Writer(data []byte) (config.Handler, error) {
    tmpName := path.Join(os.TempDir(), "pigeon", fmt.Sprintf("%d", time.Now().Nanosecond()))
    os.MkdirAll(path.Dir(tmpName), os.ModePerm)
    if err := ioutil.WriteFile(tmpName, data, 0655); err != nil {
        return nil, err
    }
    return yaml.Reader(tmpName)
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

func init() {
    config.Register("yaml", &Config{})
}
