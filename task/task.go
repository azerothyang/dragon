package main

import (
	"fmt"
	"time"
)

func TestTask() {
	ticker := time.Tick(time.Second)
	for tk := range ticker {
		fmt.Println(tk.Unix())
	}
}

func main() {
	TestTask()
}
