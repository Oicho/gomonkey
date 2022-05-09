package main

import (
	"fmt"
	"gomonkey/repl"
	"os"
)

func main() {
	fmt.Println("Hello World!")
	//reader := bufio.NewReader(os.Stdin)
	repl.MainLoop(os.Stdin, os.Stdout)

	/*
		for {
			fmt.Print("$ ")
			line, _ := reader.ReadString('\n')
			line = line[0 : len(line)-1]
			if line == "quit" {
				fmt.Println("GoodBye!")
				break
			} else {
				repl.
					fmt.Println(line)
			}
		}
	*/
}
