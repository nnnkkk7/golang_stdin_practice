package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("input your name!!!")

}

func input() string {
	// 標準入力を受け取る
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	text := stdin.Text()
	return text
}
