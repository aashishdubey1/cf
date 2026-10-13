package main

import (
	"fmt"
)

func main() {
	var dest int
	fmt.Scan(&dest)

	steps := 0

	for i := 5; i > 0; i-- {
		if dest == 0 {
			fmt.Println(steps)
			return
		}
		if i > dest {
			continue
		}
		rem := dest / i
		steps += rem
		dest = dest % i
	}
	fmt.Println(steps)
}

// Short method
// 	fmt.Println((dest + 4) / 5)
