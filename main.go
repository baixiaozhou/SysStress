package main

import (
	"SysStrain/cmd"
	"fmt"
	"os"
)

func main() {
	err := cmd.Main(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
