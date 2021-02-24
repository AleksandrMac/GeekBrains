package config

import "fmt"

func Example() {
	configuration, _ := Config("config.txt", "port", "address")

	fmt.Printf("address: = %v:%v", configuration["address"], configuration["port"])
}
