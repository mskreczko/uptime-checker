package main

import (
	"fmt"
	"github.com/mskreczko/uptime-checker/internal"
)

func main() {
	config := internal.ReadConfig("./config.yaml")
	fmt.Printf("%+v\n", config)
}
