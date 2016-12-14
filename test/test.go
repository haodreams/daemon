package main

import (
	"fmt"
	"time"

	"github.com/haodreams/daemon"
)

func main() {
	if err := daemon.Run(); err != nil {
		return
	}
	fmt.Println("start")
	var a *int
	time.Sleep(time.Second * 5)
	*a = 10
}
