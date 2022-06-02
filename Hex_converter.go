package main

// Hex 16진수 문자열 변환기
// 사용법: Decode는 0x를 맨앞에 붙여 입력, Encode는 그냥 문자열 입력
import (
	"bufio"
	"encoding/hex"
	"fmt"
	"os"
)

func main() {
	src := bufio.NewReader(os.Stdin)
	line, _ := src.ReadString('\n')

	if line[:2] == "0x" {
		line, _ := hex.DecodeString(line[2:])
		fmt.Printf("values: %s", line)
	} else {
		fmt.Print("values: ", hex.EncodeToString([]byte(line)))
	}
}
