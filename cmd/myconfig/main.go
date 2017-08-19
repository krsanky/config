package main

import (
	"fmt"

	"github.com/krsanky/config"
)

func main() {
	fmt.Println("main...")
	val1 := config.Get("section1.key1")
	fmt.Printf("val1:%s\n", val1)
}
