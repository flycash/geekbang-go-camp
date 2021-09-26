package main

import "time"

// Env 会在编译的时候注入值

// Env go build -ldflags "-X 'main.Env=aaa'"
var Env = "dev"
func main() {
	cfg := InitConfig(Env)
	println(cfg.Name)
}

type Config struct {
	Name string
	Dsn string
	Timeout time.Duration
}

func InitConfig(env string) Config {
	if env == "dev" {
		return initDevConfig()
	}
	return initProdConfig()
}

func initDevConfig() Config {
	return Config{
		Name: "This is dev",
		Dsn: "localhost:8080",
		Timeout: 10 * time.Second,
	}
}

func initProdConfig() Config {
	return Config{
		Name: "This is prod",
		Dsn: "your.com",
		Timeout: 10 * time.Second,
	}
}