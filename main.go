package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print(strings.ToLower(strings.Join(os.Args[1:], "_")))
}
