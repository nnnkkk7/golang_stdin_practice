package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("input your name!!!")
	i := input()
	//一文字ずつ分割
	slice := strings.Split(i, "")
	for _, s := range slice {
		fmt.Printf("[%s]", s)
	}

}

func input() string {
	// 標準入力を受け取る
	stdin := bufio.NewScanner(os.Stdin)
	stdin.Scan()
	text := stdin.Text()
	return text
}
