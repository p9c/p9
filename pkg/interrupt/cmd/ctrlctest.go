package main

import (
	"fmt"

	"github.com/p9c/p9/pkg/interrupt"
)

func main() {
	interrupt.AddHandler(func() {
		fmt.Println("IT'S THE END OF THE WORLD!")
	})
	<-interrupt.HandlersDone
}
