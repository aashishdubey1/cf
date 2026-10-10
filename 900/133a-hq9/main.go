package main

import (
	"fmt"
	"strings"
)

func main() {
	var p string
	fmt.Scan(&p)

	if strings.ContainsAny(p, "HQ9") {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
