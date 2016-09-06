package yaml

import (
    "sync"
    "os"
    "fmt"
    "path"
    "time"
    "strings"
    "io/ioutil"
    "gopkg.in/yaml.v2"
    "github.com/bileji/pigeon/libary/config"
    "github.com/wendal/errors"
)

type Config struct{}

type Container struct {
    data map[string]interface{}
    sync.Mutex
}

func (c *Container) Get(key string) (interface{}, error) {
    if len(key) == 0 {
        return nil, errors.New("key is empty")
    }
    if val, ok := c.data[key]; ok {
        return val, nil
    }
    return nil, fmt.Errorf("not exist key %q", key)
}

func (c *Container) Set(key string, val string) error {
    c.Lock()
    defer c.Unlock()
    c.data[key] = val
    return nil
}

func (c *Container) String(key string, def string) string {
    if val, err := c.Get(key); err == nil {
        if v, ok := val.(string); ok {
            return v
        }
    }
    return def
}

func (c *Container) Strings(key string, def []string) []string {
    if val, err := c.Get(key); err == nil {
        if v, ok := val.(string); ok {
            return strings.Split(v, ";")
        }
    }
    return def
}

func (c *Container) Int(key string, def int) int {
    if val, err := c.Get(key); err == nil {
        if v, ok := val.(int); ok {
            return v
        }
    }
    return def
}

func (c *Container) Int64(key string, def int64) int64 {
    if val, err := c.Get(key); err == nil {
        if v, ok := val.(int64); ok {
            return v
        }
    }
    return def
}

func (c *Container) Bool(key string, def bool) bool {
    if val, err := c.Get(key); err == nil {
        if v, ok := val.(bool); ok {
            return v
        }
    }
    return def
}

func (c *Container) Float(key string, def float64) float64 {
    if val, err := c.Get(key); err == nil {
        if v, ok := val.(float64); ok {
            return v
        }
    }
    return def
}

// 读取数据
func (yaml *Config) Reader(filename string) (handler config.Handler, err error) {
    data, err := ReadYaml(filename)
    if err != nil {
        return
    }
    handler = &Container{
        data: data,
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
    err = yaml.Unmarshal(bytes, &data)
    if err != nil {
        return nil, err
    }
    return
}

func init() {
    config.Register("yaml", &Config{})
}