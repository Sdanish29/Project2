package builtins

import (
	"fmt"
	"io/ioutil" )

func Ls(){
files, err := ioutil.ReadDir(".")
if err != nil {
fmt.Println("Error:", err)
return
}
for _, file := range files {
	fmt.Println(file.Name())
}
}
