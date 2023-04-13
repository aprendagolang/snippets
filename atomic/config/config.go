package config

import (
	"sync/atomic"
)

type cfg struct {
	Host string
	Port string
	User string
	Pass string
}

var (
	c cfg
	a = atomic.Pointer[cfg]{}
)

func init() {
	c = cfg{
		Host: "localhost",
		Port: "5432",
		User: "aprenda",
		Pass: "golang",
	}

	a.Store(&c)
}

func Get() cfg {
	return c
}

func GetAtomic() *cfg {
	return a.Load()
}
