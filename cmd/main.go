package main

import (
	"bufio"
	"fmt"
	"github.com/realpari/motions"
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

		output := motions.basicMotion(input)
		fmt.Println(output)
	}
}
