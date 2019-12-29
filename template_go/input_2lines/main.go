package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	in0 := stdin.Text()
	stdin.Scan()
	in1 := stdin.Text()
	out, _ := do(in0, in1)
	fmt.Println(out)
}

func do(in0, in1 string) (string, error) {
	return "Hello, world!", nil
}
