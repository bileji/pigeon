package config

import (
    "os"
    "fmt"
    "regexp"
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

func EnvValue(name string) string {
    match := regexp.MustCompile(`^\$\{([a-zA-Z_][\w]+)(\|\|[./\w]*)}$`).FindAllStringSubmatch(name, -1)
    if len(match) > 0 {
        switch true {
        case match[0] > 2:
            val := os.Getenv(match[1])
            if val != "" {
                return val
            }
            return match[2]
        case match[0] > 1:
            val := os.Getenv(match[1])
            if val != "" {
                return val
            }
        }
    }
    return name
}

func EnvValueForMap(m map[string]interface{}) map[string]interface{} {
    for k, v := range m {
        switch value := v.(type) {
        case string:
            m[k] = EnvValue(value)
        case map[string]interface{}:
            m[k] = EnvValueForMap(value)
        case map[string]string:
            for k2, v2 := range value {
                value[k2] = EnvValue(v2)
            }
            m[k] = value
        }
    }
    return m
}