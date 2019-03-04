package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"strings"
	"os/exec"
)

func main() {	
	fpath := os.Args[1]
	b, err:= ioutil.ReadFile(fpath)
	if err!=nil {
		log.Fatal(err)
	}
	funcs := strings.Split(string(b), "\n")
	for i := 0; i < len(funcs) ; i++ {
		// fmt.Println(funcs[i])
		b, err := exec.Command("go", "get", "-u", funcs[i]).Output()
		if err!=nil {
			log.Fatal("error in sub", err)
		}
		fmt.Println(string(b))
	}
} 
