package repl

import (
	"bufio"
	"fmt"
	"os"
)

func printPrompt(p string) {
	fmt.Print(p)
}

func StartRepl(c chan string, con chan bool, prompt string, termPattern string) {
	reader := bufio.NewScanner(os.Stdin)

	printPrompt(prompt)
	for reader.Scan() {
		t := reader.Text()
		if t == termPattern {
			close(c)
			break
		}

		c <- t

		// wait until we can contnue
		<-con
		printPrompt(prompt)
	}
}
