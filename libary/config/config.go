package config

type Handler interface {
    Set(key string, val string) error
    String(key, def string) (string, error)
    Strings(key string, def []string) (string, error)
    Int(key string, def int) (int, error)
    Int64(key string, def int64) (int64, error)
    Bool(key string, def bool) (bool, error)
    Float(key string, def float64) (float64, error)
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
}