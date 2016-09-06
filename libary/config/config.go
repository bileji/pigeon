package config

import (
    "fmt"
)

type Handler interface {
    Set(key string, val string) error
    String(key, def string) string
    Strings(key string, def []string) []string
    Int(key string, def int) int
    Int64(key string, def int64) int64
    Bool(key string, def bool) bool
    Float(key string, def float64) float64
}

type Config interface {
    Reader(filename string) (Handler, error)
    Writer(data []byte) (Handler, error)
}

var adapters = make(map[string]Config)

func Register(name string, adapter Config) {
    if adapter == nil {
        panic("config: Register adapter is nil")
    }
    if _, ok := adapters[name]; ok == true {
        panic("config: Register called twice for adapter " + name)
    }
    adapters[name] = adapter
}

func NewConfig(adapterName, filename string) (Handler, error) {
    if adapter, ok := adapters[adapterName]; ok {
        return adapter.Reader(filename)
    } else {
        return nil, fmt.Errorf("config: unknown adapter %q", adapterName)
    }
}