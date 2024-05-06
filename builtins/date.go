package builtins

import (
    "fmt"
    "time"
)


func Date() {
    now := time.Now()
    fmt.Println(now.Format(time.RFC1123))
}


