package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scan := bufio.NewScanner(os.Stdin)

	for scan.Scan() {
		// 大文字に変換する
		ucl := strings.ToUpper(scan.Text())
		fmt.Println(ucl)
	}
}
