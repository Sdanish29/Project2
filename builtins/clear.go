package builtins

import "os"
import "os/exec"

func Clear(){
	cmd:= exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
