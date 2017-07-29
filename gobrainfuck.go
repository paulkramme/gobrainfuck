package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"regexp"
)

func main() {
	verbose := flag.Bool("v", false, "Verbose")
	flag.Parse()
	fmt.Println(flag.Arg(0))
	file, err := ioutil.ReadFile(flag.Arg(0))
	if err != nil {
		panic(err)
	}
	re := regexp.MustCompile(`[+\-,.[\]<>]`)
	code := re.FindAll(file, -1) // I am very sorry for this, this 2d array was fucking hard

	loopfilo := make([]int, 1)
	register := make([]rune, 20)
	registerpointer := 0
	codepointer := 0
	//fmt.Println(string(code[1][0]))
	for {
		switch i := code[codepointer][0]; i {
		case '>':
			if registerpointer == len(register)-1 {
				register = append(register, 0)
			}
			registerpointer++
		case '<':
			if registerpointer != 0 {
				registerpointer--
			}
		case '+':
			register[registerpointer] = register[registerpointer] + 1
		case '-':
			register[registerpointer]--
		case '.':
			fmt.Println(string(register[registerpointer]))
		case ',':
			fmt.Println("Not yet supported")
		case '[':
			loopfilo = append(loopfilo, codepointer)
		case ']':
			if register[registerpointer] == 0 {
				_, loopfilo = loopfilo[len(loopfilo)-1], loopfilo[:len(loopfilo)-1] //pop
			} else {
				codepointer = loopfilo[len(loopfilo)-1] //jump back
			}
		default:
		}
		if *verbose {
			fmt.Println(register, registerpointer)
		}
		if codepointer != len(code)-1 {
			codepointer++
		} else {
			return
		}
	}

}
