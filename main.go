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
			// fmt.Println(funcs[i])
			b, err := exec.Command("go", "get", funcs[i]).Output()
			if err != nil {
				log.Fatal("error in sub", err)
			}
			fmt.Println(string(b))
		}
	} else {
		fmt.Println("Insufficient Arguments")
	}
}
