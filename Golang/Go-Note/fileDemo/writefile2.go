package main

import (
	"bufio"
	"fmt"
	"os"
)

// 複寫已存在檔案
func main() {
	filePath := "fileDemo/a.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 066)
	if err != nil {
		fmt.Printf("Open file err=%v\n", err)
		return
	}
	defer file.Close()

	str := "Go go go \n"
	writer := bufio.NewWriter(file) // buffer
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}
	writer.Flush() // 寫入數據
}