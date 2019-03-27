package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) == 2 {
		fpath := os.Args[1]
		b, err := ioutil.ReadFile(fpath)
		if err != nil {
			log.Fatal(err)
		}
		funcs := strings.Split(string(b), "\n")
		for i := 0; i < len(funcs); i++ {
			cmd := exec.Command("go", "get", "-v", funcs[i])
			err = cmd.Run()
			if err != nil {
				log.Fatal("Invalid Package" + funcs[i])
			}
		}
	} else {
		fmt.Println("Insufficient Arguments")
	}
}
