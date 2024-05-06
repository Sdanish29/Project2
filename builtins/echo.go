package builtins

import "fmt"

func Echo(args []string) {
	for _, arg := range args {
	fmt.Print(arg + " ")
	}
	fmt.Println()
}
