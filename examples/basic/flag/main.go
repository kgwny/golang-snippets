package main

import (
	"flag"
	"fmt"
)

/* $ go run ./examples/basic/flag/main.go -usage
 * flag provided but not defined: -usage
 * Usage of ・・・:
 *	-name string
 *	  	Your name (default "Go")
 * exit status 2
 *
 * $ go run ./examples/basic/flag/main.go -name=hogehoge
 * Hello, hogehoge
 */
func main() {
	name := flag.String("name", "Go", "Your name")
	flag.Parse()
	fmt.Println("Hello,", *name)
}
