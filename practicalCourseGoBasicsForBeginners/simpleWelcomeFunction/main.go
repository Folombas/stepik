package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func greetUser(name string) {
	fmt.Printf("Привет, %s!\n", name)
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	greetUser(name)
}
