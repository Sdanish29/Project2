package builtins

import (
	"fmt"
	"os" )

func Pwd() {
	dir, err := os.Getwd()
	if err !=nil {
	fmt.Println("Error:", err)
	return
}
fmt.Println(dir)
}
