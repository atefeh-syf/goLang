package main

import "sync"
var once sync.Once

func  GetConfigWithOnce() *Config {
	once.Do(func() {
		config = &Config{}
		println("Creating config")
	})
	println("Get config from old instance")
	return config
}