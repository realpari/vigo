package main

import (
	"bufio"
	"fmt"
	"github.com/realpari/vigo/motions"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("read error:", err)
			continue
		}
		input = strings.TrimSpace(input)

		output := motions.Vigo(input)
		fmt.Println(output)
	}
}
