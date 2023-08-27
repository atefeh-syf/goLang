package main

type Config struct {
	ConnectionString string
}

func main() {
	for i := 0; i < 100; i++ {
		//config := GetConfigWithSingleton()
		config := GetConfigWithOnce()
		println(i, ":", &config)
	}

}