package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
	fmt.Println(time.Since(start).Seconds())

	start = time.Now()
	fmt.Println(strings.Join(os.Args, " "))
	fmt.Println(time.Since(start).Seconds())
}
