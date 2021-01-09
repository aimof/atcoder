package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Buffer(make([]byte, 0, 1000000), 1000000)
	for stdin.Scan() {
		in := stdin.Text()
		out, _ := do(in)
		fmt.Println(out)
	}

}

func do(in string) (string, error) {
	return "Hello, world!", nil
}
